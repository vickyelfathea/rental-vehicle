package vehicles

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
	"carRent/src/interfaces"
)


type vehicles_service struct {
	rep interfaces.VehicleRepo
}

func NewService(rep interfaces.VehicleRepo) *vehicles_service {
	return &vehicles_service{rep}
}

func (re *vehicles_service) FindAll() (*help.Response, error) {
	data, err := re.rep.FindAll()
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *vehicles_service) Add(vehicle *models.Vehicle) (*help.Response, error) {
	data, err := re.rep.Add(vehicle)
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *vehicles_service) Update(body *models.Vehicle) (*help.Response, error) {
	data, err := re.rep.Update(body)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *vehicles_service) Delete(id int) (*help.Response, error) {
	data, err := re.rep.Delete(id)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}

func (re *vehicles_service) FindId(id int) (*help.Response, error) {
	data, err := re.rep.FindId(id)
	
	if err != nil {
		return nil, err
	}

	respone := help.New(data, 200, false)
	return respone, nil
}