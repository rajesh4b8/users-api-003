package users

import (
	"fmt"

	"github.com/rajesh4b8/users-api-003/src/utils/date_util"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() error {
	user.DateCreated = date_util.TimeNow()
	usersDB[user.Id] = user

	return nil
}

func GetUserById(id int) (*User, error) {
	user, ok := usersDB[id]
	if !ok {
		return nil, fmt.Errorf("User with id %v not found", id)
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
