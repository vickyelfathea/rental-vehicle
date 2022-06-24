package routers

import (
	"carRent/src/database"
	"carRent/src/modules/v1/auth"
	"carRent/src/modules/v1/history"
	"carRent/src/modules/v1/users"
	"carRent/src/modules/v1/vehicles"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
	// "log"
	// "fmt"
	// "encoding/json"
)
	

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("car-rent-vicky"),
		newrelic.ConfigLicense("006f5cce8ce3c4b6d3183d36a88f2a840359NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}  

	mainRoute.Use(nrgorilla.Middleware(nRelic))

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/relic", relicHandler)).Methods("GET")

	db, err := database.New()
	if err != nil {
		return nil, err
	}
	
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	history.New(mainRoute, db)
	auth.New(mainRoute, db)
	

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}

func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}
