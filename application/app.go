package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/router"

	s "github.com/MorselShogiew/ResizePhoto/service"
)

type Application struct {
	services []s.Service
	server   *http.Server
	l        logger.Logger
}

func New(conf *config.Config, l logger.Logger, services ...s.Service) *Application {
	r := route(services, l)

	l.Info("configuration:", conf.String())
	return &Application{
		server: &http.Server{
			Addr:         ":" + conf.ServerOpts.Port,
			ReadTimeout:  conf.ServerOpts.ReadTimeout.Duration,
			IdleTimeout:  conf.ServerOpts.IdleTimeout.Duration,
			WriteTimeout: conf.ServerOpts.WriteTimeout.Duration,
			Handler:      r,
		},
		services: services,
		l:        l,
	}
}

func route(services []s.Service, l logger.Logger) http.Handler {
	r := router.Router(getApis(services)...)

	return r
}

func getApis(services []s.Service) (apis []router.API) {
	for i := range services {
		if v, ok := services[i].(router.API); ok {
			apis = append(apis, v)
		}
	}

	return apis
}

func (app *Application) Start() {
	listenErr := make(chan error, 1)
	go func() {
		listenErr <- app.server.ListenAndServe()
	}()
	app.l.Info("http server started at port", app.server.Addr)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	app.startServices()

	select {
	case err := <-listenErr:
		if err != nil {
			app.l.Fatal(err)
		}
	case s := <-osSignals:
		app.l.Info("SIGNAL:", s.String())
		app.server.SetKeepAlivesEnabled(false)
		app.stopServices()
		timeout := time.Second * 5
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer func() {
			cancel()
		}()
		if err := app.server.Shutdown(ctx); err != nil {
			app.l.Fatal(err)
		}
	}
	app.l.Info("Service stopped")

}

func (app *Application) startServices() {
	app.l.Info("Starting service")
	for i := range app.services {
		if err := app.services[i].Start(); err != nil {
			app.l.Fatal(fmt.Sprintf("Couldn't start service %s: %v", app.services[i].Name(), err))
		}
	}
}

func (app *Application) stopServices() {
	for i := range app.services {
		if err := app.services[i].Stop(); err != nil {
			app.l.Fatal(fmt.Sprintf("error while stopping service %s: %v", app.services[i].Name(), err))
		}
	}
	app.l.Info("Stopping service...")
}
