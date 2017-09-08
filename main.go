package main

import (
	"net/http"
	"os"

	echo "gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<!DOCTYPE html><html><head><title>충남삼성고등학교 총동문회</title></head><body><iframe src=\"https://docs.google.com/forms/d/e/1FAIpQLSfzZ1lCGGEe21yvVbFbMIIZWE6izFySdQ9sOX3Coci5Zu03Ag/viewform?embedded=true\" style=\"position: absolute; width: 100%; height: 100%; border: none;\" frameborder=\"0\" marginheight=\"0\" marginwidth=\"0\">로드 중...</iframe></body></html>")
	})

	//e.Start(":80")
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
