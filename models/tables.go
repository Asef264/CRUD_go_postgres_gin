package models

import (
	"context"
	"database/sql"
	"log"
	"sample/constants"
)

var db *sql.DB

func CreateUserTable(context.Context) error {

	_, err := db.Query(constants.CREATE_USER_TABLE)
	if err != nil {
		log.Println("error on creating user table")
		panic(err)
	}
	log.Println("table : user  status : created")
	return nil
}
