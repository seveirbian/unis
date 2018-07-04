package handler

import (
	"encoding/json"
	"io/ioutil"

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

	_, err = os.Stat(serverFilePath.ImagesPublicPath)
	if err != nil {
		err = os.Mkdir(serverFilePath.ImagesPublicPath, os.ModePerm)
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
	// server.POST("/images/public/images", handlePublicImages)

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
	for _, user := range UsersInfo {
		if user.Username == username && user.Password == password {
			return c.String(http.StatusOK, "Signin succeeded")
		}
	}
	return c.String(http.StatusUnauthorized, "incorrect username of password")
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
	os.Create(serverFilePath.ImagesPath + "/" + username + "/" + username + ".json")
	err = ioutil.WriteFile(serverFilePath.ImagesPath+"/"+username+"/"+username+".json", imagesInfoInJSON, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}

	return c.String(http.StatusOK, "New account has created")
}

func handlePrivateImages(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	for _, user := range UsersInfo {
		if user.Username == username && user.Password == password {
			return c.String(http.StatusOK, "")
		}
	}
	return c.String(http.StatusUnauthorized, "incorrect username of password")
}
