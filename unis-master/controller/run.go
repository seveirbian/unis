package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

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
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+image.ImageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}, "dockerID": {image.DockerID}})
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

					dockerID := strings.Split(string(body), " ")[0]
					instanceID := strings.Split(string(body), " ")[1]

					// let apiserver know this node has the image
					return c.String(http.StatusAccepted, dockerID+" "+instanceID)
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
			resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+imageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}, "dockerID": {""}})
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

			dockerID := strings.Split(string(body), " ")[0]
			instanceID := strings.Split(string(body), " ")[1]

			// let apiserver know this node has the image
			return c.String(http.StatusOK, dockerID+" "+instanceID)
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
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+image.ImageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}, "dockerID": {image.DockerID}})
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

					dockerID := strings.Split(string(body), " ")[0]
					instanceID := strings.Split(string(body), " ")[1]

					// let apiserver know this node has the image
					return c.String(http.StatusAccepted, dockerID+" "+instanceID)
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
			resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/run/"+imageID, url.Values{"argument": {argument}, "command": {command}, "imageType": {imageType}, "dockerID": {""}})
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

			dockerID := strings.Split(string(body), " ")[0]
			instanceID := strings.Split(string(body), " ")[1]

			// let apiserver know this node has the image
			return c.String(http.StatusOK, dockerID+" "+instanceID)
		}
	}
	return c.String(http.StatusNotImplemented, "node info is different between apiserver and controller")
}
