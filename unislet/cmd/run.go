package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

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

		// docker run
		instance := exec.Command("docker")
		instance.Args = args
		err = instance.Start()

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusNotImplemented, err.Error())
		}

		time.Sleep(time.Second * 3)

		// get instanceID
		child1 := exec.Command("docker", "ps")
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
