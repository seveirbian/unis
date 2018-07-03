package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var signinUsage = `Usage:  unisctl signin [OPTIONS]

Options:
  -h, --help            help for signin
  -p, --password        password
  -u, --username        username
`

var siUsername, siPassword string

var signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "Sign in",
	Long:  "Sign in",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.PostForm(ConfigContent.Apiserver+"/users.json", url.Values{"username": {siUsername}, "password": {siPassword}})
		if err != nil {
			logrus.Fatal(err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			} else {
				fmt.Println(string(body))
				if resp.StatusCode == 200 {
					ConfigContent.Username = siUsername
					ConfigContent.Password = siPassword
					//write ConfigContent into configure.json
					configInJSON, err := json.Marshal(ConfigContent)
					if err != nil {
						logrus.Fatal(err)
					}
					err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
					if err != nil {
						logrus.Fatal(err)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(signinCmd)
	signinCmd.SetUsageTemplate(signinUsage)
	signinCmd.Flags().StringVarP(&siUsername, "username", "u", "", "username used to sign in")
	signinCmd.Flags().StringVarP(&siPassword, "password", "p", "", "password used to sign in")
	signinCmd.MarkFlagRequired("username")
	signinCmd.MarkFlagRequired("password")
}
