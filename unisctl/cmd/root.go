package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var defaultPath = os.Getenv("HOME") + "/.unis/"
var defaultFileName = "configure.json"

type Config struct {
	Apiserver string `json:"apiserver"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

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
	logrus.WithFields(logrus.Fields{
		"cmd": "unisctl ",
	}).Info("")

	//detect whether /home/.unis/ exists, if not create it
	createPath()

	// create configure.json
	_, err := ioutil.ReadFile(defaultPath + defaultFileName)
	if err != nil {
		logrus.Info("Failing finding " + defaultPath + defaultFileName)
		var config = Config{
			Apiserver: "",
			Username:  "",
			Password:  "",
		}

		configInJSON, err := json.Marshal(config)
		if err != nil {
			logrus.Fatal("Failing encoding config")
			fmt.Println(err)
		}

		err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, 0777)
		if err != nil {
			fmt.Println(err)
			logrus.Fatal("Failing to wirte configure into " + defaultPath + defaultFileName)
			fmt.Println(err)
		} else {
			logrus.Info("Creating " + defaultPath + defaultFileName)
		}
	} else {
		logrus.Info("Reading from " + defaultPath + defaultFileName)
	}
}

func createPath() {
	_, err := os.Stat(defaultPath)
	if err != nil {
		err = os.Mkdir(defaultPath, 0777)
		if err != nil {
			fmt.Println(err)
			logrus.Fatal("Failing creating dir: " + defaultPath)
		}
		logrus.Info("Creating dir: " + defaultPath)
	}
}
