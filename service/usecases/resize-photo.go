package usecases

import (
	"fmt"
	"image" // i
	"io"
	"net/http"

	"github.com/nfnt/resize"
)

func (r *ResizeService) ResizePhoto(reqID string, height uint64, width uint64, url string) (image.Image, error) {
	response, err := r.DownloadFile(url)
	if err != nil {
		return nil, err
	}

	// Decoding gives you an Image.
	// If you have an io.Reader already, you can give that to Decode
	// without reading it into a []byte.
	image, _, err := image.Decode(response)
	// check err

	newImage := resize.Resize(uint(width), uint(height), image, resize.Lanczos3)
	return newImage, nil

}

func (r *ResizeService) DownloadFile(URL string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return resp.Body, nil
}
