package history

import (
	// "net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"carRent/src/middleware")

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db) 
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.Do(ctrl.GetAll, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/user/{id}", ctrl.GetById).Methods("GET")
	route.HandleFunc("/popular", ctrl.GetPopular).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	// route.HandleFunc("/update", ctrl.UpdateData).Methods("PUT")
	route.HandleFunc("/delete/{id}", ctrl.DeleteData).Methods("DELETE")
} 