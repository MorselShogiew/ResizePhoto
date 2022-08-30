package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/router"

	s "github.com/MorselShogiew/ResizePhoto/service"
)

type Application struct {
	services []s.Service
	server   *http.Server
}

func New(conf *config.Config, services ...s.Service) *Application {
	r := route(services)

	log.Println("configuration:", conf.String())
	return &Application{
		server: &http.Server{
			Addr:         ":" + conf.ServerOpts.Port,
			ReadTimeout:  conf.ServerOpts.ReadTimeout.Duration,
			IdleTimeout:  conf.ServerOpts.IdleTimeout.Duration,
			WriteTimeout: conf.ServerOpts.WriteTimeout.Duration,
			Handler:      r,
		},
		services: services,
	}
}

func route(services []s.Service) http.Handler {
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
	log.Println("http server started at port", app.server.Addr)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	app.startServices()

	select {
	case err := <-listenErr:
		if err != nil {
			log.Fatal(err)
		}
	case s := <-osSignals:
		log.Println("SIGNAL:", s.String())
		app.server.SetKeepAlivesEnabled(false)
		app.stopServices()
		timeout := time.Second * 5
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer func() {
			cancel()
		}()
		if err := app.server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Service stopped")

}

func (app *Application) startServices() {
	log.Println("Starting service")
	for i := range app.services {
		if err := app.services[i].Start(); err != nil {
			log.Fatal(fmt.Sprintf("Couldn't start service %s: %v", app.services[i].Name(), err))
		}
	}
}

func (app *Application) stopServices() {
	for i := range app.services {
		if err := app.services[i].Stop(); err != nil {
			log.Fatal(fmt.Sprintf("error while stopping service %s: %v", app.services[i].Name(), err))
		}
	}
	log.Println("Stopping service...")
}
