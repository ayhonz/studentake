package handler

import (
	"github.com/ayhonz/studentake/view/dashboard"
	"github.com/labstack/echo/v4"
)

type Dashboard struct{}

func (d *Dashboard) ShowDashboard(c echo.Context) error {
	return render(c, dashboard.Show())
}
