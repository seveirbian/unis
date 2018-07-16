package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var nodesUsage = `Usage:  unisctl nodes [OPTIONS]

Options:
  -a, --all    Show all nodes (default private nodes)
  -h, --help   help for nodes
`

var allNodesFlag bool

var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Display the status of all edge nodes",
	Long:  "Display the status of all edge nodes",
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
				fmt.Println(label)
				fmt.Println(string(body))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
	nodesCmd.SetUsageTemplate(nodesUsage)
	nodesCmd.Flags().BoolVarP(&allNodesFlag, "all", "a", false, "Show all nodes")
}
