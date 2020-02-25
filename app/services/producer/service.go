package producer

import (
	"github.com/ic2hrmk/azck/app"
	"github.com/ic2hrmk/azck/app/services/producer/conf"
	"github.com/ic2hrmk/azck/app/services/producer/internal"
)

const ServiceName = "producer"

func Factory() (app.MicroService, error) {
	producerConf, err := conf.ResolveProducerConfiguration()
	if err != nil {
		return nil, err
	}

	producer, err := internal.NewProducerService(producerConf)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
