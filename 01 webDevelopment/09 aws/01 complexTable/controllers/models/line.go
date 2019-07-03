package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Line struct {
	Id     uuid.UUID
	Text   string
	Number int
	Date   time.Time
}
