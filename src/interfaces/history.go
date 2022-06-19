package interfaces

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
)

type HistoryRepo interface {
	FindAll() (*models.Histories, error)
	FindId(id int) (*models.Results, error)
	FindPopular() (*models.Populars, error)
	Add(*models.History) (*models.History, error)
	Update(*models.History) (*models.History, error)
	Delete(id int) (*models.History, error)
}

type HistoryService interface {
	FindAll() (*help.Response, error)
	FindId(id int) (*help.Response, error)
	FindPopular() (*help.Response, error)
	Add(*models.History) (*help.Response, error)
	Update(*models.History) (*help.Response, error)
	Delete(id int) (*help.Response, error)
}