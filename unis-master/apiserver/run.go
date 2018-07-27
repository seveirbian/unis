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

				dockerID := strings.Split(string(body), " ")[0]
				instanceID := strings.Split(string(body), " ")[1]

				publicNodesInfo := getPublicNodesInfo()
				for index, node := range publicNodesInfo {

					if node.NodeName == nodename {

						// node has the image
						if resp.StatusCode == http.StatusOK {
							publicNodesInfo[index].Images = append(publicNodesInfo[index].Images, ImageInfo{
								Repository: image.Repository,
								Tag:        image.Tag,
								ImageID:    image.ImageID,
								DockerID:   dockerID,
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
							DockerID:        dockerID,
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
			if strings.Contains(image.ImageID, imageID) {

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
				dockerID := strings.Split(string(body), " ")[0]
				instanceID := strings.Split(string(body), " ")[1]

				privateNodesInfo := getPrivateNodesInfo(username)
				for index, node := range privateNodesInfo {

					if node.NodeName == nodename {

						// node has the image
						if resp.StatusCode == http.StatusOK {
							privateNodesInfo[index].Images = append(privateNodesInfo[index].Images, ImageInfo{
								Repository: image.Repository,
								Tag:        image.Tag,
								ImageID:    image.ImageID,
								DockerID:   dockerID,
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
							DockerID:        dockerID,
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
