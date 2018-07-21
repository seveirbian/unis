package scheduler

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func GetPublicNodesInfo() []NodeInfo {
	publicNodesInfoInJSON, err := ioutil.ReadFile(os.Getenv("HOME") + "/.unis/apiserver/nodes/public/nodesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}

	var publicNodesInfo []NodeInfo
	err = json.Unmarshal(publicNodesInfoInJSON, &publicNodesInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	return publicNodesInfo
}

func GetPrivateNodesInfo(username string) []NodeInfo {
	privateNodesInfoInJSON, err := ioutil.ReadFile(os.Getenv("HOME") + "/.unis/apiserver/nodes/" + username + "/nodesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}

	var privateNodesInfo []NodeInfo
	err = json.Unmarshal(privateNodesInfoInJSON, &privateNodesInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	return privateNodesInfo
}
