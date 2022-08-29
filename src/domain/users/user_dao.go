package users

import (
	"fmt"

	"github.com/rajesh4b8/users-api-003/src/utils/date_util"
	"github.com/rajesh4b8/users-api-003/src/utils/error_util"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() *error_util.RestErr {
	user.DateCreated = date_util.TimeNow()
	usersDB[user.Id] = user

	return nil
}

func GetUserById(id int) (*User, *error_util.RestErr) {
	user, ok := usersDB[id]
	if !ok {
		return nil, error_util.NewNotFoundError(fmt.Sprintf("user not found with id: %v", id))
	}

	return user, nil
}
func DeleteUserById(id int) (*User, error) {
	user, ok := usersDB[id]

	if !ok {
		return nil, fmt.Errorf("User with id %v not found", id)
	}
	delete(usersDB, user.Id)
	fmt.Println("deleted user", user.Id)

	return user, nil
}
