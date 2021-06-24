package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type DBData struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

func (d *DBData) ConnectToDB() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		d.Server, d.User, d.Password, d.Port, d.Database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Println("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connected!\n")
	return db
}
