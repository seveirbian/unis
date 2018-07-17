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
		privateNodesInfo := getPrivateNodesInfo(username)

		//generate response body
		var bodyContent = ""
		var blankLenth = 10
		for _, node := range publicNodesInfo {
			bodyContent += node.NodeName
			bodyContent += EmptyString(strings.Count("Name", "") +
				blankLenth - strings.Count(node.NodeName, ""))

			bodyContent += strings.Split(node.NodeAddr, ":")[0]
			bodyContent += EmptyString(strings.Count("Addr", "") +
				blankLenth - strings.Count(strings.Split(node.NodeAddr, ":")[0], ""))

			bodyContent += node.NodeType
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(node.NodeType, ""))

			bodyContent += node.NodeEnv
			bodyContent += EmptyString(strings.Count("Env", "") +
				blankLenth - strings.Count(node.NodeEnv, ""))

			bodyContent += node.DockerInfo
			bodyContent += EmptyString(strings.Count("Docker", "") +
				blankLenth - strings.Count(node.DockerInfo, ""))

			bodyContent += node.HypervisorInfo
			bodyContent += EmptyString(strings.Count("Hypervisor", "") +
				blankLenth - strings.Count(node.HypervisorInfo, ""))

			bodyContent += strconv.Itoa(int(node.TotalCPU))
			bodyContent += EmptyString(strings.Count("Avail CPU", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalCPU)), ""))

			bodyContent += strconv.Itoa(int(node.TotalMem))
			bodyContent += EmptyString(strings.Count("Avail Mem", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalMem)), ""))

			bodyContent += "\n"
		}
		for _, node := range privateNodesInfo {
			bodyContent += node.NodeName
			bodyContent += EmptyString(strings.Count("Name", "") +
				blankLenth - strings.Count(node.NodeName, ""))

			bodyContent += strings.Split(node.NodeAddr, ":")[0]
			bodyContent += EmptyString(strings.Count("Addr", "") +
				blankLenth - strings.Count(strings.Split(node.NodeAddr, ":")[0], ""))

			bodyContent += node.NodeType
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(node.NodeType, ""))

			bodyContent += node.NodeEnv
			bodyContent += EmptyString(strings.Count("Env", "") +
				blankLenth - strings.Count(node.NodeEnv, ""))

			bodyContent += node.DockerInfo
			bodyContent += EmptyString(strings.Count("Docker", "") +
				blankLenth - strings.Count(node.DockerInfo, ""))

			bodyContent += node.HypervisorInfo
			bodyContent += EmptyString(strings.Count("Hypervisor", "") +
				blankLenth - strings.Count(node.HypervisorInfo, ""))

			bodyContent += strconv.Itoa(int(node.TotalCPU))
			bodyContent += EmptyString(strings.Count("Avail CPU", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalCPU)), ""))

			bodyContent += strconv.Itoa(int(node.TotalMem))
			bodyContent += EmptyString(strings.Count("Avail Mem", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalMem)), ""))

			bodyContent += "\n"
		}
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
		var blankLenth = 10
		for _, node := range privateNodesInfo {
			bodyContent += node.NodeName
			bodyContent += EmptyString(strings.Count("Name", "") +
				blankLenth - strings.Count(node.NodeName, ""))

			bodyContent += strings.Split(node.NodeAddr, ":")[0]
			bodyContent += EmptyString(strings.Count("Addr", "") +
				blankLenth - strings.Count(strings.Split(node.NodeAddr, ":")[0], ""))

			bodyContent += node.NodeType
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(node.NodeType, ""))

			bodyContent += node.NodeEnv
			bodyContent += EmptyString(strings.Count("Env", "") +
				blankLenth - strings.Count(node.NodeEnv, ""))

			bodyContent += node.DockerInfo
			bodyContent += EmptyString(strings.Count("Docker", "") +
				blankLenth - strings.Count(node.DockerInfo, ""))

			bodyContent += node.HypervisorInfo
			bodyContent += EmptyString(strings.Count("Hypervisor", "") +
				blankLenth - strings.Count(node.HypervisorInfo, ""))

			bodyContent += strconv.Itoa(int(node.TotalCPU))
			bodyContent += EmptyString(strings.Count("Avail CPU", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalCPU)), ""))

			bodyContent += strconv.Itoa(int(node.TotalMem))
			bodyContent += EmptyString(strings.Count("Avail Mem", "") +
				blankLenth - strings.Count(strconv.Itoa(int(node.TotalMem)), ""))

			bodyContent += "\n"
		}
		return c.String(http.StatusOK, bodyContent)
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
