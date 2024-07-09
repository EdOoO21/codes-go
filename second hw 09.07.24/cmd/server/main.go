package main

import (
	"awesomeProject/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.PATCH("/account/patch", accountsHandler.PatchAccount)
	e.DELETE("/account/delete", accountsHandler.DeleteAccount)
	e.PATCH("/account/change", accountsHandler.ChangeAccount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
