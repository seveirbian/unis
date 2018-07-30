package apiserver

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func handlePublicNodes(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		//get public nodes info
		publicNodesInfo := getPublicNodesInfo()

		//get private nodes info
		// privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		// var blankLenth = 10
		for _, node := range publicNodesInfo {
			bodyContent += node.NodeName
			bodyContent += " "

			bodyContent += strings.Split(node.NodeAddr, ":")[0]
			bodyContent += " "

			bodyContent += node.NodeType
			bodyContent += " "

			bodyContent += node.NodeEnv
			bodyContent += " "

			bodyContent += node.DockerInfo
			bodyContent += " "

			bodyContent += node.HypervisorInfo
			bodyContent += " "

			bodyContent += strconv.Itoa(int(node.TotalCPU))
			bodyContent += " "

			bodyContent += strconv.Itoa(int(node.TotalMem))
			bodyContent += " "

			if node.NodeActive {
				bodyContent += "ACTIVE"
			} else {
				bodyContent += "DEAD"
			}
			bodyContent += " "

			bodyContent += "\n"
		}
		// for _, node := range privateNodesInfo {
		// 	bodyContent += node.NodeName

		// 	bodyContent += strings.Split(node.NodeAddr, ":")[0]

		// 	bodyContent += node.NodeType

		// 	bodyContent += node.NodeEnv

		// 	bodyContent += node.DockerInfo

		// 	bodyContent += node.HypervisorInfo

		// 	bodyContent += strconv.Itoa(int(node.TotalCPU))

		// 	bodyContent += strconv.Itoa(int(node.TotalMem))

		// 	if node.NodeActive {
		// 		bodyContent += "ACTIVE"
		// 	} else {
		// 		bodyContent += "DEAD"
		// 	}

		// 	bodyContent += "\n"
		// }
		return c.String(http.StatusOK, bodyContent)
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePrivateNodes(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {

		//get private nodes info
		privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		for _, node := range privateNodesInfo {
			bodyContent += node.NodeName
			bodyContent += " "

			bodyContent += strings.Split(node.NodeAddr, ":")[0]
			bodyContent += " "

			bodyContent += node.NodeType
			bodyContent += " "

			bodyContent += node.NodeEnv
			bodyContent += " "

			bodyContent += node.DockerInfo
			bodyContent += " "

			bodyContent += node.HypervisorInfo
			bodyContent += " "

			bodyContent += strconv.Itoa(int(node.TotalCPU))
			bodyContent += " "

			bodyContent += strconv.Itoa(int(node.TotalMem))
			bodyContent += " "

			if node.NodeActive {
				bodyContent += "ACTIVE"
			} else {
				bodyContent += "DEAD"
			}
			bodyContent += " "

			bodyContent += "\n"
		}
		return c.String(http.StatusOK, bodyContent)
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
