package internal

import (
	"github.com/ic2hrmk/azck/app"
	"github.com/ic2hrmk/azck/app/services/producer/conf"
)

type producerService struct {
	conf *conf.ProducerConfiguration
}

func NewProducerService(conf *conf.ProducerConfiguration) (app.MicroService, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &producerService{conf: conf}, nil
}

