package apiserver

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleSignin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if validateUser(username, password) {
		return c.String(http.StatusOK, "Signin succeeded")
	} else {
		return c.String(http.StatusUnauthorized, "incorrect username or password")
	}
}
