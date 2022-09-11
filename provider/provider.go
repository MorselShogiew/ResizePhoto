package provider

import (
	"net/http"
	"time"

	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/provider/database"
	"github.com/jmoiron/sqlx"
)

type Provider interface {
	GetResizeDBConn() *sqlx.DB

	Close()
}

type provider struct {
	resources *resources
	apis      *apis
	c         *http.Client
	l         logger.Logger
}

// resources that should be closed manually at the end
type resources struct {
	resizedb *sqlx.DB
}

// third party apis
type apis struct {
}

func New(conf *config.Config, l logger.Logger) Provider {

	bodb, err := database.Connect(conf.ResizeDB)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("connected to resizedb")

	// создаем http клиента
	c := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &provider{
		&resources{
			bodb,
		},
		&apis{},
		c,
		l,
	}
}

func (p *provider) GetResizeDBConn() *sqlx.DB {
	return p.resources.resizedb
}

func (p *provider) GetHTTPClient() *http.Client {
	return p.c
}

func (p *provider) Close() {
	if err := p.resources.resizedb.Close(); err != nil {
		p.l.Error("error while closing resizedb:", err)
	} else {
		p.l.Info("resizedb was closed")
	}

}
