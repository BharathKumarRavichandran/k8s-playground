package models

import (
	"github.com/gocql/gocql"
)

type Record struct {
	id          gocql.UUID `json:"id"`
	message     string     `json:"message"`
	createdDate string     `json:"createdDate"`
	updatedDate string     `json:"updatedDate"`
}
