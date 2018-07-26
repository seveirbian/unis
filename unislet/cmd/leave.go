package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var leaveUsage = `Usage:  unislet leave [OPTIONS]

Options:
  -p, --public public node leave unis (default private)
  -h, --help   help for leave
`

var leavePublicFlag bool

var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "leave unis",
	Long:  "leave unis",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		var resp *http.Response
		var err error

		if leavePublicFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/nodes/leave/public/"+ConfigContent.NodeName, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/nodes/leave/"+ConfigContent.Username+"/"+ConfigContent.NodeName, url.Values{"password": {ConfigContent.Password}})
		}

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
	rootCmd.AddCommand(leaveCmd)
	leaveCmd.SetUsageTemplate(leaveUsage)

	leaveCmd.Flags().BoolVarP(&leavePublicFlag, "public", "p", false, "node type (public or private)")
}
