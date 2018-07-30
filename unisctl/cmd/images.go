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

var imagesUsage = `Usage:  unisctl images [OPTIONS]

Options:
  -p, --public Show public images (default private images)
  -h, --help   help for images
`

var allImagesFlag bool

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images in remote registry",
	Long:  "List images in remote registry",
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
		if allImagesFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/show/public/images", url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/show/"+ConfigContent.Username+"/images", url.Values{"password": {ConfigContent.Password}})
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
				label += "REPOSITORY          "
				label += "TAG          "
				label += "IMAGE ID          "
				label += "CREATED          "
				label += "SIZE          "
				label += "TYPE          "
				label += "OWNER          "
				label += "\n"

				imageLine := strings.Split(string(body), "\n")
				blankLenth := 10

				for _, image := range imageLine {
					if image == "" {
						break
					}
					messages := strings.Split(image, " ")

					label += messages[0]
					label += EmptyString(strings.Count("Repository", "") +
						blankLenth - strings.Count(messages[0], ""))

					label += messages[1]
					label += EmptyString(strings.Count("Tag", "") +
						blankLenth - strings.Count(messages[1], ""))

					label += messages[2]
					label += EmptyString(strings.Count("Image ID", "") +
						blankLenth - strings.Count(messages[2], ""))

					label += messages[3]
					label += EmptyString(strings.Count("Created", "") +
						blankLenth - strings.Count(messages[3], ""))

					label += messages[4]
					label += EmptyString(strings.Count("Size", "") +
						blankLenth - strings.Count(messages[4], ""))

					label += messages[5]
					label += EmptyString(strings.Count("Type", "") +
						blankLenth - strings.Count(messages[5], ""))

					label += messages[6]
					label += EmptyString(strings.Count("Owner", "") +
						blankLenth - strings.Count(messages[6], ""))

					label += "\n"

				}
				fmt.Println(label)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.SetUsageTemplate(imagesUsage)
	imagesCmd.Flags().BoolVarP(&allImagesFlag, "public", "p", false, "Show public images")
}
