package users

import (
	"fmt"

	"github.com/rajesh4b8/users-api-003/datastore/postgres/users_db"
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
	// user, ok := usersDB[id]
	// if !ok {
	// 	return nil, error_util.NewNotFoundError(fmt.Sprintf("user not found with id: %v", id))
	// }

	stmt, err := users_db.Client.Prepare("select user_id, first_name, last_name, email, date_created from users where user_id = $1")
	if err != nil {
		// handle error
	}
	defer stmt.Close()

	user := User{Id: id}
	result := stmt.QueryRow(id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.EmailId, &user.DateCreated); err != nil {

		// if strings.Contains(err.Error(), noRowsError) {
		// 	return errors.NewNotFoundError(fmt.Sprintf("User not found for id %d", user.Id))
		// }

		// logger.Error(fmt.Sprintf("Error while fetching user %d", user.Id), err)
		// return errors.NewInternalServerError("Database error")
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
