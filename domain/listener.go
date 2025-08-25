package domain

type Receiver func(event interface{}) error
