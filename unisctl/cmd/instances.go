package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var instancesUsage = `Usage:  unisctl instances [OPTIONS]

Options:
  -p, --public          show public instances
  -h, --help            help for instances
`

var instancesPublicFlag bool

var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Show instances",
	Long:  "Show instances",
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
				label += "\n"

				imageLine := strings.Split(string(body), "\n")
				blankLenth := 10

				for _, image := range imageLine {
					if image == "" {
						break
					}

					messages := strings.Split(image, " ")

					label += messages[0]
					label += EmptyString(strings.Count("Node", "") +
						blankLenth - strings.Count(messages[0], ""))

					label += messages[1]
					label += EmptyString(strings.Count("Repository", "") +
						blankLenth - strings.Count(messages[1], ""))

					label += messages[2]
					label += EmptyString(strings.Count("Tag", "") +
						blankLenth - strings.Count(messages[2], ""))

					label += messages[3]
					label += EmptyString(strings.Count("Image ID", "") +
						blankLenth - strings.Count(messages[3], ""))

					label += messages[4]
					label += EmptyString(strings.Count("Instance ID", "") +
						blankLenth - strings.Count(messages[4], ""))

					label += messages[5]
					label += EmptyString(strings.Count("Request CPU", "") +
						blankLenth - strings.Count(messages[5], ""))

					label += messages[6]
					label += EmptyString(strings.Count("Request Mem", "") +
						blankLenth - strings.Count(messages[6], ""))

					label += "\n"

				}
				fmt.Println(label)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(instancesCmd)
	instancesCmd.SetUsageTemplate(instancesUsage)

	instancesCmd.Flags().BoolVarP(&instancesPublicFlag, "public", "p", false, "show public instances")
}
