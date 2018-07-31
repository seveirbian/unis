package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/client"
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
	dockerID := c.FormValue("dockerID")
	imageType := c.FormValue("imageType")
	// maxCPU := c.FormValue("maxCPU")
	// maxMem := c.FormValue("maxMem")
	argument := c.FormValue("argument")
	command := c.FormValue("command")

	// run docker instance
	if imageType == "docker" {
		ctx := context.Background()
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}

		if dockerID == "" {
			// import docker image

			imageFile, err := os.Open(os.Getenv("HOME") + "/.unis/unislet/images/" + imageID)
			if err != nil {
				return c.String(http.StatusNotImplemented, "image file not exists")
			}

			resp, err := cli.ImageLoad(ctx, imageFile, false)
			if err != nil {
				return c.String(http.StatusNotImplemented, "image load error!")
			}

			var result struct {
				Stream string `json:"stream"`
			}
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				return c.String(http.StatusNotImplemented, "decode failed")
			}

			dockerID = strings.Split(strings.Split(string(result.Stream), ": ")[1], "\n")[0]

			fmt.Println(dockerID)
		}

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
		output, err := child1.Output()
		if err != nil || string(output) == "" {
			return c.String(http.StatusNotImplemented, err.Error())
		}

		instanceID := strings.Split(strings.Split(string(output), "\n")[1], " ")[0]
		fmt.Println("instanceID: " + instanceID)

		return c.String(http.StatusOK, dockerID+" "+instanceID)

	} else if imageType == "unikernel" {
		return c.String(http.StatusBadGateway, "unikernel now is not supported")
	}

	return c.String(http.StatusBadRequest, "other image cannot be deployed")
}
