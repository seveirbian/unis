package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	// server.POST()
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

func handlePublicAdd(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	nodename := c.Param("nodename")
	nodeaddr := c.Request().RemoteAddr
	environment := c.FormValue("environment")
	dockerinfo := c.FormValue("dockerinfo")
	hypervisorinfo := c.FormValue("hypervisorinfo")

	totalcpu, _ := strconv.Atoi(c.FormValue("availablecpu"))
	totalmem, _ := strconv.Atoi(c.FormValue("availablemem"))

	if validateUser(username, password) {
		publicNodesInfo := getPublicNodesInfo()
		for _, node := range publicNodesInfo {
			// detect whether node name has existed
			if node.NodeName == nodename {
				return c.String(http.StatusConflict, "node name has existed")
			}
		}

		// create new node info
		newNode := NodeInfo{
			NodeName:       nodename,
			NodeType:       "public",
			NodeEnv:        environment,
			NodeAddr:       nodeaddr,
			DockerInfo:     dockerinfo,
			HypervisorInfo: hypervisorinfo,
			TotalCPU:       int64(totalcpu),
			TotalMem:       int64(totalmem),
		}

		// add new node into nodesInfo.json
		publicNodesInfo = append(publicNodesInfo, newNode)

		publicNodesInfoInJSON, err := json.Marshal(publicNodesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(serverFilePath.NodesPublicPath+"nodesInfo.json", publicNodesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		return c.String(http.StatusOK, "Node added")
	}

	return c.String(http.StatusForbidden, "incorrect username or password")
}

func handlePrivateAdd(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	nodename := c.Param("nodename")
	nodeaddr := c.Request().RemoteAddr
	environment := c.FormValue("environment")
	dockerinfo := c.FormValue("dockerinfo")
	hypervisorinfo := c.FormValue("hypervisorinfo")

	totalcpu, _ := strconv.Atoi(c.FormValue("availablecpu"))
	totalmem, _ := strconv.Atoi(c.FormValue("availablemem"))

	if validateUser(username, password) {
		privateNodesInfo := getPrivateNodesInfo(username)
		for _, node := range privateNodesInfo {
			// detect whether node name has existed
			if node.NodeName == nodename {
				return c.String(http.StatusConflict, "node name has existed")
			}
		}

		// create new node info
		newNode := NodeInfo{
			NodeName:       nodename,
			NodeType:       "private",
			NodeEnv:        environment,
			NodeAddr:       nodeaddr,
			DockerInfo:     dockerinfo,
			HypervisorInfo: hypervisorinfo,
			TotalCPU:       int64(totalcpu),
			TotalMem:       int64(totalmem),
		}

		// add new node into nodesInfo.json
		privateNodesInfo = append(privateNodesInfo, newNode)

		privateNodesInfoInJSON, err := json.Marshal(privateNodesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(serverFilePath.NodesPath+username+"/nodesInfo.json", privateNodesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		return c.String(http.StatusOK, "Node added")
	}

	return c.String(http.StatusForbidden, "incorrect username or password")
}
