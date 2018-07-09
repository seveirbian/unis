package handler

import (
	"encoding/json"
	"io/ioutil"
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

func (rqHandler Handler) Serve(serveIP string) error {
	server := echo.New()

	//serve "unisctl connect" command
	server.GET("/", handleConnect)
	//serve "unisctl signin" command
	server.POST("/users/signin", handleSignin)
	//serve "unisctl signup" command
	server.POST("/users/signup", handleSignup)
	//serve "unisctl images" command
	server.POST("/images/show/:username/images", handlePrivateImages)
	server.POST("/images/show/public/images", handlePublicImages)
	//serve "unisctl push" command
	server.POST("/images/push/public/:imagename", handlePublicPush)
	server.POST("/images/push/:username/:imagename", handlePrivatePush)
	//serve "unisctl rmi" command
	server.POST("/images/remove/public/:imageID", handlePublicRmi)
	server.POST("/images/remove/:username/:imageID", handlePrivateRmi)
	// //serve "unisctl tag" command
	// server.POST("/images/public/:oldimage/:newimage", handleTag)
	// //serve "unisctl run" command
	// server.POST()
	// //serve "unisctl stop" command
	// server.POST()
	// //serve "unisctl stats" command
	// server.POST()
	// //serve "unisctl ps" command
	// server.POST()
	// //serve "unisctl pull" command
	// server.POST()
	// //serve "unisctl version" command
	// server.POST()
	// //serve "unistl nodes" command
	// server.POST()

	// Run request-handler, this should never exit
	return server.Start(serveIP)
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
