package v1

import (
	m "./model/mysql"
	"./rest"
	s "./service/impl"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func InitRouter(root *mux.Router) error {
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
	models := m.NewModels(c)
	services := s.NewServices(models)
	r := root.PathPrefix("/api/v1").Subrouter()
	rest.SetHandlers(r, services)
	return nil
}
