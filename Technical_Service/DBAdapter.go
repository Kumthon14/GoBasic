package technicalservice

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connectionString = "postgres://postgres:1234@127.0.0.1/Golang"

var globalAdapterInstance *gorm.DB

type DBAdapter struct {
}

func connectToDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return db
}

func (a *DBAdapter) GetGorm() *gorm.DB {
	a.StaticGetAdapterInstance()
	return globalAdapterInstance
}

func (a *DBAdapter) StaticGetAdapterInstance() *DBAdapter {
	if globalAdapterInstance == nil {
		globalAdapterInstance = connectToDatabase()
		a = &DBAdapter{}
	}
	return a
}
