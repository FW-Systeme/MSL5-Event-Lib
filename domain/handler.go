package domain

type Handler func(event interface{}) error
type ErrorHandler func(event interface{}, err error)
