package apiserver

import (
	"os"

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

	// serve "unisctl pull" command
	server.POST("/images/pull/public/:imageID", handlePublicPull)
	server.POST("/images/pull/:username/:imageID", handlePrivatePull)

	// // serve "unisctl run" command
	server.POST("/instances/run/public/:imageID", handlePublicRun)
	server.POST("/instances/run/:username/:imageID", handlePrivateRun)
	// serve "unisctl stop" command
	server.POST("/instances/stop/public/:instanceID", handlePublicStop)
	server.POST("/instances/stop/:username/:instanceID", handlePrivateStop)
	// // serve "unisctl ps" command
	// server.POST()

	// // serve "unisctl version" command
	// server.POST()
	// // serve "unistl nodes" command
	server.POST("/nodes/show/public/nodes", handlePublicNodes)
	server.POST("/nodes/show/:username/nodes", handlePrivateNodes)
	// serve "unisctl instances" command
	server.POST("/instances/show/all/instances", handleAllInstances)
	server.POST("/instances/show/:username/instances", handlePrivateInstances)

	// serve UNISLET
	// serve "unislet add" command
	server.POST("/nodes/add/public/:nodename", handlePublicAdd)
	server.POST("/nodes/add/:username/:nodename", handlePrivateAdd)

	// Run request-handler, this should never exit
	return server.Start(serveIP)
}
