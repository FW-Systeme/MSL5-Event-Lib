package application

import (
	"fmt"
)

type EventRegistry struct {
	mem map[string]chan interface{}
}

func New() EventRegistry {
	return EventRegistry{mem: make(map[string]chan interface{})}
}

func (bus EventRegistry) AddMew(typeId string) {
	bus.mem[typeId] = make(chan interface{})
}

func (bus EventRegistry) Fire(typeId string, event interface{}) error {
	ch, ok := bus.mem[typeId]
	if !ok {
		return fmt.Errorf("no subscriber for typeId: %s", typeId)
	}
	ch <- event
	return nil
}

func (bus EventRegistry) Await(typeId string) (interface{}, error) {
	ch, ok := bus.mem[typeId]
	if !ok {
		return nil, fmt.Errorf("no channel for typeId: %s", typeId)
	}
	return <-ch, nil
}
