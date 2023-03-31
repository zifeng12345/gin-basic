package users

import srv "nwd/service/users"

var uctl userCtl

type userCtl struct {
	service srv.Iusers
}

func GetUserCtl() *userCtl {
	return &uctl
}

func Init(srv srv.Iusers) {
	uctl.service = srv
}
