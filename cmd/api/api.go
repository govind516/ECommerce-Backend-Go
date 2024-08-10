package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Govind516/E-Commerce-Backend/service/user"
	"github.com/gorilla/mux"
)

type APIserver struct{
	addr string
	db *sql.DB
}

func NewAPIserver(address string, database *sql.DB) *APIserver{
	return &APIserver{
		addr : address,
		db : database,
	}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on" , s.addr)
	return http.ListenAndServe(s.addr, router)	
}