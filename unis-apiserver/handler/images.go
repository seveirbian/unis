package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePrivateImages(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	if validateUser(username, password) {
		//get user's imagesInfo
		imagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPath + "/" + username + "/imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		var imagesInfo []ImageInfo
		err = json.Unmarshal(imagesInfoInJSON, &imagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}
		//generate response body
		var bodyContent = ""
		var blankLenth = 10
		for _, image := range imagesInfo {
			bodyContent += image.Repository
			bodyContent += EmptyString(strings.Count("Repository", "") +
				blankLenth - strings.Count(image.Repository, ""))

			bodyContent += image.Tag
			bodyContent += EmptyString(strings.Count("Tag", "") +
				blankLenth - strings.Count(image.Tag, ""))

			bodyContent += Substring(image.ImageID, 0, 10)
			bodyContent += EmptyString(strings.Count("Image ID", "") +
				blankLenth - strings.Count(Substring(image.ImageID, 0, 10), ""))

			bodyContent += image.Created
			bodyContent += EmptyString(strings.Count("Created", "") +
				blankLenth - strings.Count(image.Created, ""))

			bodyContent += image.Size
			bodyContent += EmptyString(strings.Count("Size", "") +
				blankLenth - strings.Count(image.Size, ""))

			bodyContent += image.Type
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(image.Type, ""))

			bodyContent += image.Owner
			bodyContent += EmptyString(strings.Count("Owner", "") +
				blankLenth - strings.Count(image.Owner, ""))

			bodyContent += "\n"
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePublicImages(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if validateUser(username, password) {
		//get public images info
		publicImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPublicPath + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		var publicImagesInfo []ImageInfo
		err = json.Unmarshal(publicImagesInfoInJSON, &publicImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}
		//get private images info
		privateImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPath + "/" + username + "/imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		var privateImagesInfo []ImageInfo
		err = json.Unmarshal(privateImagesInfoInJSON, &privateImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		//generate response body
		var bodyContent = ""
		var blankLenth = 10
		for _, image := range publicImagesInfo {
			bodyContent += image.Repository
			bodyContent += EmptyString(strings.Count("Repository", "") +
				blankLenth - strings.Count(image.Repository, ""))

			bodyContent += image.Tag
			bodyContent += EmptyString(strings.Count("Tag", "") +
				blankLenth - strings.Count(image.Tag, ""))

			bodyContent += Substring(image.ImageID, 0, 10)
			bodyContent += EmptyString(strings.Count("Image ID", "") +
				blankLenth - strings.Count(Substring(image.ImageID, 0, 10), ""))

			bodyContent += image.Created
			bodyContent += EmptyString(strings.Count("Created", "") +
				blankLenth - strings.Count(image.Created, ""))

			bodyContent += image.Size
			bodyContent += EmptyString(strings.Count("Size", "") +
				blankLenth - strings.Count(image.Size, ""))

			bodyContent += image.Type
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(image.Type, ""))

			bodyContent += image.Owner
			bodyContent += EmptyString(strings.Count("Owner", "") +
				blankLenth - strings.Count(image.Owner, ""))

			bodyContent += "\n"
		}
		for _, image := range privateImagesInfo {
			bodyContent += image.Repository
			bodyContent += EmptyString(strings.Count("Repository", "") +
				blankLenth - strings.Count(image.Repository, ""))

			bodyContent += image.Tag
			bodyContent += EmptyString(strings.Count("Tag", "") +
				blankLenth - strings.Count(image.Tag, ""))

			bodyContent += Substring(image.ImageID, 0, 10)
			bodyContent += EmptyString(strings.Count("Image ID", "") +
				blankLenth - strings.Count(Substring(image.ImageID, 0, 10), ""))

			bodyContent += image.Created
			bodyContent += EmptyString(strings.Count("Created", "") +
				blankLenth - strings.Count(image.Created, ""))

			bodyContent += image.Size
			bodyContent += EmptyString(strings.Count("Size", "") +
				blankLenth - strings.Count(image.Size, ""))

			bodyContent += image.Type
			bodyContent += EmptyString(strings.Count("Type", "") +
				blankLenth - strings.Count(image.Type, ""))

			bodyContent += image.Owner
			bodyContent += EmptyString(strings.Count("Owner", "") +
				blankLenth - strings.Count(image.Owner, ""))

			bodyContent += "\n"
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
