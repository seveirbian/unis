package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var connectUsage = `Usage:  unisctl connect [OPTIONS] [SERVER]

Options:
  -h, --help   help for connect
`

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to unis-apiserver",
	Long:  "Connect to unis-apiserver",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithFields(logrus.Fields{
			"cmd": "unisctl connect",
		}).Info("")

		//get unis-apiserver's IP
		serverIP := args[0]
		if !strings.HasPrefix(serverIP, "http://") {
			serverIP = "http://" + serverIP
			if len(strings.Split(serverIP, ":")) == 2 {
				serverIP = serverIP + ":9898"
			}
		}
		logrus.Info("Connect to " + serverIP)

		//connect to unis-apiserver
		resp, err := http.Get(serverIP)
		if err != nil {
			logrus.Fatal(err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			} else {
				if string(body) == "OK" {
					fmt.Println("Success to connect unis-apiserver!")
				} else {
					logrus.Fatal("Getting bad response!")
				}
			}
		}

		//update configure.json
		configInJSON, err := ioutil.ReadFile(defaultPath + defaultFileName)
		var config Config
		err = json.Unmarshal(configInJSON, &config)
		if err != nil {
			logrus.Fatal("Failing decoding configure.json")
		}
		config.Apiserver = serverIP
		configInJSON, err = json.Marshal(config)
		if err != nil {
			logrus.Fatal("Failing encoding configure.json")
		}
		err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, 0777)
		if err != nil {
			logrus.Fatal("Failing writing into " + defaultPath + defaultFileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.SetUsageTemplate(connectUsage)
}
