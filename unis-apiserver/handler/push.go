package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicPush(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	imagename := c.Param("imagename")

	repository := c.FormValue("repository") + "/" + imagename
	tag := c.FormValue("tag")
	imageID := c.FormValue("imageID")
	created := c.FormValue("created")
	size := c.FormValue("size")
	imageType := c.FormValue("imageType")
	owner := c.FormValue("owner")

	if validateUser(username, password) {
		var imagesInfo []ImageInfo
		imagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPublicPath + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}

		err = json.Unmarshal(imagesInfoInJSON, &imagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		//detect whether image has existed
		for _, imageInfo := range imagesInfo {
			if imageInfo.Repository == (repository) && imageInfo.Tag == tag {
				return c.String(http.StatusForbidden, "image already exists")
			}
		}

		// transimit image
		imagefile, err := c.FormFile(imagename)
		if err != nil {
			logrus.Fatal(err)
		}

		src, err := imagefile.Open()
		if err != nil {
			logrus.Fatal(err)
		}
		defer src.Close()

		dst, err := os.Create(serverFilePath.ImagesPublicPath + imagename)
		if err != nil {
			logrus.Fatal(err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			logrus.Fatal(err)
		}

		// use imageID as file name
		err = os.Rename(serverFilePath.ImagesPublicPath+imagename, serverFilePath.ImagesPublicPath+imageID)
		if err != nil {
			logrus.Fatal(err)
		}

		// make change to imagesInfo.json
		imagesInfo = append(imagesInfo, ImageInfo{
			Repository: repository,
			Tag:        tag,
			ImageID:    imageID,
			Created:    created,
			Size:       size,
			Type:       imageType,
			Owner:      owner,
		})

		imagesInfoInJSON, err = json.Marshal(imagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(serverFilePath.ImagesPublicPath+"imagesInfo.json", imagesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		return c.String(http.StatusOK, "image pushed")
	} else {
		return c.String(http.StatusUnauthorized, "incorrect username or password")
	}
}

func handlePrivatePush(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	imagename := c.Param("imagename")

	repository := c.FormValue("repository") + "/" + imagename
	tag := c.FormValue("tag")
	imageID := c.FormValue("imageID")
	created := c.FormValue("created")
	size := c.FormValue("size")
	imageType := c.FormValue("imageType")
	owner := c.FormValue("owner")

	if validateUser(username, password) {
		var imagesInfo []ImageInfo
		imagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPath + username + "/" + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}

		err = json.Unmarshal(imagesInfoInJSON, &imagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		// detect whether image has existed
		for _, imageInfo := range imagesInfo {
			if imageInfo.Repository == (repository) && imageInfo.Tag == tag {
				return c.String(http.StatusForbidden, "image already exists")
			}
		}

		// transimit image
		imagefile, err := c.FormFile(imagename)
		if err != nil {
			logrus.Fatal(err)
		}

		src, err := imagefile.Open()
		if err != nil {
			logrus.Fatal(err)
		}
		defer src.Close()

		dst, err := os.Create(serverFilePath.ImagesPath + username + "/" + imagename)
		if err != nil {
			logrus.Fatal(err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			logrus.Fatal(err)
		}

		// use imageID as file name
		err = os.Rename(serverFilePath.ImagesPath+username+"/"+imagename, serverFilePath.ImagesPath+username+"/"+imageID)
		if err != nil {
			logrus.Fatal(err)
		}

		// make change to imagesInfo.json
		imagesInfo = append(imagesInfo, ImageInfo{
			Repository: repository,
			Tag:        tag,
			ImageID:    imageID,
			Created:    created,
			Size:       size,
			Type:       imageType,
			Owner:      owner,
		})

		imagesInfoInJSON, err = json.Marshal(imagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(serverFilePath.ImagesPath+username+"/"+"imagesInfo.json", imagesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		return c.String(http.StatusOK, "image pushed")
	} else {
		return c.String(http.StatusUnauthorized, "incorrect username or password")
	}
}
