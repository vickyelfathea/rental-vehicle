package history

import ("net/http"
		"fmt"
		"encoding/json"
		"github.com/gorilla/mux"
		"strconv"
		"reflect"
		"carRent/src/database/gorm/models"
		"carRent/src/interfaces"
	)

// var respone helpers.Respone

type history_ctrl struct {
	repo interfaces.HistoryService
}

func NewCtrl(rep interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{rep}
}

func (rep *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, err := rep.repo.FindAll()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *history_ctrl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

func (rep *history_ctrl) GetPopular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := rep.repo.FindPopular()

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *history_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.History
	json.NewDecoder(r.Body).Decode(&data)
	
	result, err := rep.repo.Add(&data)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

// func (rep *history_ctrl) UpdateData(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	vars := r.URL.Query()

// 	var data History
// 	id, err := strconv.Atoi(vars["id"][0]) 

// 	data.ID = id
// 	data.Name = vars["name"][0]
// 	data.City = vars["city"][0]
// 	data.Type = vars["type"][0]

// 	if err != nil {
// 		fmt.Println(id, err, reflect.TypeOf(id))
// 	}

// 	result, err := rep.repo.Update(&data)

// 	if err != nil {
// 		fmt.Fprint(w, err.Error())
// 	}

// 	json.NewEncoder(w).Encode(&result)
// }


func (rep *history_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
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