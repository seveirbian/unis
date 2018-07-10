package apiserver

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ABANDONED

func handlePublicPull(c echo.Context) error {
	unisctlAddr := c.FormValue("unisctlAddr")
	imageID := c.Param("imageID")

	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		publicImagesInfo := getPublicImagesInfo()

		for _, image := range publicImagesInfo {
			if image.ImageID == imageID {
				arg0 := "curl"
				arg1 := "-F"
				arg2 := image.Repository + ":" + image.Tag + "=@" + serverFilePath.ImagesPublicPath + image.ImageID
				arg3 := unisctlAddr + "/pull/"

				child := exec.Command(arg0, arg1, arg2, arg3)
				output, err := child.Output()
				if err != nil {
					logrus.Fatal(err)
				}

				fmt.Println(output)

				return c.String(http.StatusOK, "image pulled")
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID")
}
