package repos

import (
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/provider"
	"github.com/jmoiron/sqlx"
)

type ResizeDBRepo interface {
	PostUrl(url string) error
}

type resizeDB struct {
	db *sqlx.DB
	logger.Logger
}

func NewResizeDBRepo(p provider.Provider, l logger.Logger) ResizeDBRepo {
	return &resizeDB{p.GetResizeDBConn(), l}
}

func (r resizeDB) PostUrl(url string) error {
	var query = `INSERT INTO url(id,url) values($1,$2);`
	if _, err := r.db.Exec(query, 1, url); err != nil {
		return err
	}
	return nil
}
