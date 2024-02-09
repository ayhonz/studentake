package main

import (
	"github.com/ayhonz/studentake/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	homeHandler := handler.Home{}
	dashboard := handler.Dashboard{}
	e.GET("/", homeHandler.Index)
	e.GET("/dashboard", dashboard.ShowDashboard)

	e.Logger.Fatal(e.Start(":6969"))
}
