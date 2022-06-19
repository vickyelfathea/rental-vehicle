package auth

import (
	"encoding/json"
	"net/http"

	"carRent/src/database/gorm/models"
	"carRent/src/interfaces"
	// "fmt"
)

type auth_ctrl struct {
	rep interfaces.AuthService
}

func NewCtrk(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (re *auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)
	// fmt.Println(data.Password)

	result, err := re.rep.Login(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)

}