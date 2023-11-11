package main

import (
	"my-go-starter/db"
	"log"
	"my-go-starter/handler"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.HideBanner = true
	e.HTTPErrorHandler = handler.HTTPErrorHandler

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	e.GET("/users", handler.GetUsers)
	e.POST("/users/mask", handler.MaskUserHandler)
	
	e.Start(":8080")

}