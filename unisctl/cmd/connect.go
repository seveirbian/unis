package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		//get unis-apiserver's IP
		serverIP := args[0]
		if !strings.HasPrefix(serverIP, "http://") {
			serverIP = "http://" + serverIP
			if len(strings.Split(serverIP, ":")) == 2 {
				serverIP = serverIP + ":9898"
			}
		}

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
					logrus.Fatal(err)
				}
			}
		}

		//update configure.json
		ConfigContent.Apiserver = serverIP
		configInJSON, err := json.Marshal(ConfigContent)
		if err != nil {
			logrus.Fatal(err)
		}
		err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.SetUsageTemplate(connectUsage)
}
