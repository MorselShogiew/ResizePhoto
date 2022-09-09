package database

import (
	"fmt"
	"net/url"

	"github.com/MorselShogiew/ResizePhoto/config"
	"github.com/jmoiron/sqlx"
)

func Connect(conf *config.DB) (*sqlx.DB, error) {
	query := url.Values{}

	switch conf.Scheme {
	case "postgres":
		query.Add("dbname", conf.Database)
		if !conf.SSLMode {
			query.Add("sslmode", "disable")
		}
	case "sqlserver":
		query.Add("database", conf.Database)
		query.Add("failoverpartner", conf.FailoverHost)
	default:
		return nil, fmt.Errorf("unknown db scheme")
	}

	host := conf.Server
	if conf.Port != "" {
		host += ":" + conf.Port
	}

	u := &url.URL{
		Scheme:   conf.Scheme,
		User:     url.UserPassword(conf.Username, conf.Password),
		Host:     host,
		RawQuery: query.Encode(),
	}

	db, err := sqlx.Open(conf.Scheme, u.String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxOpenConns)
	if conf.MaxOpenConns <= 0 {
		db.SetMaxOpenConns(10)
	}
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime.Duration)

	return db, nil
}
