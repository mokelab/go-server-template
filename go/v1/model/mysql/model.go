package mysql

import (
	m "../"
	"log"
	"os"
)

func NewModels(connection *Connection) *m.Models {
	return &m.Models{
		Logger: log.New(os.Stdout, "logger:", log.Lshortfile),
	}
}
