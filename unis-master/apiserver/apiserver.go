package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/seveirbian/unis/unis-master/scheduler"
	"github.com/sirupsen/logrus"
)

var UsersInfo []userInfo

var serverFilePath = ServerFilePath{
	UnisPath:         os.Getenv("HOME") + "/.unis/",
	RootPath:         os.Getenv("HOME") + "/.unis/apiserver/",
	ImagesPath:       os.Getenv("HOME") + "/.unis/apiserver/images/",
	NodesPath:        os.Getenv("HOME") + "/.unis/apiserver/nodes/",
	ImagesPublicPath: os.Getenv("HOME") + "/.unis/apiserver/images/public/",
	NodesPublicPath:  os.Getenv("HOME") + "/.unis/apiserver/nodes/public/",
	UsersJSONPath:    os.Getenv("HOME") + "/.unis/apiserver/",
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	//create server dir
	serverFilePath.createFilePath()

	//load users.json
	serverFilePath.loadUsersJSON()

}

func (apiServer Server) Serve(serveIP string) error {
	// handler version
	apiServer.Version = "0.001.0.0"

	server := echo.New()

	// serve UNISCTL and UNISLET
	// serve "unisctl connect" command
	server.GET("/", handleConnect)

	// serve "unisctl signin" command
	server.POST("/users/signin", handleSignin)

	// serve UNISCTL
	// serve "unisctl signup" command
	server.POST("/users/signup", handleSignup)

	// serve "unisctl images" command
	server.POST("/images/show/:username/images", handlePrivateImages)
	server.POST("/images/show/public/images", handlePublicImages)

	// serve "unisctl push" command
	server.POST("/images/push/public/:imagename", handlePublicPush)
	server.POST("/images/push/:username/:imagename", handlePrivatePush)

	// serve "unisctl rmi" command
	server.POST("/images/remove/public/:imageID", handlePublicRmi)
	server.POST("/images/remove/:username/:imageID", handlePrivateRmi)

	// serve "unisctl tag" command
	server.POST("/images/tag/public/:oldimage/:oldtag/:newimage/:newtag", handlePublicTag)
	server.POST("/images/tag/:username/:oldimage/:oldtag/:newimage/:newtag", handlePrivateTag)

	// serve "unisctl pull" command
	server.POST("/images/pull/public/:imageID", handlePublicPull)
	server.POST("/images/pull/:username/:imageID", handlePrivatePull)

	// // serve "unisctl run" command
	server.POST("/instances/run/public/:imageID", handlePublicRun)
	server.POST("/instances/run/:username/:imageID", handlePrivateRun)
	// // serve "unisctl stop" command
	// server.POST()
	// // serve "unisctl ps" command
	// server.POST()

	// // serve "unisctl version" command
	// server.POST()
	// // serve "unistl nodes" command
	server.POST("/nodes/show/public/nodes", handlePublicNodes)
	server.POST("/nodes/show/:username/nodes", handlePrivateNodes)
	// // serve "unisctl stats" command
	// server.POST()

	// serve UNISLET
	// serve "unislet add" command
	server.POST("/nodes/add/public/:nodename", handlePublicAdd)
	server.POST("/nodes/add/:username/:nodename", handlePrivateAdd)

	// Run request-handler, this should never exit
	return server.Start(serveIP)
}

func handlePublicRun(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	imageID := c.FormValue("imageid")
	argument := c.FormValue("argument")
	command := c.FormValue("command")
	requestCPU := c.FormValue("requestcpu")
	requestMem := c.FormValue("requestmem")
	// maxCPU := c.FormValue("maxcpu")
	// maxMem := c.FormValue("maxmem")

	if validateUser(username, password) {
		// validate that the image exists
		publicImagesInfo := getPublicImagesInfo()
		for _, image := range publicImagesInfo {
			if strings.Contains(image.ImageID, imageID) {

				// get the node that to deploy instance
				publicNodesInfoSchedule := scheduler.GetPublicNodesInfo()
				nodename, err := scheduler.Schedule(scheduler.FirstFit, publicNodesInfoSchedule, image.Type, requestCPU, requestMem)
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}

				fmt.Println(nodename)

				// send node and image to controller
				resp, err := http.PostForm("http://127.0.0.1:10000/instances/run/public/"+image.ImageID, url.Values{"nodename": {nodename}, "argument": {argument}, "command": {command}, "type": {image.Type}})
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}
				if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
					return c.String(http.StatusBadRequest, "failed to deploy instance")
				}

				// add image and instance to node
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return c.String(http.StatusNotImplemented, "failed to deploy instance")
				}
				instanceID := string(body)

				publicNodesInfo := getPublicNodesInfo()
				for index, node := range publicNodesInfo {

					if node.NodeName == nodename {

						// node has the image
						if resp.StatusCode == http.StatusOK {
							publicNodesInfo[index].Images = append(publicNodesInfo[index].Images, ImageInfo{
								Repository: image.Repository,
								Tag:        image.Tag,
								ImageID:    image.ImageID,
								Created:    image.Created,
								Size:       image.Size,
								Type:       image.Type,
								Owner:      image.Owner,
							})
						}

						neededCPU, _ := strconv.Atoi(requestCPU)
						neededMem, _ := strconv.Atoi(requestMem)
						publicNodesInfo[index].Instances = append(publicNodesInfo[index].Instances, Instance{
							ImageRepository: image.Repository,
							ImageTag:        image.Tag,
							ImageID:         image.ImageID,
							InstanceID:      instanceID,
							RequestCPU:      int64(neededCPU),
							RequestMem:      int64(neededMem),
						})
					}
				}
				fmt.Println(publicNodesInfo)
				publicNodesInfoInJSON, err := json.Marshal(publicNodesInfo)
				if err != nil {
					logrus.Fatal(err)
				}
				err = ioutil.WriteFile(serverFilePath.NodesPublicPath+"nodesInfo.json", publicNodesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}
				return c.String(http.StatusOK, "instance deployed")
			}
		}
		return c.String(http.StatusNotFound, "incorrect image ID")
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePrivateRun(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	imageID := c.FormValue("imageid")
	argument := c.FormValue("argument")
	command := c.FormValue("command")
	requestCPU := c.FormValue("requestcpu")
	requestMem := c.FormValue("requestmem")
	// maxCPU := c.FormValue("maxcpu")
	// maxMem := c.FormValue("maxmem")

	if validateUser(username, password) {
		// validate that the image exists
		privateImagesInfo := getPrivateImagesInfo(username)
		for _, image := range privateImagesInfo {
			fmt.Println(image.ImageID)
			fmt.Println(imageID)
			if strings.Contains(image.ImageID, imageID) {

				fmt.Println(image.ImageID)
				fmt.Println(imageID)

				// get the node that to deploy instance
				privateNodesInfoSchedule := scheduler.GetPrivateNodesInfo(username)
				nodename, err := scheduler.Schedule(scheduler.FirstFit, privateNodesInfoSchedule, image.Type, requestCPU, requestMem)
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}

				fmt.Println(nodename)

				// send node and image to controller
				resp, err := http.PostForm("http://127.0.0.1:10000/instances/run/"+username+"/"+image.ImageID, url.Values{"nodename": {nodename}, "argument": {argument}, "command": {command}, "type": {image.Type}})
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}
				if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
					return c.String(http.StatusBadRequest, "failed to deploy instance")
				}

				// add image and instance to node
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return c.String(http.StatusNotImplemented, "failed to deploy instance")
				}
				instanceID := string(body)

				privateNodesInfo := getPrivateNodesInfo(username)
				for index, node := range privateNodesInfo {

					if node.NodeName == nodename {

						// node has the image
						if resp.StatusCode == http.StatusOK {
							privateNodesInfo[index].Images = append(privateNodesInfo[index].Images, ImageInfo{
								Repository: image.Repository,
								Tag:        image.Tag,
								ImageID:    image.ImageID,
								Created:    image.Created,
								Size:       image.Size,
								Type:       image.Type,
								Owner:      image.Owner,
							})
						}

						neededCPU, _ := strconv.Atoi(requestCPU)
						neededMem, _ := strconv.Atoi(requestMem)
						privateNodesInfo[index].Instances = append(privateNodesInfo[index].Instances, Instance{
							ImageRepository: image.Repository,
							ImageTag:        image.Tag,
							ImageID:         image.ImageID,
							InstanceID:      instanceID,
							RequestCPU:      int64(neededCPU),
							RequestMem:      int64(neededMem),
						})
					}
				}
				fmt.Println(privateNodesInfo)
				privateNodesInfoInJSON, err := json.Marshal(privateNodesInfo)
				if err != nil {
					logrus.Fatal(err)
				}
				err = ioutil.WriteFile(serverFilePath.NodesPath+username+"/nodesInfo.json", privateNodesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}
				return c.String(http.StatusOK, "instance deployed")
			}
		}
		return c.String(http.StatusNotFound, "incorrect image ID")
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
