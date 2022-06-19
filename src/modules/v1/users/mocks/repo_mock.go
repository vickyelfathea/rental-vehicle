package mocks

import ("github.com/stretchr/testify/mock"
		"carRent/src/database/gorm/models")

type RepoMock struct {
	Mock mock.Mock
}

func (pr *RepoMock) FindId(id int) (*models.User, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) FindAll() (*models.Users, error){
	args := pr.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr *RepoMock) FindByUsername(username string) (*models.User, error){
	args := pr.Mock.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr *RepoMock) Add(data *models.User) (*models.User, error) {
	return data, nil
}
	
func (pr *RepoMock) Update(data *models.User) (*models.User, error){
	return data, nil
}

func (pr *RepoMock) Delete(id int) (*models.User, error){
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}
