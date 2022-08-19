package app

import "github.com/rajesh4b8/users-api-003/src/controllers/ping"

func mapUrls() {
	router.HandleFunc("/ping", ping.PingHandler)
}
