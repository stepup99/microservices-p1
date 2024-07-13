package services

import (
	"github.com/stepup99/microservices-p1/mvc/domain"
	"github.com/stepup99/microservices-p1/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
