package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

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
		for index, node := range publicNodesInfo {
			// detect whether node name has existed
			if node.NodeName == nodename {
				if strings.Split(node.NodeAddr, ":")[0] == strings.Split(nodeaddr, ":")[0] {
					publicNodesInfo[index].NodeActive = true

					// public node reactive
					publicNodesInfoInJSON, err := json.Marshal(publicNodesInfo)
					if err != nil {
						logrus.Fatal(err)
					}

					err = ioutil.WriteFile(serverFilePath.NodesPublicPath+"nodesInfo.json", publicNodesInfoInJSON, os.ModePerm)
					if err != nil {
						logrus.Fatal(err)
					}

					return c.String(http.StatusOK, "node reactive")
				}
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
			NodeActive:     true,
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
	username := c.Param("username")
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
		for index, node := range privateNodesInfo {
			// detect whether node name has existed
			if node.NodeName == nodename {
				if strings.Split(node.NodeAddr, ":")[0] == strings.Split(nodeaddr, ":")[0] {
					privateNodesInfo[index].NodeActive = true

					// private node reactive
					privateNodesInfoInJSON, err := json.Marshal(privateNodesInfo)
					if err != nil {
						logrus.Fatal(err)
					}

					err = ioutil.WriteFile(serverFilePath.NodesPath+username+"/nodesInfo.json", privateNodesInfoInJSON, os.ModePerm)
					if err != nil {
						logrus.Fatal(err)
					}

					return c.String(http.StatusOK, "node reactive")
				}
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
			NodeActive:     true,
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
