package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Version string
}

type ImageInfo struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
	Type       string
	Owner      string
}

type Instance struct {
	ImageRepository string
	ImageTag        string
	ImageID         string
	DockerID        string
	InstanceID      string

	RequestCPU int64
	RequestMem int64
	// LimitCPU   int64
	// LimitMem   int64

	Type string
}

type NodeInfo struct {
	NodeName       string
	NodeAddr       string
	NodeType       string // public or private
	NodeEnv        string // Docker or Unikernel
	DockerInfo     string
	HypervisorInfo string

	TotalCPU int64
	TotalMem int64

	Images    []ImageInfo
	Instances []Instance
}

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var publicNodes = []NodeInfo{}
var usersInfo = []userInfo{}
var privateNodes = make(map[string][]NodeInfo)

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	// wait for apiserver to start
	time.Sleep(1000)

	// load nodes
	// err := loadNodes()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
}

func (controller Controller) Start(serveIP string) error {
	controller.Version = "0.0001.0.0"

	cServer := echo.New()

	// handle test
	cServer.GET("/", handleConnect)
	// // handle add user
	// cServer.POST("/users/add/:username", handleAddUser)
	// // handle add node
	// cServer.POST("/nodes/add/public/:nodename", handleAddPublicNode)
	// cServer.POST("/nodes/add/:username/:nodename", handleAddPrivateNode)
	// // handle remove node
	// cServer.POST("/nodes/remove/:nodename", handleRemoveNode)
	// // handle run instance
	cServer.POST("/instances/run/public/:imageID", handleRunPublicInstance)
	cServer.POST("/instances/run/:username/:imageID", handleRunPrivateInstance)
	// // handle stop instance
	// cServer.POST("/instance/stop/:instanceID", handleStopInstance)

	return cServer.Start(serveIP)
}

func handleRunPublicInstance(c echo.Context) error {
	imageID := c.Param("imageID")
	nodename := c.FormValue("nodename")
	imageType := c.FormValue("type")
	// maxCPU := c.FormValue("maxCPU")
	// maxMem := c.FormValue("maxMem")
	argument := c.FormValue("argument")
	command := c.FormValue("command")

	// load nodes
	err := loadNodes()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, node := range publicNodes {
		if node.NodeName == nodename {
			for _, image := range node.Images {
				if image.ImageID == imageID {
					// ask unislet to run instance
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+image.ImageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}})
					if err != nil {
						return c.String(http.StatusBadRequest, err.Error())
					}
					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusBadRequest, "failed to deploy instance")
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						logrus.Fatal(err)
					}
					instanceID := string(body)

					// let apiserver know this node has the image
					return c.String(http.StatusAccepted, instanceID)
				}
			}

			fmt.Println("sending image")

			// send image to node
			arg0 := "curl"
			arg1 := "-F"
			arg2 := imageID + "=@" + os.Getenv("HOME") + "/.unis/apiserver/images/public/" + imageID
			arg3 := "http://" + strings.Split(node.NodeAddr, ":")[0] + ":9899/images/receive/" + imageID

			fmt.Println(arg0, arg1, arg2, arg3)

			child := exec.Command(arg0, arg1, arg2, arg3)
			output, err := child.Output()

			fmt.Println(string(output))

			if err != nil {
				return c.String(http.StatusNotImplemented, "fail to run curl")
			}
			if string(output) != "image sended" {
				return c.String(http.StatusNotImplemented, "fail to send image")
			}

			// ask unislet to run instance
			resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+imageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}})
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			if resp.StatusCode != http.StatusOK {
				return c.String(http.StatusBadRequest, "failed to deploy instance")
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			}
			instanceID := string(body)

			return c.String(http.StatusOK, instanceID)
		}
	}
	return c.String(http.StatusNotImplemented, "node info is different between apiserver and controller")
}

func handleRunPrivateInstance(c echo.Context) error {
	username := c.Param("username")
	imageID := c.Param("imageID")
	nodename := c.FormValue("nodename")
	imageType := c.FormValue("type")
	// maxCPU := c.FormValue("maxCPU")
	// maxMem := c.FormValue("maxMem")
	argument := c.FormValue("argument")
	command := c.FormValue("command")

	// load nodes
	err := loadNodes()
	if err != nil {
		logrus.Fatal(err)
	}

	userNodes := privateNodes[username]

	for _, node := range userNodes {
		if node.NodeName == nodename {
			for _, image := range node.Images {
				if image.ImageID == imageID {
					// ask unislet to run instance
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+image.ImageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}})
					if err != nil {
						return c.String(http.StatusBadRequest, err.Error())
					}
					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusBadRequest, "failed to deploy instance")
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						logrus.Fatal(err)
					}
					instanceID := string(body)

					// let apiserver know this node has the image
					return c.String(http.StatusAccepted, instanceID)
				}
			}

			fmt.Println("sending image")

			// send image to node
			arg0 := "curl"
			arg1 := "-F"
			arg2 := imageID + "=@" + os.Getenv("HOME") + "/.unis/apiserver/images/" + username + "/" + imageID
			arg3 := "http://" + strings.Split(node.NodeAddr, ":")[0] + ":9899/images/receive/" + imageID

			fmt.Println(arg0, arg1, arg2, arg3)

			child := exec.Command(arg0, arg1, arg2, arg3)
			output, err := child.Output()

			fmt.Println(string(output))

			if err != nil {
				return c.String(http.StatusNotImplemented, "fail to run curl")
			}
			if string(output) != "image sended" {
				return c.String(http.StatusNotImplemented, "fail to send image")
			}

			// ask unislet to run instance
			resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+imageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}})
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			if resp.StatusCode != http.StatusOK {
				return c.String(http.StatusBadRequest, "failed to deploy instance")
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			}
			instanceID := string(body)

			return c.String(http.StatusOK, instanceID)
		}
	}
	return c.String(http.StatusNotImplemented, "node info is different between apiserver and controller")
}
