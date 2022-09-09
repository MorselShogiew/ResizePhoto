package repos

import (
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/provider"
	"github.com/jmoiron/sqlx"
)

type ResizeDBRepo interface {
	//пустой пока что
}

type resizeDB struct {
	db *sqlx.DB
	logger.Logger
}

func NewResizeDBRepo(p provider.Provider, l logger.Logger) ResizeDBRepo {
	return &resizeDB{p.GetBODBConn(), l}
}
