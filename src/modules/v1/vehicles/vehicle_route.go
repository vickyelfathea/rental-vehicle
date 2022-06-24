package vehicles

import (
	// "net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"carRent/src/middleware")

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles").Subrouter()

	repo := NewRepo(db)  
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/update/{id}", middleware.Do(ctrl.UpdateData, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/delete/{id}", ctrl.DeleteData).Methods("DELETE")
	route.HandleFunc("/{id}", ctrl.GetById).Methods("GET")
}