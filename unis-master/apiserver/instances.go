package apiserver

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func handleAllInstances(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		//get public Nodes info
		publicNodesInfo := getPublicNodesInfo()

		// //get private Nodes info
		// privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		// var blankLenth = 10
		for _, node := range publicNodesInfo {

			for _, instance := range node.Instances {
				bodyContent += node.NodeName
				bodyContent += " "

				bodyContent += instance.ImageRepository
				bodyContent += " "

				bodyContent += instance.ImageTag
				bodyContent += " "

				bodyContent += Substring(instance.ImageID, 0, 10)
				bodyContent += " "

				bodyContent += instance.InstanceID
				bodyContent += " "

				bodyContent += string(instance.RequestCPU)
				bodyContent += " "

				bodyContent += string(instance.RequestMem)
				bodyContent += " "

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
		// var blankLenth = 10

		for _, node := range privateNodesInfo {
			for _, instance := range node.Instances {
				bodyContent += node.NodeName
				bodyContent += " "

				bodyContent += instance.ImageRepository
				bodyContent += " "

				bodyContent += instance.ImageTag
				bodyContent += " "

				bodyContent += Substring(instance.ImageID, 0, 10)
				bodyContent += " "

				bodyContent += instance.InstanceID
				bodyContent += " "

				bodyContent += strconv.Itoa(int(instance.RequestCPU))
				bodyContent += " "

				bodyContent += strconv.Itoa(int(instance.RequestMem))
				bodyContent += " "

				bodyContent += "\n"
			}
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
