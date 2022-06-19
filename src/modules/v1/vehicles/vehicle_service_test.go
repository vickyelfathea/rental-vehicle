package vehicles

import (
	"testing"


	"carRent/src/modules/v1/vehicles/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"

	"carRent/src/database/gorm/models"

	// help "carRent/src/helpers"
	// "carRent/src/interfaces"
)

var modelMock = models.Vehicle{
	ID: 1,
}

var modelMocks = models.Vehicles{
	models.Vehicle {ID: 1},
	models.Vehicle {ID: 2},
}


func TestFindId(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = vehicles_service{&repo}
 
	repo.Mock.On("FindId", 1).Return(&modelMock, nil)
	data, err := service.FindId(1)

	vehicles := data.Data.(*models.Vehicle)
	assert.Equal(t, 1, vehicles.ID, "Expect id = 1")
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = vehicles_service{&repo}
 
	repo.Mock.On("FindAll").Return(&modelMocks, nil)
	data, err := service.FindAll()

	vehicles := data.Data.(*models.Vehicles)
	
	expectedModelMocks := &models.Vehicles{
		models.Vehicle {ID: 1},
		models.Vehicle {ID: 2},
	}

	assert.Equal(t, expectedModelMocks, vehicles, "Expect id = 1 & 2")
	assert.Nil(t, err)
	
}





















// type vehicles_service struct {
// 	rep interfaces.VehicleRepo
// }

// func NewService(rep interfaces.VehicleRepo) *vehicles_service {
// 	return &vehicles_service{rep}
// }

// func (re *vehicles_service) FindAll() (*help.Response, error) {
// 	data, err := re.rep.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	respone := help.New(data, 200, false)
// 	return respone, nil
// }

// func (re *vehicles_service) Add(vehicle *models.Vehicle) (*help.Response, error) {
// 	data, err := re.rep.Add(vehicle)
// 	if err != nil {
// 		return nil, err
// 	}

// 	respone := help.New(data, 200, false)
// 	return respone, nil
// }

// func (re *vehicles_service) Update(body *models.Vehicle) (*help.Response, error) {
// 	data, err := re.rep.Update(body)
	
// 	if err != nil {
// 		return nil, err
// 	}

// 	respone := help.New(data, 200, false)
// 	return respone, nil
// }

// func (re *vehicles_service) Delete(id int) (*help.Response, error) {
// 	data, err := re.rep.Delete(id)
	
// 	if err != nil {
// 		return nil, err
// 	}

// 	respone := help.New(data, 200, false)
// 	return respone, nil
// }

// func (re *vehicles_service) FindId(id int) (*help.Response, error) {
// 	data, err := re.rep.FindId(id)
	
// 	if err != nil {
// 		return nil, err
// 	}

// 	respone := help.New(data, 200, false)
// 	return respone, nil
// }