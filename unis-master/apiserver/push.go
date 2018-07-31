package apiserver

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePublicPush(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	imagename := c.Param("imagename")
	tag := c.FormValue("tag")
	imageType := c.FormValue("imageType")

	repository := "public" + "/" + imagename
	var imageID string
	var created string
	var size string
	owner := username

	if validateUser(username, password) {
		imagesInfo := getPublicImagesInfo()

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

		// move from unisctl to here
		//get image size
		fileInfo, err := os.Stat(serverFilePath.ImagesPublicPath + imagename)
		if err != nil {
			logrus.Fatal(err)
		}
		size = strconv.FormatInt(fileInfo.Size()/1024/1024, 10)

		created = string(time.Now().Format("2006-01-02"))
		i := string(time.Now().Format("2006-01-02T15:04:05Z"))

		if content, err := ioutil.ReadFile(serverFilePath.ImagesPublicPath + imagename); err != nil {
			logrus.Fatal(err)
		} else {
			temp := sha256.Sum256([]byte(string(content) + i))
			imageID = hex.EncodeToString(temp[:])
		}
		// move from unisctl to here

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

		imagesInfoInJSON, err := json.Marshal(imagesInfo)
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
	tag := c.FormValue("tag")
	imageType := c.FormValue("imageType")

	repository := username + "/" + imagename

	var imageID string
	var created string
	var size string
	owner := username

	if validateUser(username, password) {
		imagesInfo := getPrivateImagesInfo(username)

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

		// move from unisctl to here
		//get image size
		fileInfo, err := os.Stat(serverFilePath.ImagesPath + username + "/" + imagename)
		if err != nil {
			logrus.Fatal(err)
		}
		size = strconv.FormatInt(fileInfo.Size()/1024/1024, 10)

		created = string(time.Now().Format("2006-01-02"))
		i := string(time.Now().Format("2006-01-02T15:04:05Z"))

		if content, err := ioutil.ReadFile(serverFilePath.ImagesPath + username + "/" + imagename); err != nil {
			logrus.Fatal(err)
		} else {
			temp := sha256.Sum256([]byte(string(content) + i))
			imageID = hex.EncodeToString(temp[:])
		}
		// move from unisctl to here

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

		imagesInfoInJSON, err := json.Marshal(imagesInfo)
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
