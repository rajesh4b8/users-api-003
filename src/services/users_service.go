package services

import "github.com/rajesh4b8/users-api-003/src/domain/users"

func CreateUser(u users.User) (*users.User, error) {
	// send dummy response
	u.Save()

	return &u, nil
}

func GetUser(userId int) (*users.User, error) {
	return users.GetUserById(userId)
}

func DeleteUser(userId int) (*users.User, error) {
	return users.DeleteUserById(userId)
}
