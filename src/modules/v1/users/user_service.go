package users

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
	"carRent/src/interfaces"
	"github.com/go-playground/validator/v10"
)


type users_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *users_service {
	return &users_service{rep}
}

func (re *users_service) FindAll() (*help.Response, error) {
	data, err := re.rep.FindAll()
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) Add(user *models.User) (*help.Response, error) {
	validate := validator.New()
	
	err := validate.Struct(user)
	if err != nil {
		return nil, err
	}
	
	hassPassword, err := help.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hassPassword
	
	data, err := re.rep.Add(user)
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) Update(body *models.User) (*help.Response, error) {
	data, err := re.rep.Update(body)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) Delete(id int) (*help.Response, error) {
	data, err := re.rep.Delete(id)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) FindId(id int) (*help.Response, error) {
	data, err := re.rep.FindId(id)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *users_service) FindByUsername(username string) (*help.Response, error) {
	data, err := re.rep.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}