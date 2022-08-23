package user

import (
	"encoding/json"
	"net/http"

	"github.com/rajesh4b8/users-api-003/src/domain/users"
	"github.com/rajesh4b8/users-api-003/src/services"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// error handling
		return
	}

	var u *users.User
	u, err = services.CreateUser(user)

	if err != nil {
		// error handling
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(u)
}
