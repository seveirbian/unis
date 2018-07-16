package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var addUsage = `Usage:  unislet add [OPTIONS]

Options:
  -e, --environment       Set environment (Docker or Unikernel)
  -n, --node-name         Set node name
  -p, --public-node       Set node type (default private)
      --help              Print usage
`

var addEnvFlag string
var addPublicFlag bool
var nodeName string
var reservedCPU string
var reservedMem string

type Instance struct {
	ImageRepository string
	ImageTag        string
	ImageID         string
	InstanceID      string

	RequestCPU int64
	RequestMem int64
	LimitCPU   int64
	LimitMem   int64
}

var instances = []Instance{}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a node to unis",
	Long:  "add a node to unis",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ConfigContent.NodeName = nodeName
		ConfigContent.NodeEnv = addEnvFlag

		if addEnvFlag == "docker" {
			child := exec.Command("docker", "version")
			output, err := child.Output()
			if err != nil {
				logrus.Fatal(err)
			}
			ConfigContent.DockerInfo = strings.Replace(strings.Split(strings.Split(string(output), "\n")[9], ":")[1], " ", "", -1)
		} else {

		}

		tempRestCPU, _ := strconv.Atoi(reservedCPU)
		tempRestMem, _ := strconv.Atoi(reservedMem)
		ConfigContent.RestCPU = ConfigContent.TotalCPU - int64(tempRestCPU)
		ConfigContent.RestMem = ConfigContent.TotalMem - int64(tempRestMem)

		if addPublicFlag {
			// this node is public
			ConfigContent.Nodetype = "public"

			configInJSON, err := json.Marshal(ConfigContent)
			if err != nil {
				logrus.Fatal(err)
			}
			err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
			if err != nil {
				logrus.Fatal(err)
			}

			resp, err := http.PostForm(ConfigContent.Apiserver+"/nodes/add/public/"+nodeName, url.Values{"username": {ConfigContent.Username}, "password": {ConfigContent.Password}, "environment": {addEnvFlag}, "nodename": {nodeName}, "dockerinfo": {ConfigContent.DockerInfo}, "hypervisorinfo": {ConfigContent.HypervisorInfo}, "availablecpu": {strconv.Itoa(int(ConfigContent.RestCPU))}, "availablemem": {strconv.Itoa(int(ConfigContent.RestMem))}})
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

			// create a server to response apiserver
			if resp.StatusCode == http.StatusOK {
				unisletServer(":9899")
			}

		} else {
			// this node is private
			ConfigContent.Nodetype = "private"

			configInJSON, err := json.Marshal(ConfigContent)
			if err != nil {
				logrus.Fatal(err)
			}
			err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
			if err != nil {
				logrus.Fatal(err)
			}

			resp, err := http.PostForm(ConfigContent.Apiserver+"/nodes/add/"+ConfigContent.Username+"/"+nodeName, url.Values{"password": {ConfigContent.Password}, "environment": {addEnvFlag}, "nodename": {nodeName}, "dockerinfo": {ConfigContent.DockerInfo}, "hypervisorinfo": {ConfigContent.HypervisorInfo}, "availablecpu": {strconv.Itoa(int(ConfigContent.RestCPU))}, "availablemem": {strconv.Itoa(int(ConfigContent.RestMem))}})
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

			// create a server to response apiserver
			if resp.StatusCode == http.StatusOK {
				unisletServer(":9899")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.SetUsageTemplate(addUsage)

	// set node environment: docker or unikernel
	addCmd.Flags().StringVarP(&addEnvFlag, "environment", "e", "", "Set node Environment")
	addCmd.MarkFlagRequired("environment")

	// set node name
	addCmd.Flags().StringVarP(&nodeName, "node-name", "n", "", "Set node name")
	addCmd.MarkFlagRequired("node-name")

	// set node type: dafault private
	addCmd.Flags().BoolVarP(&addPublicFlag, "public", "p", false, "Set node type (public or private)")

	// set reserved cpu
	addCmd.Flags().StringVarP(&reservedCPU, "reserved-cpu", "c", "1", "Set reserved cpu for system and unislet")

	// set reserved mem
	addCmd.Flags().StringVarP(&reservedMem, "reserved-mem", "m", "3", "Set reserved memory for system and unislet")

}

func unisletServer(ipaddr string) error {
	server := echo.New()

	return server.Start(ipaddr)
}
