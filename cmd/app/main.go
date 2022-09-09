package main

import (
	"github.com/MorselShogiew/ResizePhoto/application"
	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/logger/opt"
	"github.com/MorselShogiew/ResizePhoto/provider"
	"github.com/MorselShogiew/ResizePhoto/repos"
	"github.com/MorselShogiew/ResizePhoto/service/api"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func main() {
	conf := config.LoadConfig()
	conf.InstanceID = uuid.New()
	opts := makeLoggerOpts(conf)
	l := logger.New(opts)
	p := provider.New(conf, l)

	repositories := repos.New(p, l)

	resizePhotoService := api.New(l, repositories)

	app := application.New(conf, l, resizePhotoService)
	app.Start()
}

func makeLoggerOpts(c *config.Config) *opt.LoggerOpts {
	return &opt.LoggerOpts{
		Opts: &opt.GeneralOpts{
			InstanceID: c.InstanceID,
			Env:        c.Environment,
			AppName:    c.ApplicationName,
			Level:      c.Logger.Level,
		},
		StdLoggerOpts: &opt.StdLoggerOpts{
			LogFile:  c.Logger.LoggerStd.LogFile,
			Stdout:   c.Logger.LoggerStd.Stdout,
			Disabled: c.Logger.LoggerStd.Disabled,
		},
	}
}
