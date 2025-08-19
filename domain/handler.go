package domain

type EventHandler[T any] interface {
	Handle(event T) error
}
