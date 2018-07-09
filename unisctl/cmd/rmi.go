package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rmiUsage = `Usage:  unisctl rmi [OPTIONS] IMAGE [IMAGE...]

Options:
  -h, --help            help for rmi
  -p, --public          Remove public images (default private images)
`

var rmiPublicFlag bool

var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "reomve one or more images in registry",
	Long:  "remobe one or more images in registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if rmiPublicFlag {
			for _, imageID := range args {
				resp, err := http.PostForm(ConfigContent.Apiserver+"/images/remove/public/"+imageID, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
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
			}
		} else {
			for _, imageID := range args {
				resp, err := http.PostForm(ConfigContent.Apiserver+"/images/remove/"+ConfigContent.Username+"/"+imageID, url.Values{"password": {ConfigContent.Password}})
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
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmiCmd)
	rmiCmd.SetUsageTemplate(rmiUsage)
	rmiCmd.Flags().BoolVarP(&rmiPublicFlag, "public", "p", false, "Remove a public image")
}
