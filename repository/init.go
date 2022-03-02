package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sample/constants"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "docker"
	host     = "localhost"
	port     = 5432
	dbname   = "postgres"
)

var DBP *sql.DB

func Connect() {
	dbConf := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user,
		password, host, port, dbname)

	db, err := sql.Open("postgres", dbConf)
	if err != nil {
		log.Println("error in connection function on sql.Open part")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("error in connection part and on pinging")
		panic(err)
	}
	fmt.Println("connection...")
	DBP = db
	_, err = db.Exec(constants.CREATE_USER_TABLE)
	if err != nil {
		log.Println("error on creating user table")
		panic(err)
	}
	log.Println("table : user status : created")

}
