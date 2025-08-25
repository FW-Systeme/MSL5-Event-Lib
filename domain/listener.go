package domain

type Receiver interface {
	Publish(event interface{}) error
}
