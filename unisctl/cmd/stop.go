package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var stopUsage = `Usage:  unisctl stop [OPTIONS] INSTANCE [INSTANCE...]

Options:
  -h, --help            help for stop
  -p, --public          stop public instance
`

var stopPublicFlag bool

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or more running instances",
	Long:  "Stop one or more running instances",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if stopPublicFlag {
			for _, instance := range args {
				resp, err := http.PostForm(ConfigContent.Apiserver+"/instances/stop/public/"+instance, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
				if err != nil {
					logrus.Fatal(err)
				} else {
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						logrus.Fatal(err)
					}

					fmt.Println(string(body))
				}
			}
		} else {
			for _, instance := range args {
				resp, err := http.PostForm(ConfigContent.Apiserver+"/instances/stop/"+ConfigContent.Username+"/"+instance, url.Values{"password": {ConfigContent.Password}})
				if err != nil {
					logrus.Fatal(err)
				} else {
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						logrus.Fatal(err)
					}

					fmt.Println(string(body))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.SetUsageTemplate(stopUsage)

	stopCmd.Flags().BoolVarP(&stopPublicFlag, "public", "p", false, "Stop one or more running instances")
}
