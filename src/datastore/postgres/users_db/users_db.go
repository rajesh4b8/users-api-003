package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rajesh4b8/users-api-003/src/logger"
)

const (
	username = "postgres"
	password = "dev"
	host     = "127.0.0.1"
	schema   = "postgres"
)

var (
	Client *sql.DB
)

func init() {
	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("postgres", datasource)
	if err != nil {
		logger.GetLogger().Error("DB driver failed with " + err.Error())
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		logger.GetLogger().Error("DB connection failed with " + err.Error())
		panic(err)
	}

	logger.GetLogger().Info("DB connection successful!")
}
