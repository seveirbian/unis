package apiserver

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func handleAllInstances(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		//get public Nodes info
		publicNodesInfo := getPublicNodesInfo()

		//get private Nodes info
		privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		var blankLenth = 10
		for _, node := range publicNodesInfo {

			for _, instance := range node.Instances {
				bodyContent += node.NodeName
				bodyContent += EmptyString(strings.Count("Node", "") +
					blankLenth - strings.Count(node.NodeName, ""))

				bodyContent += instance.ImageRepository
				bodyContent += EmptyString(strings.Count("Repository", "") +
					blankLenth - strings.Count(instance.ImageRepository, ""))

				bodyContent += instance.ImageTag
				bodyContent += EmptyString(strings.Count("Tag", "") +
					blankLenth - strings.Count(instance.ImageTag, ""))

				bodyContent += Substring(instance.ImageID, 0, 10)
				bodyContent += EmptyString(strings.Count("Image ID", "") +
					blankLenth - strings.Count(Substring(instance.ImageID, 0, 10), ""))

				bodyContent += instance.InstanceID
				bodyContent += EmptyString(strings.Count("Instance ID", "") +
					blankLenth - strings.Count(instance.InstanceID, ""))

				bodyContent += string(instance.RequestCPU)
				bodyContent += EmptyString(strings.Count("Request CPU", "") +
					blankLenth - strings.Count(string(instance.RequestCPU), ""))

				bodyContent += string(instance.RequestMem)
				bodyContent += EmptyString(strings.Count("Request Mem", "") +
					blankLenth - strings.Count(string(instance.RequestMem), ""))

				bodyContent += "\n"
			}
		}
		for _, node := range privateNodesInfo {
			for _, instance := range node.Instances {
				bodyContent += node.NodeName
				bodyContent += EmptyString(strings.Count("Node", "") +
					blankLenth - strings.Count(node.NodeName, ""))

				bodyContent += instance.ImageRepository
				bodyContent += EmptyString(strings.Count("Repository", "") +
					blankLenth - strings.Count(instance.ImageRepository, ""))

				bodyContent += instance.ImageTag
				bodyContent += EmptyString(strings.Count("Tag", "") +
					blankLenth - strings.Count(instance.ImageTag, ""))

				bodyContent += Substring(instance.ImageID, 0, 10)
				bodyContent += EmptyString(strings.Count("Image ID", "") +
					blankLenth - strings.Count(Substring(instance.ImageID, 0, 10), ""))

				bodyContent += instance.InstanceID
				bodyContent += EmptyString(strings.Count("Instance ID", "") +
					blankLenth - strings.Count(instance.InstanceID, ""))

				bodyContent += strconv.Itoa(int(instance.RequestCPU))
				bodyContent += EmptyString(strings.Count("Request CPU", "") +
					blankLenth - strings.Count(strconv.Itoa(int(instance.RequestCPU)), ""))

				bodyContent += strconv.Itoa(int(instance.RequestMem))
				bodyContent += EmptyString(strings.Count("Request Mem", "") +
					blankLenth - strings.Count(strconv.Itoa(int(instance.RequestMem)), ""))

				bodyContent += "\n"
			}
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePrivateInstances(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {

		//get private Nodes info
		privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		var blankLenth = 10

		for _, node := range privateNodesInfo {
			for _, instance := range node.Instances {
				bodyContent += node.NodeName
				bodyContent += EmptyString(strings.Count("Node", "") +
					blankLenth - strings.Count(node.NodeName, ""))

				bodyContent += instance.ImageRepository
				bodyContent += EmptyString(strings.Count("Repository", "") +
					blankLenth - strings.Count(instance.ImageRepository, ""))

				bodyContent += instance.ImageTag
				bodyContent += EmptyString(strings.Count("Tag", "") +
					blankLenth - strings.Count(instance.ImageTag, ""))

				bodyContent += Substring(instance.ImageID, 0, 10)
				bodyContent += EmptyString(strings.Count("Image ID", "") +
					blankLenth - strings.Count(Substring(instance.ImageID, 0, 10), ""))

				bodyContent += instance.InstanceID
				bodyContent += EmptyString(strings.Count("Instance ID", "") +
					blankLenth - strings.Count(instance.InstanceID, ""))

				bodyContent += strconv.Itoa(int(instance.RequestCPU))
				bodyContent += EmptyString(strings.Count("Request CPU", "") +
					blankLenth - strings.Count(strconv.Itoa(int(instance.RequestCPU)), ""))

				bodyContent += strconv.Itoa(int(instance.RequestMem))
				bodyContent += EmptyString(strings.Count("Request Mem", "") +
					blankLenth - strings.Count(strconv.Itoa(int(instance.RequestMem)), ""))

				bodyContent += "\n"
			}
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
