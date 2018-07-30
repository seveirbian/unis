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

var nodesUsage = `Usage:  unisctl nodes [OPTIONS]

Options:
  -p, --public Show public nodes (default private nodes)
  -h, --help   help for nodes
`

var allNodesFlag bool

var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Display the status of  edge nodes",
	Long:  "Display the status of  edge nodes",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// send different request based on allImagesFlag
		var resp *http.Response
		var err error
		if allNodesFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/nodes/show/public/nodes", url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/nodes/show/"+ConfigContent.Username+"/nodes", url.Values{"password": {ConfigContent.Password}})
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
				label += "NAME          "
				label += "ADDR          "
				label += "TYPE          "
				label += "ENV          "
				label += "DOCKER          "
				label += "HYPERVISOR          "
				label += "AVAIL CPU          "
				label += "AVAIL MEM          "
				label += "NODE STATUS          "
				label += "\n"

				imageLine := strings.Split(string(body), "\n")
				blankLenth := 10

				for _, image := range imageLine {
					if image == "" {
						break
					}
					messages := strings.Split(image, " ")

					label += messages[0]
					label += EmptyString(strings.Count("Name", "") +
						blankLenth - strings.Count(messages[0], ""))

					label += messages[1]
					label += EmptyString(strings.Count("Addr", "") +
						blankLenth - strings.Count(messages[1], ""))

					label += messages[2]
					label += EmptyString(strings.Count("Type", "") +
						blankLenth - strings.Count(messages[2], ""))

					label += messages[3]
					label += EmptyString(strings.Count("Env", "") +
						blankLenth - strings.Count(messages[3], ""))

					label += messages[4]
					label += EmptyString(strings.Count("Docker", "") +
						blankLenth - strings.Count(messages[4], ""))

					label += messages[5]
					label += EmptyString(strings.Count("Hypervisor", "") +
						blankLenth - strings.Count(messages[5], ""))

					label += messages[6]
					label += EmptyString(strings.Count("Avail CPU", "") +
						blankLenth - strings.Count(messages[6], ""))

					label += messages[7]
					label += EmptyString(strings.Count("Avail Mem", "") +
						blankLenth - strings.Count(messages[7], ""))

					label += messages[8]
					label += EmptyString(strings.Count("NODE STATUS", "") +
						blankLenth - strings.Count(messages[8], ""))

					label += "\n"
				}
				fmt.Println(label)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
	nodesCmd.SetUsageTemplate(nodesUsage)
	nodesCmd.Flags().BoolVarP(&allNodesFlag, "public", "p", false, "Show public nodes")
}
