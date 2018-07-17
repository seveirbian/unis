package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleConnect(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
