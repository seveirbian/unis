package apiserver

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handleConnect(c echo.Context) error {
	logrus.Info("Handling connect request from " + c.RealIP())
	return c.String(http.StatusOK, "OK")
}
