package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"carRent/src/database/gorm/models"
	"carRent/src/modules/v1/history/mocks"
)

var modelMockResult = models.Results{
	models.Result {ID: 1, Firstname: "dua"},
	models.Result {ID: 2, Firstname: "lipa"},
}

var modelMockPopular = models.Populars{
	models.Popular {Vehicleid: 1},
	models.Popular {Vehicleid: 2},
}

var modelMocks = models.Histories{
	models.History {ID: 1},
	models.History {ID: 2},
}

func TestFindId(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = histories_service{&repo}
 
	repo.Mock.On("FindId", 1).Return(&modelMockResult, nil)
	data, err := service.FindId(1)

	var modelMockResult = &models.Results{
		models.Result {ID: 1, Firstname: "dua"},
		models.Result {ID: 2, Firstname: "lipa"},
	}

	history := data.Data.(*models.Results)
	assert.Equal(t, modelMockResult, history, "Expect id = 1 & 2")
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = histories_service{&repo}
 
	repo.Mock.On("FindAll").Return(&modelMocks, nil)
	data, err := service.FindAll()

	history := data.Data.(*models.Histories)
	
	expectedModelMocks := &models.Histories{
		models.History {ID: 1},
		models.History {ID: 2},
	}

	assert.Equal(t, expectedModelMocks, history, "Expect id = 1 & 2")
	assert.Nil(t, err)
	
}

func TestFindPopular(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = histories_service{&repo}
 
	repo.Mock.On("FindPopular").Return(&modelMockPopular, nil)
	data, err := service.FindPopular()

	var modelMockPopular = &models.Populars{
		models.Popular {Vehicleid: 1},
		models.Popular {Vehicleid: 2},
	}

	history := data.Data.(*models.Populars)
	assert.Equal(t, modelMockPopular, history, "Expect username = admin")
	assert.Nil(t, err)	
}