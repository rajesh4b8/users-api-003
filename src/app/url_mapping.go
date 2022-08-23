package app

import (
	"github.com/rajesh4b8/users-api-003/src/controllers/ping"
	"github.com/rajesh4b8/users-api-003/src/controllers/user"
)

func mapUrls() {
	router.HandleFunc("/ping", ping.PingHandler)

	// Create a user
	router.HandleFunc("/users", user.CreateUserHandler).Methods("POST")
}
