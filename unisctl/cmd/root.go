package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var defaultPath = os.Getenv("HOME") + "/.unis/unisctl/"
var defaultFileName = "configure.json"

type Config struct {
	Apiserver string `json:"apiserver"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

var ConfigContent Config

var rootCmd = &cobra.Command{
	Use:   "unisctl",
	Short: "A client to communicate with unis-apiserver",
	Long:  "A client to communicate with unis-apiserver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	//detect whether /home/.unis/unisctl exists, if not create it
	createPath()

	// create configure.json
	_, err := ioutil.ReadFile(defaultPath + defaultFileName)
	if err != nil {
		ConfigContent.Apiserver = ""
		ConfigContent.Username = ""
		ConfigContent.Password = ""

		configInJSON, err := json.Marshal(ConfigContent)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		configInJSON, err := ioutil.ReadFile(defaultPath + defaultFileName)
		err = json.Unmarshal(configInJSON, &ConfigContent)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func createPath() {
	_, err := os.Stat(defaultPath)
	if err != nil {
		err = os.Mkdir(defaultPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			logrus.Fatal(err)
		}
	}
}
