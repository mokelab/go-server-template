package mysql

import (
	m "../"
	"log"
)

func NewModels(connection *Connection, logger *log.Logger) *m.Models {
	return &m.Models{
		Logger: logger,
	}
}
