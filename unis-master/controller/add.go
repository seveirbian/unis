package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func handleAddUser(c echo.Context) error {
	username := c.Param("username")

	privateNodes[username] = []NodeInfo{}

	return c.String(http.StatusOK, "add user to privateNodesInfo")
}

func handleAddPublicNode(c echo.Context) error {
	nodename := c.Param("nodename")
	nodetype := c.FormValue("nodetype")
	nodeenv := c.FormValue("nodeenv")
	nodeaddr := c.FormValue("nodeaddr")
	dockerinfo := c.FormValue("dockerinfo")
	hypervisorinfo := c.FormValue("hypervisorinfo")
	totalcpu, _ := strconv.Atoi(c.FormValue("totalcpu"))
	totalmem, _ := strconv.Atoi(c.FormValue("totalmem"))

	publicNodes = append(publicNodes, NodeInfo{
		NodeName:       nodename,
		NodeType:       nodetype,
		NodeEnv:        nodeenv,
		NodeAddr:       nodeaddr,
		DockerInfo:     dockerinfo,
		HypervisorInfo: hypervisorinfo,
		TotalCPU:       int64(totalcpu),
		TotalMem:       int64(totalmem),
	})

	fmt.Println(publicNodes)

	return c.String(http.StatusOK, "public node added")
}

func handleAddPrivateNode(c echo.Context) error {
	username := c.Param("username")
	nodename := c.Param("nodename")
	nodetype := c.FormValue("nodetype")
	nodeenv := c.FormValue("nodeenv")
	nodeaddr := c.FormValue("nodeaddr")
	dockerinfo := c.FormValue("dockerinfo")
	hypervisorinfo := c.FormValue("hypervisorinfo")
	totalcpu, _ := strconv.Atoi(c.FormValue("totalcpu"))
	totalmem, _ := strconv.Atoi(c.FormValue("totalmem"))

	privateNodes[username] = append(privateNodes[username], NodeInfo{
		NodeName:       nodename,
		NodeType:       nodetype,
		NodeEnv:        nodeenv,
		NodeAddr:       nodeaddr,
		DockerInfo:     dockerinfo,
		HypervisorInfo: hypervisorinfo,
		TotalCPU:       int64(totalcpu),
		TotalMem:       int64(totalmem),
	})

	fmt.Println(privateNodes)
	return c.String(http.StatusOK, "private node added")
}
