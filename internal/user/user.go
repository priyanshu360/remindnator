package user

import (
	"github.com/google/uuid"
	"github.com/priyanshu360/remindnator/internal/source"
)

type User struct {
	ID     uuid.UUID
	Email  string
	Source []*source.Source
}
