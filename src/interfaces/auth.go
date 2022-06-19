package interfaces

import (
	"carRent/src/database/gorm/models"
	help "carRent/src/helpers"
)

type AuthService interface {
	Login(body models.User) (*help.Response, error)
}