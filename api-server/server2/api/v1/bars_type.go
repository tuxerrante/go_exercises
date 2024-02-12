package v1

import (
	"github.com/google/uuid"
)

type Bar struct {
	Name string    `json:"name"`
	Id   uuid.UUID `json:"id"`
}

type BarList struct {
	Id map[uuid.UUID]*Bar `json:"id"`
}
