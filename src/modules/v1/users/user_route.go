package users

import (
	// "net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	// "carRent/src/middleware"
	)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db) 
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/update", ctrl.UpdateData).Methods("PUT")
	route.HandleFunc("/delete/{id}", ctrl.DeleteData).Methods("DELETE")
}