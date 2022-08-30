package v1

import (
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/MorselShogiew/ResizePhoto/errs"
	"github.com/MorselShogiew/ResizePhoto/middleware"
	// _ "image/jpeg"
)

func (h *Handlers) GetResizePhoto(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	// employeeID := middleware.GetEmployeeID(r)

	// if employeeID <= 0 {
	// 	err := errs.New(errors.New(errs.ErrWrongEmployeeID), errs.ErrBadRequest, false, 400)
	// 	h.CheckErrWriteResp(err, w, reqID)
	// 	return
	// }
	heightStr := r.URL.Query().Get("height")
	widthStr := r.URL.Query().Get("width")
	url := r.URL.Query().Get("url")

	if heightStr == "" || widthStr == "" || url == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	height, err := strconv.ParseUint(heightStr, 10, 32)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	width, err := strconv.ParseUint(widthStr, 10, 32)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	res, err := h.u.ResizePhoto(reqID, height, width, url)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

	if err = jpeg.Encode(w, res, nil); err != nil {
		err := errs.New(err, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

}
