package v1

import (
	"net/http"

	"github.com/MorselShogiew/ResizePhoto/errs"
	"github.com/MorselShogiew/ResizePhoto/middleware"
	// _ "image/jpeg"
)

func (h *Handlers) PostUrl(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	url := r.URL.Query().Get("url")

	if url == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	err := h.u.PostUrl(reqID, url)
	h.CheckErrWriteResp(err, w, reqID)

}
