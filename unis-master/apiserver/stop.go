package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func handlePublicStop(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	instanceID := c.Param("instanceID")

	if validateUser(username, password) {
		nodesInfo := getPublicNodesInfo()
		for indexNode, node := range nodesInfo {
			for indexInstance, instance := range node.Instances {
				if strings.Contains(instance.InstanceID, instanceID) {
					resp, err := http.PostForm("http://127.0.0.1:10000/instance/stop/public/"+instanceID, url.Values{"nodename": {node.NodeName}})
					if err != nil {
						return c.String(http.StatusNotImplemented, err.Error())
					}

					body, err := ioutil.ReadAll(resp.Body)

					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to read resp.Body")
					}

					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusNotImplemented, string(body))
					}

					// delete instance
					nodesInfo[indexNode].Instances = append(nodesInfo[indexNode].Instances[:indexInstance], nodesInfo[indexNode].Instances[indexInstance+1:]...)
					nodesInfoInJSON, err := json.Marshal(nodesInfo)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to delete instance")
					}

					err = ioutil.WriteFile(serverFilePath.NodesPublicPath+"nodesInfo.json", nodesInfoInJSON, os.ModePerm)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to write nodesInfo.json")
					}

					return c.String(http.StatusOK, string(body))
				}
			}
		}
		return c.String(http.StatusNotImplemented, "wrong instance ID")
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePrivateStop(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	instanceID := c.Param("instanceID")

	if validateUser(username, password) {
		nodesInfo := getPrivateNodesInfo(username)
		for indexNode, node := range nodesInfo {
			for indexInstance, instance := range node.Instances {
				if strings.Contains(instance.InstanceID, instanceID) {
					resp, err := http.PostForm("http://127.0.0.1:10000/instance/stop/"+username+"/"+instanceID, url.Values{"nodename": {node.NodeName}})
					if err != nil {
						return c.String(http.StatusNotImplemented, err.Error())
					}

					body, err := ioutil.ReadAll(resp.Body)

					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to read resp.Body")
					}

					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusNotImplemented, string(body))
					}

					// delete instance
					nodesInfo[indexNode].Instances = append(nodesInfo[indexNode].Instances[:indexInstance], nodesInfo[indexNode].Instances[indexInstance+1:]...)
					nodesInfoInJSON, err := json.Marshal(nodesInfo)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to delete instance")
					}

					err = ioutil.WriteFile(serverFilePath.NodesPath+username+"/nodesInfo.json", nodesInfoInJSON, os.ModePerm)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail to write nodesInfo.json")
					}

					return c.String(http.StatusOK, string(body))
				}
			}
		}
		return c.String(http.StatusNotImplemented, "wrong instance ID")
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
