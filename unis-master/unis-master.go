package main

import (
	"github.com/seveirbian/unis/unis-master/apiserver"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.FatalLevel)

	var apiServer = apiserver.Server{}

	apiServer.Serve(":9898")
}
