package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicTag(c echo.Context) error {
	oldImageName := c.Param("oldimage")
	oldTag := c.Param("oldtag")
	newImageName := c.Param("newimage")
	newTag := c.Param("newtag")

	username := c.FormValue("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		// get public images info
		publicImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPublicPath + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		var publicImagesInfo []ImageInfo
		err = json.Unmarshal(publicImagesInfoInJSON, &publicImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		// change images info
		for index, image := range publicImagesInfo {
			if strings.Split(image.Repository, "/")[1] == oldImageName && image.Tag == oldTag && image.Owner == username {
				for _, imageRef := range publicImagesInfo {
					if strings.Split(imageRef.Repository, "/")[1] == newImageName && imageRef.Tag == newTag {
						return c.String(http.StatusForbidden, "TARGET_SOURCE:[TAG] has existed")
					}
				}

				publicImagesInfo[index].Repository = "public/" + newImageName
				publicImagesInfo[index].Tag = newTag

				// write public images info back
				publicImagesInfoInJSON, err = json.Marshal(publicImagesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.ImagesPublicPath+"imagesInfo.json", publicImagesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, "Tag succeeded")
			}
		}
	}
	return c.String(http.StatusForbidden, "Wrong IMAGE:[TAG] or You are unauthorized")
}

func handlePrivateTag(c echo.Context) error {
	oldImageName := c.Param("oldimage")
	oldTag := c.Param("oldtag")
	newImageName := c.Param("newimage")
	newTag := c.Param("newtag")

	username := c.Param("username")
	password := c.FormValue("password")

	if validateUser(username, password) {
		// get private images info
		privateImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPath + username + "/" + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		var privateImagesInfo []ImageInfo
		err = json.Unmarshal(privateImagesInfoInJSON, &privateImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		// change images info
		for index, image := range privateImagesInfo {
			if strings.Split(image.Repository, "/")[1] == oldImageName && image.Tag == oldTag && image.Owner == username {
				for _, imageRef := range privateImagesInfo {
					if strings.Split(imageRef.Repository, "/")[1] == newImageName && imageRef.Tag == newTag {
						return c.String(http.StatusForbidden, "TARGET_SOURCE:[TAG] has existed")
					}
				}

				privateImagesInfo[index].Repository = username + "/" + newImageName
				privateImagesInfo[index].Tag = newTag

				// write private images info back
				privateImagesInfoInJSON, err = json.Marshal(privateImagesInfo)
				if err != nil {
					logrus.Fatal(err)
				}

				err = ioutil.WriteFile(serverFilePath.ImagesPath+username+"/"+"imagesInfo.json", privateImagesInfoInJSON, os.ModePerm)
				if err != nil {
					logrus.Fatal(err)
				}

				return c.String(http.StatusOK, "Tag succeeded")
			}
		}
	}
	return c.String(http.StatusForbidden, "Wrong IMAGE:[TAG] or You are unauthorized")
}
