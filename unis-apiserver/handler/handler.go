package handler

import (
	// "fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
)

type Handler struct {
	Version  string
}

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func (rqHandler Handler) Serve(serveIP string) error {
	server := echo.New()

	server.GET("/", handleConnect)

	// Run request-handler, this should never exit
	return server.Start(serveIP)
}

func handleConnect(c echo.Context) error{
	logrus.Info("Handled connect request from " + c.RealIP())
	return c.String(http.StatusOK, "OK")
}