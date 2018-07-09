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

var tagUsage = `Usage:  unisctl tag [OPTIONS] SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

Options:
  -h, --help            help for tag
  -p, --public          tag a public image
`

var tagPublicFlag bool

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE",
	Long:  "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var oldImageName string
		var oldTag string
		var newImageName string
		var newTag string

		if strings.Contains(args[0], ":") {
			temp := strings.Split(args[0], ":")
			oldImageName = temp[0]
			oldTag = temp[1]
			if oldImageName == "" {
				logrus.Fatal("Old image name must set")
			}
			if oldTag == "" {
				oldTag = "latest"
			}
		}

		if strings.Contains(args[1], ":") {
			temp := strings.Split(args[1], ":")
			newImageName = temp[0]
			newTag = temp[1]
			if newImageName == "" {
				logrus.Fatal("Old image name must set")
			}
			if newTag == "" {
				newTag = "latest"
			}
		}

		if tagPublicFlag {
			// tag public image
			resp, err := http.PostForm(ConfigContent.Apiserver+"/images/tag/public/"+oldImageName+"/"+oldTag+"/"+newImageName+"/"+newTag, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
			if err != nil {
				logrus.Fatal(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatal(err)
			} else {
				fmt.Println(string(body))
			}

		} else {
			// tag private image
			resp, err := http.PostForm(ConfigContent.Apiserver+"/images/tag/"+ConfigContent.Username+"/"+oldImageName+"/"+oldTag+"/"+newImageName+"/"+newTag, url.Values{"password": {ConfigContent.Password}})
			if err != nil {
				logrus.Fatal(err)
			}

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
	rootCmd.AddCommand(tagCmd)
	tagCmd.SetUsageTemplate(tagUsage)
	tagCmd.Flags().BoolVarP(&tagPublicFlag, "public", "p", false, "Tag a public image")
}
