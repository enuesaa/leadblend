package controller

import "github.com/labstack/echo/v4"

func ListSpaces(c echo.Context) error {
	return WithData(c, "")
}
