package users

import (
	"nwd/model/request"
	urepo "nwd/repository/users"
)

type users struct {
	userRepo urepo.IUserRepo
}

type Iusers interface {
	Login(req request.Login) (string, error)
}

var u users = users{}

func GetUsers() Iusers {
	return &u
}

func Init(repo urepo.IUserRepo) {
	u.userRepo = repo
}
