package app

import (
	"github.com/rajesh4b8/users-api-003/src/controllers/ping"
	"github.com/rajesh4b8/users-api-003/src/controllers/user"
)

func mapUrls() {
	router.HandleFunc("/ping", ping.PingHandler)

	// Create a user
	router.HandleFunc("/users", user.CreateUserHandler).Methods("POST")
	// Get user by id
	router.HandleFunc("/users/{userId}", user.GetUserHandler).Methods("GET")

	// Delete user by id
	router.HandleFunc("/users/{userId}", user.DeleteUserHandler).Methods("DELETE")

	// list all the users
	// router.HandleFunc("/listAllUsers", user.ListAllUsersHandler).Methods("GET")
}
