package users

import (
	"nwd/model/request"
	"nwd/shared/database"
)

type userRepo struct {
}

type IUserRepo interface {
	Login(db database.IConnection, req request.Login) (bool, error)
}

var w = userRepo{}

func GetWaitingRepo() IUserRepo {
	return &w
}
