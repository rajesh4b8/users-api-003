package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rajesh4b8/users-api-003/src/domain/users"
	"github.com/rajesh4b8/users-api-003/src/services"
	"github.com/rajesh4b8/users-api-003/src/utils/error_util"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// error handling - bad json
		error_util.NewBadRequestError("Body must be a valid json").HandleError(w)
		return
	}

	u, restErr := services.CreateUser(user)
	if restErr != nil {
		// error handling
		restErr.HandleError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		// userId is not a number
		error_util.NewBadRequestError("userId must be a number").HandleError(w)

		return
	}

	user, getUserErr := services.GetUser(userId)
	if getUserErr != nil {
		// error handling
		getUserErr.HandleError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		// error handling
		return
	}

	user, delUserErr := services.DeleteUser(userId)
	if delUserErr != nil {
		// error handling
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}
