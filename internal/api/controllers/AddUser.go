package controllers

import (
	"myapp/internal/model"
	"myapp/internal/repository/socials"
	"myapp/internal/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Registration godoc
// @Summary Добавление пользователя
// @Description Добавляет нового пользователя, если его ID в Telegram (TgId) еще не существует в системе, если существует добавляет новый канал соцсети
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body  model.UserAdd  true  "Данные для добавления пользователя"
// @Success 201 {object} model.UserAdd "Пользователь успешно зарегистрирован"
// @Failure 400 {object} map[string]string "Неправильный ввод данных"
// @Router /users/add [post]
func (h *Handler) AddUser(c echo.Context) error {
	u := new(model.UserAdd)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	userFromDB, err := user.GetFromTgId(u.TgId, h.repository)
	if err == nil {
		socials.InRepo(*u, userFromDB.UserID, h.repository)
		user.IncChannels(userFromDB.TgId, h.repository)
		return c.JSON(http.StatusBadRequest, u)
	} else {
		user.InRepo(*u, h.repository)
		userId, err := user.GetIdFromTgId(u.TgId, h.repository)
		if err != nil {
			panic(err.Error)
		}
		socials.InRepo(*u, userId, h.repository)
		user.IncChannels(u.TgId, h.repository)
	}

	return c.JSON(http.StatusCreated, u)
}
