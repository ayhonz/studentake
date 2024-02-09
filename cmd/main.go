package main

import (
	"github.com/ayhonz/studentake/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	homeHandler := handler.Home{}
	e.GET("/", homeHandler.Index)

	e.Logger.Fatal(e.Start(":6969"))
}
