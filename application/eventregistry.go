package application

import (
	"fmt"
	"sync"
)

type EventRegistry struct {
	wg  *sync.WaitGroup
	mem map[string]chan interface{}
}

func New() EventRegistry {
	return EventRegistry{
		mem: make(map[string]chan interface{}),
		wg:  &sync.WaitGroup{}}
}

func (bus EventRegistry) AddMew(typeId string) {
	bus.wg.Wait()
	bus.wg.Add(1)
	defer bus.wg.Done()
	bus.mem[typeId] = make(chan interface{})
}

func (bus EventRegistry) Fire(typeId string, event interface{}) error {
	bus.wg.Wait()
	bus.wg.Add(1)
	defer bus.wg.Done()
	ch, ok := bus.mem[typeId]
	if !ok {
		return fmt.Errorf("no subscriber for typeId: %s", typeId)
	}
	ch <- event
	bus.wg.Done()
	return nil
}

func (bus EventRegistry) Await(typeId string) (interface{}, error) {
	bus.wg.Wait()
	bus.wg.Add(1)
	defer bus.wg.Done()
	ch, ok := bus.mem[typeId]
	if !ok {
		return nil, fmt.Errorf("no channel for typeId: %s", typeId)
	}
	return <-ch, nil
}

func (bus EventRegistry) Listen(typeId string) (chan interface{}, error) {
	bus.wg.Wait()
	bus.wg.Add(1)
	defer bus.wg.Done()
	ch, ok := bus.mem[typeId]
	if !ok {
		return nil, fmt.Errorf("no channel for typeId: %s", typeId)
	}
	return ch, nil
}
