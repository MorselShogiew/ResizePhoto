package provider

import (
	"net/http"
	"time"

	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/jmoiron/sqlx"
)

type Provider interface {
	GetBODBConn() *sqlx.DB

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
	bodb *sqlx.DB
}

// third party apis
type apis struct {
}

func New(conf *config.Config, l logger.Logger) Provider {

	// bodb, err := database.Connect(conf.BODB)
	// if err != nil {
	// 	l.Fatal(err)
	// }
	l.Info("connected to bodb")

	// создаем http клиента
	c := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &provider{
		&resources{
			//bodb,
		},
		&apis{},
		c,
		l,
	}
}

func (p *provider) GetBODBConn() *sqlx.DB {
	return p.resources.bodb
}

func (p *provider) GetHTTPClient() *http.Client {
	return p.c
}

func (p *provider) Close() {
	if err := p.resources.bodb.Close(); err != nil {
		p.l.Error("error while closing bodb:", err)
	} else {
		p.l.Info("bodb was closed")
	}

}
