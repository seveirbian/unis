package main

import (
	// "fmt"

	"github.com/sirupsen/logrus"
	"github.com/seveirbian/unis/unis-apiserver/handler"
)

func main() {
	logrus.SetLevel(logrus.FatalLevel)

	var rqHandler = handler.Handler {
		Version: "", 
	}

	rqHandler.Serve(":9898")
}