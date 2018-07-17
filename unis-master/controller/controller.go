package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	// handle add node
	// cServer.POST("/nodes/add/:nodename", handleAddNode)
	// // handle remove node
	// cServer.POST("/nodes/remove/:nodename", handleRemoveNode)
	// // handle run instance
	// cServer.POST("/instances/run/:instancename", handleRunInstance)
	// // handle stop instance
	// cServer.POST("/instance/stop/:instancename", handleStopInstance)

	return cServer.Start(serveIP)
}

func loadNodes() error {
	// load public nodes info
	_, err := os.Stat(os.Getenv("HOME") + "/.unis/apiserver/nodes/public/nodesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}
	publicNodesInJSON, err := ioutil.ReadFile(os.Getenv("HOME") + "/.unis/apiserver/nodes/public/nodesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}
	err = json.Unmarshal(publicNodesInJSON, &publicNodes)
	if err != nil {
		logrus.Fatal(err)
	}

	// load private nodes info
	// get all users info
	_, err = os.Stat(os.Getenv("HOME") + "/.unis/apiserver/users.json")
	if err != nil {
		logrus.Fatal(err)
	}
	usersInfoInJSON, err := ioutil.ReadFile(os.Getenv("HOME") + "/.unis/apiserver/users.json")
	if err != nil {
		logrus.Fatal(err)
	}
	err = json.Unmarshal(usersInfoInJSON, &usersInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, user := range usersInfo {
		// detect whether user has private nodes
		f, err := os.Stat(os.Getenv("HOME") + "/.unis/apiserver/nodes/" + user.Username)
		if err != nil {
			logrus.Fatal(err)
		}
		if f.IsDir() {
			_, err = os.Stat(os.Getenv("HOME") + "/.unis/apiserver/nodes/" + user.Username + "/nodesInfo.json")
			if err != nil {
				logrus.Fatal(err)
			}
			nodesInfoInJSON, err := ioutil.ReadFile(os.Getenv("HOME") + "/.unis/apiserver/nodes/" + user.Username + "/nodesInfo.json")
			if err != nil {
				logrus.Fatal(err)
			}
			nodesInfo := []NodeInfo{}
			err = json.Unmarshal(nodesInfoInJSON, &nodesInfo)
			if err != nil {
				logrus.Fatal(err)
			}

			privateNodes[user.Username] = nodesInfo

		}
	}

	logrus.Info(publicNodes)
	logrus.Info(privateNodes)

	return nil
}
