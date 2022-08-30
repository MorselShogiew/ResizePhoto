package main

import (
	"github.com/MorselShogiew/ResizePhoto/application"
	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/repos"
	"github.com/MorselShogiew/ResizePhoto/service/api"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func main() {
	conf := config.LoadConfig()
	conf.InstanceID = uuid.New()

	repositories := repos.New()

	resizePhotoService := api.New(repositories)

	app := application.New(conf, resizePhotoService)
	app.Start()
}
