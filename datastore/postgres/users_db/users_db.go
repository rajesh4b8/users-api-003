package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
		fmt.Printf("DB driver failed with %s\n", err.Error())
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		fmt.Printf("DB connection failed with %s\n", err.Error())
		panic(err)
	}

	fmt.Println("DB connection successful!")
}
