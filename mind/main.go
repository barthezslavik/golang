package main

import (
	"github.com/barthezslavik/golang/mind/app"
	"github.com/barthezslavik/golang/mind/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
