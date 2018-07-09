package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var imagesUsage = `Usage:  unisctl images [OPTIONS]

Options:
  -a, --all    Show all images (default private images)
  -h, --help   help for images
`

var allImagesFlag bool

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images in remote registry",
	Long:  "List images in remote registry",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		//detect whether user has signed in
		if ConfigContent.Username == "" || ConfigContent.Password == "" {
			fmt.Println("Please signin first!")
			logrus.Fatal("Please signin first!")
		}
		//send different request based on allImagesFlag
		var resp *http.Response
		var err error
		if allImagesFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/show/public/images", url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/show/"+ConfigContent.Username+"/images", url.Values{"password": {ConfigContent.Password}})
		}

		//print response
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
				fmt.Println(label)
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.SetUsageTemplate(imagesUsage)
	imagesCmd.Flags().BoolVarP(&allImagesFlag, "all", "a", false, "Show all images")
}
