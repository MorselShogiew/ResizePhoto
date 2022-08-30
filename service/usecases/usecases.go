package usecases

import (
	"github.com/MorselShogiew/ResizePhoto/repos"
)

type ResizeService struct {
	resizeAPIRepo repos.ResizeDBRepo
}

func New(r *repos.Repositories) *ResizeService {
	return &ResizeService{
		r.ResizeDBRepo,
	}
}
