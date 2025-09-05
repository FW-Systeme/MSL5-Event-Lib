package application

type Aggregate[T any] interface {
	Ident() string
	// Loads the current state internly
	load(entity T)
	// Performs an Event on the current state
	Handle(event interface{}) error
	// Publishes the current state to the outside
	Get() T
}
