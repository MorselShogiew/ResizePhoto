package usecases

// i

func (r *ResizeService) PostUrl(reqID string, url string) error {
	return r.resizeAPIRepo.PostUrl(url)
}
