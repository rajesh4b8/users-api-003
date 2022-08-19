package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		fmt.Println("Something wrong!")
	}
}
