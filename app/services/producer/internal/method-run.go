package internal

import (
	"log"
	"time"
)

func (rcv *producerService) Run() error {
	ticker := time.NewTicker(getPeriodFromFrequency(rcv.conf))

	for range ticker.C {
		log.Printf("new value: -> %d", rand.Int())
	}

	return nil
}

func getPeriodFromFrequency(freq float64) time.Duration {
	if freq == 0 {
		return 0
	}

	//
	//		1000 Milliseconds / frequency
	//
	return time.Duration(1000.0/freq) * time.Millisecond
}
