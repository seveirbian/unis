package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var pullUsage = `Usage:  unisctl pull [OPTIONS] NAME[:TAG]

Options:
  -h, --help   help for pull
  -p, --public pull a public image
`

var pullPublicFlag bool

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image from registry",
	Long:  "Pull an image from registry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// start server to receive image
		child := exec.Command("fileReceiver", ":10001")
		go child.Output()

		// get imageID
		imageID := args[0]

		var resp *http.Response
		var err error
		if pullPublicFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/pull/public/"+imageID, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/images/pull/"+ConfigContent.Username+"/"+imageID, url.Values{"password": {ConfigContent.Password}})
		}

		// print response
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

		// kill fileReceiver
		err = child.Process.Kill()
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.SetUsageTemplate(pullUsage)
	createPullPath()

	pullCmd.Flags().BoolVarP(&pullPublicFlag, "public", "p", false, "pull a public image")
}

func createPullPath() error {
	_, err := os.Stat(os.Getenv("HOME") + "/.unis/unisctl/pulled/")
	if err != nil {
		err = os.Mkdir(os.Getenv("HOME")+"/.unis/unisctl/pulled/", os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}
	return nil
}
