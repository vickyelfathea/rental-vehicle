package routers

import ("net/http"
	"github.com/gorilla/mux"
	"carRent/src/modules/v1/products"
	"carRent/src/modules/v1/users"
	"carRent/src/modules/v1/vehicles"
	"carRent/src/modules/v1/history"
	"carRent/src/modules/v1/auth"
	"carRent/src/database"
	// "log"
	// "fmt"
	// "encoding/json"
	)
	

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	db, err := database.New()

	if err != nil {
		return nil, err
	}
	
	products.New(mainRoute, db)
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	history.New(mainRoute, db)
	auth.New(mainRoute, db)
	

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}
