package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

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
