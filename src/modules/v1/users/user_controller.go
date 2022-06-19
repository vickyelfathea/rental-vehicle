package users

import ("net/http"
		"fmt"
		"encoding/json"
		"github.com/gorilla/mux"
		"strconv"
		"reflect"
		"carRent/src/database/gorm/models"
		"carRent/src/interfaces"
	)

type user_ctrl struct {
	repo interfaces.UserService
}

func NewCtrl(rep interfaces.UserService) *user_ctrl {
	return &user_ctrl{rep}
}

func (rep *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindAll()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *user_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User
	json.NewDecoder(r.Body).Decode(&data)
	
	result, err := rep.repo.Add(&data)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (rep *user_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := r.URL.Query()

	var data models.User
	id, err := strconv.Atoi(vars["id"][0]) 

	data.ID = id
	data.Firstname = vars["firstname"][0]
	data.Lastname = vars["lastname"][0]

	if err != nil {
		fmt.Println(id, err, reflect.TypeOf(id))
	}

	result, err := rep.repo.Update(&data)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}


func (rep *user_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) 

	if err != nil {
		fmt.Println(id, err, reflect.TypeOf(id))
	}

	// fmt.Fprintf(w, "Params : %v", id["id"])

	// var data User
	// json.NewDecoder(r.Body).Decode(&data)

	// fmt.Println(data)
	
	result, err := rep.repo.Delete(id)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}