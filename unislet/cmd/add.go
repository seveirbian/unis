package cmd

import (
	"encoding/json"
	"fmt"
	"io"
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
	DockerID        string
	InstanceID      string

	RequestCPU int64
	RequestMem int64
	// LimitCPU   int64
	// LimitMem   int64
}

var instances = []Instance{}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a node to unis",
	Long:  "add a node to unis",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		child := exec.Command("lsof", "-i", ":9899")
		output, _ := child.Output()
		if len(string(output)) != 0 {
			logrus.Fatal("unislet has being running")
		}

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
			logrus.Fatal("now only support docker")
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
	addCmd.Flags().StringVarP(&reservedMem, "reserved-mem", "m", "1024", "Set reserved memory for system and unislet")

}

func unisletServer(ipaddr string) error {
	server := echo.New()

	// receive image
	server.POST("/images/receive/:imageID", handleReceiveImage)
	// run instance
	server.POST("/instances/run/:imageID", handleRunImage)

	return server.Start(ipaddr)
}

func handleReceiveImage(c echo.Context) error {
	imagename := c.Param("imageID")
	fmt.Println(imagename)

	// transimit image
	imagefile, err := c.FormFile(imagename)
	if err != nil {
		logrus.Fatal(err)
	}

	src, err := imagefile.Open()
	if err != nil {
		logrus.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(os.Getenv("HOME") + "/.unis/unislet/images/" + imagename)
	if err != nil {
		logrus.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		logrus.Fatal(err)
	}

	return c.String(http.StatusOK, "image sended")
}

func handleRunImage(c echo.Context) error {
	imageID := c.Param("imageID")
	imageType := c.FormValue("imageType")
	// maxCPU := c.FormValue("maxCPU")
	// maxMem := c.FormValue("maxMem")
	argument := c.FormValue("argument")
	command := c.FormValue("command")

	// run docker instance
	if imageType == "docker" {
		// import docker image
		arg0 := "docker"
		arg1 := "load"
		arg2 := "-i"
		arg3 := os.Getenv("HOME") + "/.unis/unislet/images/" + imageID

		child := exec.Command(arg0, arg1, arg2, arg3)
		output, err := child.Output()

		fmt.Println(string(output))
		dockerID := strings.Split(strings.Split(string(output), ": ")[1], "\n")[0]

		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(dockerID)

		// run docker image
		arguments := strings.Split(argument, " ")
		commands := strings.Split(command, " ")
		args := []string{
			"docker",
			"run",
		}
		if argument != "" {
			args = append(args, arguments...)
		}

		args = append(args, dockerID)

		if command != "" {
			args = append(args, commands...)
		}

		for _, arg := range args {
			fmt.Println(arg)
		}

		instance := exec.Command("docker")
		instance.Args = args
		output, err = instance.Output()

		fmt.Println(string(output))

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusNotImplemented, err.Error())
		}

		// get instanceID
		child1 := exec.Command("docker", "ps", "-a")
		output, err = child1.Output()
		if err != nil {
			return c.String(http.StatusNotImplemented, err.Error())
		}
		instanceID := strings.Split(strings.Split(string(output), "\n")[1], " ")[0]
		fmt.Println("instanceID: " + instanceID)

		return c.String(http.StatusOK, instanceID)

	} else if imageType == "unikernel" {
		return c.String(http.StatusBadGateway, "unikernel now is not supported")
	}

	return c.String(http.StatusBadRequest, "other image cannot be deployed")
}
