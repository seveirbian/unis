package apiserver

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func handlePublicPull(c echo.Context) error {
	imageID := c.Param("imageID")

	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		publicImagesInfo := getPublicImagesInfo()

		for _, image := range publicImagesInfo {
			if strings.Contains(image.ImageID, imageID) {

				return c.Attachment(serverFilePath.ImagesPublicPath+image.ImageID, strings.Split(image.Repository, "/")[1])
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID")
}

func handlePrivatePull(c echo.Context) error {
	imageID := c.Param("imageID")

	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		privateImagesInfo := getPrivateImagesInfo(username)

		for _, image := range privateImagesInfo {
			if strings.Contains(image.ImageID, imageID) {
				return c.Attachment(serverFilePath.ImagesPath+username+"/"+image.ImageID, strings.Split(image.Repository, "/")[1])
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID")
}
