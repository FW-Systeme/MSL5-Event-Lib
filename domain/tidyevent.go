package domain

import (
	"github.com/google/uuid"
)

const (
	TidyEventId = "Tidy Evnet"
)

type TidyEvent struct {
	Uuid uuid.UUID
}
