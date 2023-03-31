package users

import (
	"fmt"

	"nwd/model/request"
	"nwd/model/tables"
	"nwd/shared/database"
)

func (u *userRepo) Login(db database.IConnection, req request.Login) (bool, error) {
	var user tables.Users
	err := db.Tables("users").Where("username = ? and password =? ", req.UserName, req.Password).First(&user).Error
	if err != nil {
		fmt.Printf("User Login error: %v", err)
		return false, err
	}

	if user.ID == 0 {
		fmt.Printf("User is null")
		return false, err
	}

	return true, nil
}
