package application

import (
	"fmt"
	"sync"

	"github.com/FW-Systeme/MSL5-Event-Lib/domain"
)

type EventRegistry struct {
	wg  *sync.WaitGroup
	mem map[string]map[chan interface{}]bool
}

func New() EventRegistry {
	return EventRegistry{
		mem: make(map[string]map[chan interface{}]bool),
		wg:  &sync.WaitGroup{}}
}

func (bus EventRegistry) Publish(typeId string, event interface{}) error {
	chanMap, ok := bus.mem[typeId]
	if !ok {
		return fmt.Errorf("no subscriber for typeId: %s", typeId)
	}
	for channel := range chanMap {
		go func() {
			channel <- event
		}()
	}
	return nil
}

func (bus EventRegistry) Listen(typeId string, receiver domain.Receiver) error {
	channels, ok := bus.mem[typeId]
	if !ok {
		bus.mem[typeId] = make(map[chan interface{}]bool)
	}
	var publisher chan interface{}
	for channel, taken := range channels {
		if !taken {
			bus.mem[typeId][channel] = true
			publisher = channel
		}
	}
	if publisher == nil {
		publisher = make(chan interface{})
		bus.mem[typeId][publisher] = true
	}
	go func(publisher chan interface{}) {
		for {
			event := <-publisher
			receiver.Publish(event)
			if _, ok := event.(domain.TidyEvent); ok {
				return
			}
		}
	}(publisher)
	return nil
}
