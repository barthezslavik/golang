package main

import (
	"github.com/barthezslavik/golang/open-mind/app"
	"github.com/barthezslavik/golang/open-mind/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
