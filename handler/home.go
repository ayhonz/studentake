package handler

import (
	"github.com/ayhonz/studentake/view/home"
	"github.com/labstack/echo/v4"
)

type Home struct{}

func (h *Home) Index(c echo.Context) error {
	return render(c, home.Show())
}
