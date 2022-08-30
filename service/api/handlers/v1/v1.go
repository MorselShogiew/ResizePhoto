package v1

import (
	"log"
	"net/http"

	"github.com/MorselShogiew/ResizePhoto/service/usecases"
)

type Handlers struct {
	u *usecases.ResizeService
}

func New(u *usecases.ResizeService) *Handlers {
	return &Handlers{u}
}

func (h *Handlers) CheckErrWriteResp(e error, w http.ResponseWriter, requestID string) {
	if e == nil {
		w.WriteHeader(200)
		return
	}

	if err, ok := e.(interface{ StatusCode() int }); ok {
		log.Fatal(requestID, e)
		w.WriteHeader(err.StatusCode())
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(e.Error()))
		return
	}

	log.Fatal(requestID, e)
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(e.Error()))
}
