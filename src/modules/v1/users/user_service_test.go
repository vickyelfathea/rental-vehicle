package users

import (
	"testing"

	"carRent/src/modules/v1/users/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"

	"carRent/src/database/gorm/models"
)

var modelMock = models.User{
	ID: 1,
	Username: "admin",
}

var modelMocks = models.Users{
	models.User {ID: 1},
	models.User {ID: 2},
}


func TestFindId(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = users_service{&repo}
 
	repo.Mock.On("FindId", 1).Return(&modelMock, nil)
	data, err := service.FindId(1)

	users := data.Data.(*models.User)
	assert.Equal(t, 1, users.ID, "Expect id = 1")
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = users_service{&repo}
 
	repo.Mock.On("FindAll").Return(&modelMocks, nil)
	data, err := service.FindAll()

	users := data.Data.(*models.Users)
	
	expectedModelMocks := &models.Users{
		models.User {ID: 1},
		models.User {ID: 2},
	}

	assert.Equal(t, expectedModelMocks, users, "Expect id = 1 & 2")
	assert.Nil(t, err)
	
}

func TestFindByUSername(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = users_service{&repo}
 
	repo.Mock.On("FindByUsername", "admin").Return(&modelMock, nil)
	data, err := service.FindByUsername("admin")

	users := data.Data.(*models.User)
	assert.Equal(t, "admin", users.Username, "Expect username = admin")
	assert.Nil(t, err)	
}