package model

import (
	"database/sql"
	"fmt"

	"github.com/hydeenoble/mux-rest-api/config"
	_ "github.com/lib/pq"
)

// var db
var (
	host     = config.GetConfig().Postgres.Host
	port     = config.GetConfig().Postgres.Port
	user     = config.GetConfig().Postgres.Username
	password = config.GetConfig().Postgres.Password
	dbname   = config.GetConfig().Postgres.DbName
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
