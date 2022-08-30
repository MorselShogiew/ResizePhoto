package repos

import (
	"github.com/jmoiron/sqlx"
)

type ResizeDBRepo interface {
	//пустой пока что
}

type resizeDB struct {
	db *sqlx.DB
}

func NewResizeDBRepo() ResizeDBRepo {
	return &resizeDB{}
}
