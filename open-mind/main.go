package main

import (
	"os"

	"github.com/barthezslavik/golang/open-mind/config"
	"github.com/barthezslavik/golang/open-mind/controllers"
	"github.com/labstack/echo"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	e := echo.New()

	config.Setup(e)
	controllers.Setup(e.Router())

	e.File("/", "index.html")
	e.Static("/static", "static")

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}

// vi:syntax=go
