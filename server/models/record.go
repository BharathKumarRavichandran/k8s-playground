package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Record struct {
	ID          gocql.UUID `json:"id"`
	Message     string     `json:"message"`
	CreatedDate time.Time  `json:"created_date"`
}
