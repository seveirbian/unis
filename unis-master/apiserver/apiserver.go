package apiserver

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
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

	// // serve "unisctl run" command
	// server.POST()
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

	// ABANDONED
	// serve "unisctl pull" command
	// server.POST("/images/pull/public/:imageID", handlePublicPull)
	// server.POST("/images/pull/:username/:imageID", handlePrivatePull)

	// Run request-handler, this should never exit
	return server.Start(serveIP)
}

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
