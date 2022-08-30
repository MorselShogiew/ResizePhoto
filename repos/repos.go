package repos

type Repositories struct {
	ResizeDBRepo ResizeDBRepo
}

func New() *Repositories {
	ResizeDBRepo := NewResizeDBRepo()
	return &Repositories{
		ResizeDBRepo,
	}
}
