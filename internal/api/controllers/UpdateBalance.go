package controllers

import (
	"fmt"
	"myapp/internal/repository/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UpdateBalance godoc
// @Summary Обновление баланса пользователя
// @Description Обновляет баланс пользователя на заданное значение.
// @Tags users
// @Accept  json
// @Produce  json
// @Param tgID query string true "Telegram ID пользователя"
// @Param value query string true "Значение для обновления баланса"
// @Success 200 {object} map[string]string "succes"
// @Success 500 {object} map[string]string "ошибка обновления баланса"
// @Router /users/balance [put]
func (h *Handler) UpdateBalance(c echo.Context) error {
	id := c.QueryParam("tgID")
	value := c.QueryParam("value")
	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		fmt.Println("Ошибка при преобразовании:", err)
	}
	result := user.UpdateBalance(id, float32(floatValue), h.repository)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "ошибка обновления баланса"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "succes"})

}
