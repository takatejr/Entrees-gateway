package datebase

import (
	"database/sql"
	"fmt"
	"gateway/utils"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "users"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	//defer DB.Close()

	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)

	utils.DefaultErrorHandler(err)

	err = db.Ping()

	utils.DefaultErrorHandler(err)

	logrus.Println("Datebase ready.")

	return db
}
