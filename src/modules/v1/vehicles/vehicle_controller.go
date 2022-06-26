package vehicles

import ("net/http"
		"fmt"
		"encoding/json"
		"github.com/gorilla/mux"
		"strconv"
		"reflect"
		"carRent/src/database/gorm/models"
		"carRent/src/interfaces"
	)

type vehicle_ctrl struct {
	repo interfaces.VehicleService
}

func NewCtrl(rep interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{rep}
}

func (rep *vehicle_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, err := rep.repo.FindAll()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Vehicle

	json.NewDecoder(r.Body).Decode(&data)
	result, err := rep.repo.Add(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(result)
}

func (rep *vehicle_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) 

	if err != nil {
		fmt.Println(id, err, reflect.TypeOf(id))
	}

	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	data.ID = id

	result, err := rep.repo.Update(&data)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}


func (rep *vehicle_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) 

	if err != nil {
		fmt.Println(id, err, reflect.TypeOf(id))
	}
	
	result, err := rep.repo.Delete(id)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *vehicle_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) 

	if err != nil {
		fmt.Println(id, err, reflect.TypeOf(id))
	}
	
	result, err := rep.repo.FindId(id)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}