package application

import (
	"fmt"
	"sync"

	"github.com/FW-Systeme/MSL5-Event-Lib/domain"
)

type EventBroker struct {
	wg     *sync.WaitGroup
	mu     sync.RWMutex
	evtMem map[string][]chan interface{}
}

func New() EventBroker {
	return EventBroker{
		evtMem: make(map[string][]chan interface{}),
		wg:     &sync.WaitGroup{},
		mu:     sync.RWMutex{}}
}

func (bus EventBroker) Publish(typeId string, event interface{}) error {
	bus.mu.RLock()
	chanMap, ok := bus.evtMem[typeId]
	bus.mu.RUnlock()
	if !ok {
		return fmt.Errorf("no subscriber for typeId: %s", typeId)
	}
	for _, listener := range chanMap {
		go func() {
			listener <- event
		}()
	}
	return nil
}

func (bus EventBroker) Listen(typeId string, handler domain.Handler) error {
	bus.mu.Lock()
	_, ok := bus.evtMem[typeId]
	if !ok {
		bus.evtMem[typeId] = make([]chan interface{}, 1)
	}
	publisher := make(chan interface{})
	bus.evtMem[typeId] = append(bus.evtMem[typeId], publisher)
	bus.mu.Unlock()
	go func(publisher chan interface{}) {
		for {
			event := <-publisher
			if err := handler(event); err != nil {
			}
			if _, ok := event.(domain.TidyEvent); ok {
				return
			}
		}
	}(publisher)
	return nil
}
