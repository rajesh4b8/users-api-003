package ping

import (
	"encoding/json"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Pong!")
}
