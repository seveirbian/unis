package apiserver

import (
	"net/http"

	"github.com/labstack/echo"
)

func handlePrivateImages(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	if validateUser(username, password) {
		//get user's imagesInfo
		imagesInfo := getPrivateImagesInfo(username)

		//generate response body
		var bodyContent = ""
		// var blankLenth = 10
		for _, image := range imagesInfo {
			bodyContent += image.Repository
			bodyContent += " "

			bodyContent += image.Tag
			bodyContent += " "

			bodyContent += Substring(image.ImageID, 0, 10)
			bodyContent += " "

			bodyContent += image.Created
			bodyContent += " "

			bodyContent += image.Size
			bodyContent += " "

			bodyContent += image.Type
			bodyContent += " "

			bodyContent += image.Owner
			bodyContent += " "

			bodyContent += "\n"
		}
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}

func handlePublicImages(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if validateUser(username, password) {
		//get public images info
		publicImagesInfo := getPublicImagesInfo()

		// //get private images info
		// privateImagesInfo := getPrivateImagesInfo(username)

		//generate response body
		var bodyContent = ""
		// var blankLenth = 10
		for _, image := range publicImagesInfo {
			bodyContent += image.Repository
			bodyContent += " "

			bodyContent += image.Tag
			bodyContent += " "

			bodyContent += Substring(image.ImageID, 0, 10)
			bodyContent += " "

			bodyContent += image.Created
			bodyContent += " "

			bodyContent += image.Size
			bodyContent += " "

			bodyContent += image.Type
			bodyContent += " "

			bodyContent += image.Owner
			bodyContent += " "

			bodyContent += "\n"
		}
		// for _, image := range privateImagesInfo {
		// 	bodyContent += image.Repository

		// 	bodyContent += image.Tag
		// 	bodyContent += " "

		// 	bodyContent += Substring(image.ImageID, 0, 10)
		// 	bodyContent += " "

		// 	bodyContent += image.Created
		// 	bodyContent += " "

		// 	bodyContent += image.Size
		// 	bodyContent += " "

		// 	bodyContent += image.Type
		// 	bodyContent += " "

		// 	bodyContent += image.Owner
		// 	bodyContent += " "

		// 	bodyContent += "\n"
		// }
		return c.String(http.StatusOK, bodyContent)
	}
	return c.String(http.StatusUnauthorized, "incorrect username or password")
}
