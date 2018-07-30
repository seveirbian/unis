package apiserver

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func handlePublicImagesNum(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		publicImages := getPublicImagesInfo()

		return c.String(http.StatusOK, strconv.Itoa(len(publicImages)))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}

func handlePrivateImagesNum(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		privateImages := getPrivateImagesInfo(username)

		return c.String(http.StatusOK, strconv.Itoa(len(privateImages)))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}

func handlePublicNodesNum(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		publicNodes := getPublicNodesInfo()

		return c.String(http.StatusOK, strconv.Itoa(len(publicNodes)))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}

func handlePrivateNodesNum(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		privateNodes := getPrivateNodesInfo(username)

		return c.String(http.StatusOK, strconv.Itoa(len(privateNodes)))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}

func handlePublicInstancesNum(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	instancesNum := 0

	if validateUser(username, password) {
		publicNodes := getPublicNodesInfo()

		for _, node := range publicNodes {
			instancesNum += len(node.Instances)
		}

		return c.String(http.StatusOK, strconv.Itoa(instancesNum))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}

func handlePrivateInstancesNum(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")

	instanceNum := 0

	if validateUser(username, password) {
		privateNodes := getPrivateNodesInfo(username)

		for _, node := range privateNodes {
			instanceNum += len(node.Instances)
		}

		return c.String(http.StatusOK, strconv.Itoa(instanceNum))
	}

	return c.String(http.StatusNotImplemented, "incorrect username or password")
}
