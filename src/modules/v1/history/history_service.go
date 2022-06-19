package history

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
	"carRent/src/interfaces"
)


type histories_service struct {
	rep interfaces.HistoryRepo
}

func NewService(rep interfaces.HistoryRepo) *histories_service {
	return &histories_service{rep}
}

	func (re *histories_service) FindAll() (*help.Response, error){
		data, err := re.rep.FindAll()
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}

	func (re *histories_service) FindId(id int) (*help.Response, error){
		data, err := re.rep.FindId(id)
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}

	func (re *histories_service) FindPopular() (*help.Response, error){
		data, err := re.rep.FindPopular()
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}

	func (re *histories_service) Add(history *models.History) (*help.Response, error){
		data, err := re.rep.Add(history)
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}

	func (re *histories_service) Update(body *models.History) (*help.Response, error){
		data, err := re.rep.Update(body)
	
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}

	func (re *histories_service) Delete(id int) (*help.Response, error){
		data, err := re.rep.Delete(id)
	
		if err != nil {
			return nil, err
		}

		respone := help.New(data, 200, false)
		return respone, nil
	}