package main

import (
	"github.com/seveirbian/unis/unis-master/apiserver"
	"github.com/seveirbian/unis/unis-master/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.FatalLevel)

	var apiServer = apiserver.Server{}
	var controller = controller.Controller{}

	// start controlller
	go controller.Start(":10000")

	// start apiserver
	apiServer.Serve(":9898")
}
