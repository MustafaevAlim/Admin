package controllers

import (
	"myapp/internal/repository/socials"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllUserNote godoc
// @Summary Получить общее количество просмотров по соцсетям
// @Description Возвращает все просмотры соцсетей
// @Tags socials
// @Produce  json
// @Success 200 {object} map[string]int "Список общих просмотров"
// @Failure 400 {object} map[string]string "Неверный email"
// @Router /socials/get [get]
func (h *Handler) GetAllSocial(c echo.Context) error {
	social := []string{"youtube", "tiktok", "instagram", "facebook"}
	views := make(map[string]int)
	var sum int
	for _, s := range social {
		count := socials.GetAllViews(s, h.repository)
		views[s] = count
		sum += count
	}
	views["total"] = sum

	return c.JSON(http.StatusOK, views)
}
