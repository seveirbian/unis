package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var signupUsage = `Usage:  unisctl signup [OPTIONS]

Options:
  -h, --help            help for signin
  -p, --password        password
  -u, --username        username
`

var suUsername, suPassword string

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Sign up",
	Long:  "Sign up",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.PostForm(ConfigContent.Apiserver+"/users/signup", url.Values{"username": {suUsername}, "password": {suPassword}})
		if err != nil {
			logrus.Fatal(err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
	signupCmd.SetUsageTemplate(signupUsage)
	signupCmd.Flags().StringVarP(&suUsername, "username", "u", "", "username used to sign up")
	signupCmd.Flags().StringVarP(&suPassword, "password", "p", "", "password used to sign up")
	signupCmd.MarkFlagRequired("username")
	signupCmd.MarkFlagRequired("password")
}
