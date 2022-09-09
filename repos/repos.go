package repos

import (
	"github.com/MorselShogiew/ResizePhoto/logger"
	"github.com/MorselShogiew/ResizePhoto/provider"
)

type Repositories struct {
	ResizeDBRepo ResizeDBRepo
}

func New(p provider.Provider, l logger.Logger) *Repositories {
	ResizeDBRepo := NewResizeDBRepo(p, l)
	return &Repositories{
		ResizeDBRepo,
	}
}
