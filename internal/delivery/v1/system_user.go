package v1

import (
	"github.com/labstack/echo"
	"gitlab.com/abiewardani/scaffold/internal/delivery/middleware"
	"gitlab.com/abiewardani/scaffold/internal/usecase"
	"gitlab.com/abiewardani/scaffold/internal/utils"
)

type SystemUserHandler struct {
	systemUserUc usecase.SystemUser
}

func NewSystemUserHandler(systemUserUc usecase.SystemUser) *SystemUserHandler {
	return &SystemUserHandler{systemUserUc: systemUserUc}
}

func (h *SystemUserHandler) Mount(group *echo.Group, secretKeyBase string) {
	group.GET("/:id", h.Show, middleware.VerifyToken(secretKeyBase))
}

// Show ...
func (h *SystemUserHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param(`id`)
	res, err := h.systemUserUc.Show(ctx, id)
	if err != nil {
		return utils.JSONError(c, err)
	}
	return utils.JSONSuccess(c, res)
}
