package services

import "github.com/stepup99/microservices-p1/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
