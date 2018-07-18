package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func recieveFile(addr string) error {
	rServer := echo.New()

	rServer.POST("/images/pull/:imagename", func(c echo.Context) error {
		// transimit image
		imagename := c.Param("imagename")
		imagefile, err := c.FormFile(imagename)
		if err != nil {
			logrus.Fatal(err)
		}

		src, err := imagefile.Open()
		if err != nil {
			logrus.Fatal(err)
		}
		defer src.Close()

		dst, err := os.Create(os.Getenv("HOME") + "/.unis/unisctl/pulled/" + imagename)
		if err != nil {
			logrus.Fatal(err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			logrus.Fatal(err)
		}

		return c.String(http.StatusOK, "image pulled")
	})

	return rServer.Start(addr)
}

func main() {
	addr := os.Args[1]

	recieveFile(addr)
}
