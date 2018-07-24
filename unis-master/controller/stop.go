package controller

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handleStopPublicInstance(c echo.Context) error {
	instanceID := c.Param("instanceID")
	nodename := c.FormValue("nodename")

	// load nodes
	err := loadNodes()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, node := range publicNodes {
		if strings.Contains(node.NodeName, nodename) {
			for _, instance := range node.Instances {
				if strings.Contains(instance.InstanceID, instanceID) {
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/stop/"+instanceID, url.Values{})
					if err != nil {
						return c.String(http.StatusNotImplemented, err.Error())
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail read resp.Body")
					}

					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusNotImplemented, string(body))
					}

					return c.String(http.StatusOK, string(body))
				}
			}
		}
	}

	return c.String(http.StatusNotImplemented, "node info is different between apiserver and controller")
}

func handleStopPrivateInstance(c echo.Context) error {
	username := c.Param("username")
	instanceID := c.Param("instanceID")
	nodename := c.FormValue("nodename")

	// load nodes
	err := loadNodes()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, node := range privateNodes[username] {
		if strings.Contains(node.NodeName, nodename) {
			for _, instance := range node.Instances {
				if strings.Contains(instance.InstanceID, instanceID) {
					resp, err := http.PostForm("http://"+strings.Split(node.NodeAddr, ":")[0]+":9899/instances/stop/"+instanceID, url.Values{})
					if err != nil {
						return c.String(http.StatusNotImplemented, err.Error())
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return c.String(http.StatusNotImplemented, "fail read resp.Body")
					}

					if resp.StatusCode != http.StatusOK {
						return c.String(http.StatusNotImplemented, string(body))
					}

					return c.String(http.StatusOK, string(body))
				}
			}
		}
	}

	return c.String(http.StatusNotImplemented, "node info is different between apiserver and controller")
}
