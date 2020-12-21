package common

import (
	"database/sql"
	_ "github.com/lib/pq"
	"errors"
	"fmt"
)

func Connect() (*sql.DB, error) {
	user := Config.DbUser
	password := Config.DbPassword
	host := Config.DbHost
	port := Config.DbPort
	dbname := Config.DbName
	driver := Config.DbDriver

	if driver == "postgres" {
		URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

		db, err := sql.Open(driver, URL)

		if err != nil {
			return nil, err
		}

		return db, nil
	} else {
		return nil, errors.New("Driver is not supported")
	}

}