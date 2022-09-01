package ping

import (
	"encoding/json"
	"net/http"

	"github.com/rajesh4b8/users-api-003/src/logger"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("Responding with pong!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Pong!")
}
