package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	// "fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Version string
}

type ServerFilePath struct {
	RootPath         string
	ImagesPath       string
	NodesPath        string
	ImagesPublicPath string
	NodesPublicPath  string
	UsersJSONPath    string
}

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ImageInfo struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
	Type       string
	Owner      string
}

var UsersInfo []userInfo

var serverFilePath = ServerFilePath{
	RootPath:         os.Getenv("HOME") + "/.unis/apiserver/",
	ImagesPath:       os.Getenv("HOME") + "/.unis/apiserver/images/",
	NodesPath:        os.Getenv("HOME") + "/.unis/apiserver/nodes/",
	ImagesPublicPath: os.Getenv("HOME") + "/.unis/apiserver/images/public/",
	NodesPublicPath:  os.Getenv("HOME") + "/.unis/apiserver/nodes/public/",
	UsersJSONPath:    os.Getenv("HOME") + "/.unis/apiserver/",
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	//create server dir
	serverFilePath.createFilePath()

	//read users.json
	serverFilePath.readUsersJSON()
	os.Stat(serverFilePath.UsersJSONPath)
}

func (serverFilePath ServerFilePath) createFilePath() error {
	_, err := os.Stat(serverFilePath.RootPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.RootPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	_, err = os.Stat(serverFilePath.ImagesPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.ImagesPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	//detect $HOME/.unis/images/public/
	_, err = os.Stat(serverFilePath.ImagesPublicPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.ImagesPublicPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	//detect $HOME/.unis/images/public/imagesInfo.json
	_, err = os.Stat(serverFilePath.ImagesPublicPath + "imagesInfo.json")
	if err != nil {
		_, err = os.Create(serverFilePath.ImagesPublicPath + "imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}
		publicImagesInfo := []ImageInfo{}
		publicImagesInfoInJSON, err := json.Marshal(publicImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}
		err = ioutil.WriteFile(serverFilePath.ImagesPublicPath+"imagesInfo.json", publicImagesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	_, err = os.Stat(serverFilePath.NodesPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.NodesPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	_, err = os.Stat(serverFilePath.NodesPublicPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.NodesPublicPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	return nil
}

func (serverFilePath ServerFilePath) readUsersJSON() error {
	_, err := os.Stat(serverFilePath.UsersJSONPath + "users.json")
	if err != nil {
		_, err = os.Create(serverFilePath.UsersJSONPath + "users.json")
		if err != nil {
			logrus.Fatal(err)
		}
		// UsersInfo = append(UsersInfo, userInfo{Username: "admin", Password: "admin"})
		UsersInfoInJSON, err := json.Marshal(UsersInfo)
		if err != nil {
			logrus.Fatal(err)
		}
		err = ioutil.WriteFile(serverFilePath.UsersJSONPath+"/users.json", UsersInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		UsersInfoInJSON, err := ioutil.ReadFile(serverFilePath.UsersJSONPath + "/users.json")
		if err != nil {
			logrus.Fatal(err)
		}
		err = json.Unmarshal(UsersInfoInJSON, &UsersInfo)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	return nil
}

func (rqHandler Handler) Serve(serveIP string) error {
	server := echo.New()

	//serve "unisctl connect" command
	server.GET("/", handleConnect)
	//serve "unisctl signin" command
	server.GET("/users.json/:username/:password", handleSignin)
	//serve "unisctl signup" command
	server.POST("/users.json", handleSignup)
	//serve "unisctl images" command
	server.POST("/images/:username/images", handlePrivateImages)
	server.POST("/images/public/images", handlePublicImages)
	//serve "unisctl push" command
	server.POST("/images/public/:imagename", handlePublicPush)
	server.POST("/images/:username/:imagename", handlePrivatePush)

	// Run request-handler, this should never exit
	return server.Start(serveIP)
}

func handleConnect(c echo.Context) error {
	logrus.Info("Handling connect request from " + c.RealIP())
	return c.String(http.StatusOK, "OK")
}

func handleSignin(c echo.Context) error {
	username := c.Param("username")
	password := c.Param("password")
	if validateUser(username, password) {
		return c.String(http.StatusOK, "Signin succeeded")
	} else {
		return c.String(http.StatusUnauthorized, "incorrect username or password")
	}
}

func handleSignup(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	//detect whether username has existed
	for _, user := range UsersInfo {
		if user.Username == username {
			return c.String(http.StatusConflict, "Username has existed")
		}
	}
	//add new account into user.json
	UsersInfo = append(UsersInfo, userInfo{Username: username, Password: password})
	UsersInfoInJSON, err := json.Marshal(UsersInfo)
	if err != nil {
		logrus.Fatal(err)
	}
	err = ioutil.WriteFile(serverFilePath.UsersJSONPath+"/users.json", UsersInfoInJSON, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
	//create new account folder and info file
	err = os.Mkdir(serverFilePath.ImagesPath+"/"+username, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
	imagesInfo := []ImageInfo{}
	imagesInfoInJSON, err := json.Marshal(imagesInfo)
	if err != nil {
		logrus.Fatal(err)
	}
	os.Create(serverFilePath.ImagesPath + "/" + username + "/" + "imagesInfo.json")
	err = ioutil.WriteFile(serverFilePath.ImagesPath+"/"+username+"/"+"imagesInfo.json", imagesInfoInJSON, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}

	return c.String(http.StatusOK, "New account has created")
}

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

		//make change to imagesInfo.json
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

		//transimit image
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

		//detect whether image has existed
		for _, imageInfo := range imagesInfo {
			if imageInfo.Repository == (repository) && imageInfo.Tag == tag {
				return c.String(http.StatusForbidden, "image already exists")
			}
		}

		//make change to imagesInfo.json
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

		//transimit image
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

		return c.String(http.StatusOK, "image pushed")
	} else {
		return c.String(http.StatusUnauthorized, "incorrect username or password")
	}
}
