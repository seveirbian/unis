package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicLeave(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	nodename := c.Param("nodename")

	if validateUser(username, password) {
		publicNodesInfo := getPublicNodesInfo()
		for index, node := range publicNodesInfo {
			if node.NodeName == nodename {
				publicNodesInfo[index].NodeActive = false

				publicNodesInfoInJSON, err := json.Marshal(publicNodesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.NodesPublicPath+"nodesInfo.json", publicNodesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, "node leaved")
			}
		}
		return c.String(http.StatusNotImplemented, "node name error")
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePrivateLeave(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	nodename := c.Param("nodename")

	if validateUser(username, password) {
		privateNodesInfo := getPrivateNodesInfo(username)
		for index, node := range privateNodesInfo {
			if node.NodeName == nodename {
				privateNodesInfo[index].NodeActive = false

				privateNodesInfoInJSON, err := json.Marshal(privateNodesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.NodesPath+username+"/nodesInfo.json", privateNodesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, "node leaved")
			}
		}
		return c.String(http.StatusNotImplemented, "node name error")
	}

	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
