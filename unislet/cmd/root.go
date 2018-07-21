package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var defaultPath = os.Getenv("HOME") + "/.unis/unislet/"
var defaultFileName = "configure.json"

type Config struct {
	Apiserver string `json:"apiserver"`
	Username  string `json:"username"`
	Password  string `json:"password"`

	NodeName       string `json:"nodename"`
	NodeAddr       string `json:"nodeaddr"`
	Nodetype       string `json:"nodetype"`
	NodeEnv        string `json:"nodeenv"`
	DockerInfo     string `json:"dockerinfo"`
	HypervisorInfo string `json:"hypervisorinfo"`

	TotalCPU int64 `json:"availablecpu"`
	RestCPU  int64 `json:"restcpu"`
	TotalMem int64 `json:"availablemem"`
	RestMem  int64 `json:"restmem"`
}

var ConfigContent Config

type ImageInfo struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
	Type       string
	Owner      string
}

var ImagesInfo = []ImageInfo{}

var rootCmd = &cobra.Command{
	Use:   "unislet",
	Short: "An edge node to execute apiserver cmd",
	Long:  "An edge node to execute apiserver cmd",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	// create ~/.unis/unislet/images/imagesInfo.json and ~/.unis/unislet/configure.json
	createPath()

	//
}

func createPath() {
	_, err := os.Stat(defaultPath)
	if err != nil {
		err = os.Mkdir(defaultPath, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	_, err = os.Stat(defaultPath + "images/")
	if err != nil {
		err = os.Mkdir(defaultPath+"images/", os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	// create ~/.unis/unislet/images/imagesInfo.json
	_, err = os.Stat(defaultPath + "images/imagesInfo.json")
	if err != nil {
		_, err = os.Create(defaultPath + "images/imagesInfo.json")
		if err != nil {
			logrus.Fatal(err)
		}

		imagesInfoInJSON, err := json.Marshal(ImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(defaultPath+"images/imagesInfo.json", imagesInfoInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		imagesInfoInJSON, err := ioutil.ReadFile(defaultPath + "images/imagesInfo.json")
		err = json.Unmarshal(imagesInfoInJSON, &ImagesInfo)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	// create configure.json
	_, err = ioutil.ReadFile(defaultPath + defaultFileName)
	if err != nil {
		// initialize ConfigContent
		ConfigContent.Apiserver = ""
		ConfigContent.Username = ""
		ConfigContent.Password = ""

		ConfigContent.NodeName = ""
		ConfigContent.NodeAddr = ""
		ConfigContent.Nodetype = ""
		ConfigContent.NodeEnv = ""
		ConfigContent.DockerInfo = ""
		ConfigContent.HypervisorInfo = ""

		ConfigContent.TotalCPU = int64(runtime.NumCPU())
		ConfigContent.RestCPU = 0
		ConfigContent.TotalMem = int64(sysTotalMemory()) / 1048576.0
		ConfigContent.RestMem = 0

		configInJSON, err := json.Marshal(ConfigContent)
		if err != nil {
			logrus.Fatal(err)
		}

		err = ioutil.WriteFile(defaultPath+defaultFileName, configInJSON, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		configInJSON, err := ioutil.ReadFile(defaultPath + defaultFileName)
		err = json.Unmarshal(configInJSON, &ConfigContent)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func sysTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}
