package services

type Service interface {
	Start() error
	Stop() error
	Name() string
}
