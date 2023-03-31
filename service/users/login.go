package users

import (
	"fmt"

	"nwd/middleware"
	"nwd/model/request"
	"nwd/shared/database"
)

func (u *users) Login(req request.Login) (string, error) {
	var res string
	success, err := u.userRepo.Login(database.GetConnection(), req)
	if err != nil {
		return res, err
	}

	if success {
		res, flag := middleware.GenerateJwt(req.UserName, req.Password, 24*30)
		if flag {
			fmt.Println("Generate jwt success")
			return res, nil
		}
	}

	return "", nil
}
