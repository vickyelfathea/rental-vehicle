package interfaces

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
)

type VehicleRepo interface {
	FindAll() (*models.Vehicles, error)
	Add(*models.Vehicle) (*models.Vehicle, error)
	Update(*models.Vehicle) (*models.Vehicle, error)
	Delete(id int) (*models.Vehicle, error)
	FindId(id int) (*models.Vehicle, error)
}

type VehicleService interface {
	FindAll() (*help.Response, error)
	Add(*models.Vehicle) (*help.Response, error)
	Update(*models.Vehicle) (*help.Response, error)
	Delete(id int) (*help.Response, error)
	FindId(id int) (*help.Response, error)
}

