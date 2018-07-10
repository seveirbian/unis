package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

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
