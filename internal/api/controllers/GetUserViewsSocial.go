package controllers

import (
	"myapp/internal/repository/socials"
	"myapp/internal/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUserViewsSocial godoc
// @Summary      Получить просмотры о конкретной социальной платформе
// @Description  Извлекает количество просмотров для социальной платформы пользователя на основе предоставленного идентификатора Telegram и названия социальной платформы.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        social  query  string  true  "Social platform name"
// @Param        tgID    query  string  true  "Telegram user ID"
// @Success      200     {object} map[string]int  "Count of views and link"
// @Failure      400     {object} map[string]string "Error message"
// @Router       /users/views [get]
func (h *Handler) GetUserViewsSocial(c echo.Context) error {
	s := c.QueryParam("social")
	id := c.QueryParam("tgID")
	uid, err := user.GetIdFromTgId(id, h.repository)
	if err != nil {
		panic(err)
	}
	count_views := socials.GetViews(uid, s, h.repository)
	link := socials.GetLink(uid, s, h.repository)

	return c.JSON(http.StatusOK, map[string]string{"count_views": count_views, "link": link})

}
