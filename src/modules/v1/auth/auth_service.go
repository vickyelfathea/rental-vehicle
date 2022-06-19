package auth

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
	"carRent/src/interfaces"
	// "fmt"
)

type token_respone struct {
	Tokens string `json:"token"`
}

type auth_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *auth_service {
	return &auth_service{rep}
}

func (r *auth_service) Login(body models.User) (*help.Response, error) {
	user, err := r.rep.FindByUsername(body.Username)
	if err != nil {
		return nil, err
	}

	if !help.CheckPassword(user.Password, body.Password) {
		return help.New("Password salah", 401, true), nil
	}



	token := help.NewToken(body.Username)
	// fmt.Println(token)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}

	respone := help.New(token_respone{Tokens: theToken}, 200, false)
	return respone, nil

}