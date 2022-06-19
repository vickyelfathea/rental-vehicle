package interfaces

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	Add(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
	FindId(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

type UserService interface {
	FindAll() (*help.Response, error)
	Add(*models.User) (*help.Response, error)
	Update(*models.User) (*help.Response, error)
	Delete(id int) (*help.Response, error)
	FindId(id int) (*help.Response, error)
	FindByUsername(username string) (*help.Response, error)
}