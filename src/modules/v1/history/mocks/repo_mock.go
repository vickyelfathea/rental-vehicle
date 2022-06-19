package mocks

import ("github.com/stretchr/testify/mock"
		"carRent/src/database/gorm/models")

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindId(id int) (*models.Results, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.Results), args.Error(1)
}

func (pr *RepoMock) FindAll() (*models.Histories, error){
	args := pr.Mock.Called()
	return args.Get(0).(*models.Histories), args.Error(1)
}

func (pr *RepoMock) FindPopular() (*models.Populars, error){
	args := pr.Mock.Called()
	return args.Get(0).(*models.Populars), args.Error(1)
}

func (pr *RepoMock) Add(data *models.History) (*models.History, error) {
	return data, nil
}
	
func (pr *RepoMock) Update(data *models.History) (*models.History, error){
	return data, nil
}

func (pr *RepoMock) Delete(id int) (*models.History, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.History), args.Error(1)
}
