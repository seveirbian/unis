package cmd

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func handleStopInstance(c echo.Context) error {
	instanceID := c.Param("instanceID")

	arg0 := "docker"
	arg1 := "kill"
	arg2 := "rm"

	stopInstance := exec.Command(arg0, arg1, instanceID)
	output, err := stopInstance.Output()
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusNotImplemented, "fail to kill instance")
	}

	rmInstance := exec.Command(arg0, arg2, instanceID)
	output, err = rmInstance.Output()
	if err != nil {
		return c.String(http.StatusNotImplemented, "fail to rm instance")
	}

	return c.String(http.StatusOK, string(output))
}
