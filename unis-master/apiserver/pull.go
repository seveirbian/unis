package apiserver

import (
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicPull(c echo.Context) error {
	unisctlAddr := strings.Split(c.Request().RemoteAddr, ":")[0] + ":10001"
	imageID := c.Param("imageID")

	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		publicImagesInfo := getPublicImagesInfo()

		for _, image := range publicImagesInfo {
			if strings.Contains(image.ImageID, imageID) {
				arg0 := "curl"
				arg1 := "-F"
				arg2 := strings.Split(image.Repository, "/")[1] + "=@" + serverFilePath.ImagesPublicPath + image.ImageID
				arg3 := "http://" + unisctlAddr + "/images/pull/" + strings.Split(image.Repository, "/")[1]

				time.Sleep(2000)

				child := exec.Command(arg0, arg1, arg2, arg3)
				_, err := child.Output()
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, "image pulled")
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID")
}
