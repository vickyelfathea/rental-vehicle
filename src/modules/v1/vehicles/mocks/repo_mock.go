package mocks

import ("github.com/stretchr/testify/mock"
		"carRent/src/database/gorm/models")

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindId (id int) (*models.Vehicle, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}

func (pr *RepoMock) FindAll() (*models.Vehicles, error){
	args := pr.Mock.Called()
	return args.Get(0).(*models.Vehicles), args.Error(1)
}

func (pr *RepoMock) Add(data *models.Vehicle) (*models.Vehicle, error) {
	return data, nil
}
	
func (pr *RepoMock) Update(data *models.Vehicle) (*models.Vehicle, error){
	return data, nil
}

func (pr *RepoMock) Delete(id int) (*models.Vehicle, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Vehicle), args.Error(1)
}
