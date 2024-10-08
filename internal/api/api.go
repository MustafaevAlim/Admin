package api

import (
	_ "myapp/docs"
	"myapp/internal/api/controllers"
	"myapp/internal/app"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func RouteController(a *app.App) {
	handler := controllers.NewHandler(a.DB)

	a.ServerEcho.POST("/auth", handler.Authorization)
	a.ServerEcho.POST("/users/add", handler.AddUser)
	a.ServerEcho.GET("/socials/get", handler.GetAllSocial)
	a.ServerEcho.GET("/users/all", handler.GetAllUsers)
	a.ServerEcho.GET("/users", handler.GetUser)
	a.ServerEcho.GET("/users/views", handler.GetUserViewsSocial)
	a.ServerEcho.PUT("/users/balance", handler.UpdateBalance)
	a.ServerEcho.GET("/swagger/*", echoSwagger.WrapHandler)
}
