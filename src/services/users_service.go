package services

import (
	"github.com/rajesh4b8/users-api-003/src/domain/users"
	"github.com/rajesh4b8/users-api-003/src/utils/error_util"
)

func CreateUser(u users.User) (*users.User, *error_util.RestErr) {
	// Validate input request
	// Validate if emailId is valid
	// if not validate return a 400 error
	// TODO: Assignment for validating the User

	err := u.Save()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(userId int) (*users.User, *error_util.RestErr) {
	return users.GetUserById(userId)
}

func DeleteUser(userId int) (*users.User, error) {
	return users.DeleteUserById(userId)
}
