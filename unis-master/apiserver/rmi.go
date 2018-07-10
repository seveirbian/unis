package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicRmi(c echo.Context) error {
	imageID := c.Param("imageID")
	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		// get public images info
		publicImagesInfo := getPublicImagesInfo()

		// get the image that to be removed
		for index, imageinfo := range publicImagesInfo {
			if strings.Contains(imageinfo.ImageID, imageID) && imageinfo.Owner == username {
				// delete the image file
				err := os.Remove(serverFilePath.ImagesPublicPath + imageinfo.ImageID)
				if err != nil {
					logrus.Fatal(err)
				}
				// delete the image info
				publicImagesInfo = append(publicImagesInfo[:index], publicImagesInfo[index+1:]...)
				// write public images info back
				publicImagesInfoInJSON, err := json.Marshal(publicImagesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.ImagesPublicPath+"imagesInfo.json", publicImagesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, imageID+" removed")
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID or You are unauthorized")
}

func handlePrivateRmi(c echo.Context) error {
	imageID := c.Param("imageID")
	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		// get private images info
		privateImagesInfo := getPrivateImagesInfo(username)

		// get the image that to be removed
		for index, imageinfo := range privateImagesInfo {
			if strings.Contains(imageinfo.ImageID, imageID) && imageinfo.Owner == username {
				// delete the image file
				err := os.Remove(serverFilePath.ImagesPath + username + "/" + imageinfo.ImageID)
				if err != nil {
					logrus.Fatal(err)
				}
				// delete the image info
				privateImagesInfo = append(privateImagesInfo[:index], privateImagesInfo[index+1:]...)
				// write private images info back
				privateImagesInfoInJSON, err := json.Marshal(privateImagesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.ImagesPath+username+"/"+"imagesInfo.json", privateImagesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, imageID+" removed")
			}
		}
	}
	return c.String(http.StatusNotFound, "Wrong imageID or You are unauthorized")
}
