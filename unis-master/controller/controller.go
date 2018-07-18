package controller

import (
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
	InstanceID      string

	RequestCPU int64
	RequestMem int64
	LimitCPU   int64
	LimitMem   int64
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
	err := loadNodes()
	if err != nil {
		logrus.Fatal(err)
	}
}

func (controller Controller) Start(serveIP string) error {
	controller.Version = "0.0001.0.0"

	cServer := echo.New()

	// handle test
	cServer.GET("/", handleConnect)
	// handle add user
	cServer.POST("/users/add/:username", handleAddUser)
	// handle add node
	cServer.POST("/nodes/add/public/:nodename", handleAddPublicNode)
	cServer.POST("/nodes/add/:username/:nodename", handleAddPrivateNode)
	// // handle remove node
	// cServer.POST("/nodes/remove/:nodename", handleRemoveNode)
	// // handle run instance
	cServer.POST("/instances/run/:instancename", handleRunInstance)
	// // handle stop instance
	// cServer.POST("/instance/stop/:instancename", handleStopInstance)

	return cServer.Start(serveIP)
}

func handleRunInstance(c echo.Context) error {
	return nil
}
