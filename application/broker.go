package application

import "github.com/FW-Systeme/MSL5-Event-Lib/domain"

type Broker interface {
	Publish(typeId string, event interface{}) error
	Listen(typeId string, handler domain.Handler) error
}
