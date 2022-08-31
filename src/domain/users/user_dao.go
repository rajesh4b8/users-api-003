package users

import (
	"fmt"
	"strings"

	"github.com/rajesh4b8/users-api-003/src/datastore/postgres/users_db"
	"github.com/rajesh4b8/users-api-003/src/utils/error_util"
)

const (
	noRowsError   = "no rows in result set"
	queryUserById = "select user_id, first_name, last_name, email, date_created from users where user_id = $1"
	querySaveUser = "insert into users (first_name, last_name, email) values ($1, $2, $3) returning user_id, date_created"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Save() *error_util.RestErr {
	// user.DateCreated = date_util.TimeNow()
	// usersDB[user.Id] = user

	stmt, err := users_db.Client.Prepare(querySaveUser)
	if err != nil {
		return error_util.NewInternalServerError(err.Error())
	}

	// defer will call the close function before returning from the current function
	defer stmt.Close()

	result := stmt.QueryRow(user.FirstName, user.LastName, user.EmailId)

	if err := result.Scan(&user.Id, &user.DateCreated); err != nil {
		return error_util.NewInternalServerError(err.Error())
	}

	return nil
}

func GetUserById(id int) (*User, *error_util.RestErr) {
	// user, ok := usersDB[id]
	// if !ok {
	// 	return nil, error_util.NewNotFoundError(fmt.Sprintf("user not found with id: %v", id))
	// }

	stmt, err := users_db.Client.Prepare(queryUserById)
	if err != nil {
		return nil, error_util.NewInternalServerError(err.Error())
	}

	// defer will call the close function before returning from the current function
	defer stmt.Close()

	user := User{}
	result := stmt.QueryRow(id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.EmailId, &user.DateCreated); err != nil {

		if strings.Contains(err.Error(), noRowsError) {
			return nil, error_util.NewNotFoundError(fmt.Sprintf("User not found for id %d", id))
		}

		return nil, error_util.NewInternalServerError(err.Error())
	}

	return &user, nil
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
