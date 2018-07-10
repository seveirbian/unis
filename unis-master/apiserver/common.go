package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// validate user based on username and password
func validateUser(username string, password string) bool {
	for _, user := range UsersInfo {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// generate a substring
func Substring(str string, start int, end int) string {
	stringSlice := []rune(str)
	stringLen := len(stringSlice)

	if start < 0 || start >= stringLen {
		return ""
	} else if end < 0 || end >= stringLen || end < start {
		return ""
	} else {
		return string(stringSlice[start:end])
	}
}

// generate empty string
func EmptyString(length int) string {
	var stringSlice []rune

	for i := 0; i < length; i++ {
		stringSlice = append(stringSlice, rune(' '))
	}

	return string(stringSlice[:])
}

// get public images info
func getPublicImagesInfo() []ImageInfo {
	publicImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPublicPath + "imagesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}

	var publicImagesInfo []ImageInfo
	err = json.Unmarshal(publicImagesInfoInJSON, &publicImagesInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	return publicImagesInfo
}

func getPrivateImagesInfo(username string) []ImageInfo {
	privateImagesInfoInJSON, err := ioutil.ReadFile(serverFilePath.ImagesPath + "/" + username + "/imagesInfo.json")
	if err != nil {
		logrus.Fatal(err)
	}

	var privateImagesInfo []ImageInfo
	err = json.Unmarshal(privateImagesInfoInJSON, &privateImagesInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	return privateImagesInfo
}

// create the server file path
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

// load users info if users.json exists or create one
func (serverFilePath ServerFilePath) loadUsersJSON() error {
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
