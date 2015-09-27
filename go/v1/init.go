package v1

import (
	m "./model/mysql"
	"./rest"
	s "./service/impl"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(root *mux.Router, logger *log.Logger) error {
	db, err := sql.Open("mysql", "<user>:<pass>@/<db-name>")
	if err != nil {
		return err
	}
	/*
		c := mi.NewConnection(db)
		models := mi.NewModels(c)
		services := si.NewServices(models)
	*/
	c := m.NewConnection(db)
	models := m.NewModels(c, logger)
	services := s.NewServices(models)
	r := root.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r, services)
	return nil
}
