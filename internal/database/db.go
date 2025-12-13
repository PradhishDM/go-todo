package database

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"project.com/todo/internal/config"
)


func ConnectDB(config *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName )

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return db,nil



}