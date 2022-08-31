package error_util

import (
	"encoding/json"
	"net/http"
)

type RestErr struct {
	Status  int    `json:"Status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (err *RestErr) HandleError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

func NewBadRequestError(message string) *RestErr {
	err := RestErr{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "Bad Request",
	}

	return &err
}

func NewNotFoundError(message string) *RestErr {
	err := RestErr{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "Not Found",
	}

	return &err
}

func NewInternalServerError(message string) *RestErr {
	err := RestErr{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "Internal Server Error",
	}

	return &err
}
