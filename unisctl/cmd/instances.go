package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var instancesUsage = `Usage:  unisctl instances [OPTIONS]

Options:
  -a, --all             show all instances
  -h, --help            help for instances
`

var instancesPublicFlag bool

var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Display the status of all components of unis",
	Long:  "Display the status of all components of unis",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// detect whether user has signed in
		if ConfigContent.Username == "" || ConfigContent.Password == "" {
			fmt.Println("Please signin first!")
			logrus.Fatal("Please signin first!")
		}

		// send different request based on allImagesFlag
		var resp *http.Response
		var err error
		if instancesPublicFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/instances/show/all/instances", url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/instances/show/"+ConfigContent.Username+"/instances", url.Values{"password": {ConfigContent.Password}})
		}

		// print response
		if err != nil {
			logrus.Fatal(err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			} else {
				var label = ""
				label += "NODE          "
				label += "REPOSITORY          "
				label += "TAG          "
				label += "IMAGE ID          "
				label += "INSTANCE ID          "
				label += "REQUEST CPU          "
				label += "REQUEST MEM          "
				fmt.Println(label)
				fmt.Println(string(body))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(instancesCmd)
	instancesCmd.SetUsageTemplate(instancesUsage)

	instancesCmd.Flags().BoolVarP(&instancesPublicFlag, "all", "a", false, "show all instances")
}
