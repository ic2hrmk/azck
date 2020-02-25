package app

type MicroService interface {
	Run() error
	Stop() error
}
