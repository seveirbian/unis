package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var runUsage = `Usage:  unisctl run [OPTIONS] IMAGE

Options:
	  --argument        argument to start instance
      --command         command to start instance
  -c, --request-cpu     cpu request
      --max-cpu         cpu limit 
  -m, --request-memory  memory request (MB)
	  --max-memory      memory limit (MB)
  -h, --help            help for run
  -p, --public          run as a public instance
`

var runPublicFlag bool
var runInstanceArgument string
var runInstanceCommand string
var cpuRequest string
var memRequest string

// var cpuMax string
// var memMax string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a instance on a edge node",
	Long:  "Run a instance on a edge node",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageID := args[0]
		var resp *http.Response
		var err error
		if runPublicFlag {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/instances/run/public/"+imageID, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}, "argument": {runInstanceArgument}, "command": {runInstanceCommand}, "requestcpu": {cpuRequest}, "requestmem": {memRequest}, "imageid": {imageID}})
		} else {
			resp, err = http.PostForm(ConfigContent.Apiserver+"/instances/run/"+ConfigContent.Username+"/"+imageID, url.Values{"password": {ConfigContent.Password}, "argument": {runInstanceArgument}, "command": {runInstanceCommand}, "requestcpu": {cpuRequest}, "requestmem": {memRequest}, "imageid": {imageID}})
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
	rootCmd.AddCommand(runCmd)
	runCmd.SetUsageTemplate(runUsage)
	runCmd.Flags().BoolVarP(&runPublicFlag, "public", "p", false, "Run as a public instance")

	runCmd.Flags().StringVarP(&runInstanceArgument, "argument", "", "", "argument to start instance")
	runCmd.Flags().StringVarP(&runInstanceCommand, "command", "", "", "command to start instance")

	runCmd.Flags().StringVarP(&cpuRequest, "request-cpu", "c", "0.5", "request cpu to start instance")
	runCmd.MarkFlagRequired("request-cpu")

	runCmd.Flags().StringVarP(&memRequest, "request-memory", "m", "200", "request mem to start instance")
	runCmd.MarkFlagRequired("request-memory")

	// runCmd.Flags().StringVarP(&cpuMax, "max-cpu", "", "", "instance cpu limit")
	// runCmd.MarkFlagRequired("max-cpu")

	// runCmd.Flags().StringVarP(&memMax, "max-memory", "", "", "instance memory limit")
	// runCmd.MarkFlagRequired("max-memory")
}
