package controllers

import (
	"fmt"
	"myapp/internal/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser godoc
// @Summary      Получить информацию о пользователе с помощью Telegram ID
// @Description  Извлекает данные пользователя на основе предоставленного идентификатора Telegram.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tgID    query  string  true  "Telegram user ID"
// @Success      200   {object} model.UserInfo  "User details"
// @Failure      404   {object} map[string]string "Error message: User not found"
// @Failure      500   {object} map[string]string "Error message: Internal server error"
// @Router       /users [get]
func (h *Handler) GetUser(c echo.Context) error {
	id := c.QueryParam("tgID")
	fmt.Println("Received tgID:", id) // Логируем полученный tgID
	u := user.Get(id, h.repository)

	return c.JSON(http.StatusOK, u)

}
