package main

import (
	"GoPractice/Config"
	"GoPractice/Models"
	"GoPractice/Routes"
	_ "GoPractice/Technical_Service"
	"fmt"
	_ "strings"
	_ "text/template"

	_ "GoPractice/docs"

	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var err error

// @title Go Hello Api
// @version 1.0
// @destination Go Learning Project
// @termOfService http://agilerap.com/

// @contact.name API Support
// @contact.url http://agilerap.com
// @contact.email support@agilerap.com

// @license.name Agilerap
// @license.url http://agilerap.com/

// @schemas https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// var adapter *technicalservice.DBAdapter

	// a := adapter.GetGorm()

	// a.AutoMigrate()

	// db := adapter.StaticGetAdapterInstance()

	// fmt.Println("Database instance: ", db)

	// dsn := Config.DbURL(&Config.DBConfig{})

	dbConfig := Config.BuildDBConfig()
	dsn := Config.DbURL(dbConfig)
	Config.DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Status: ", err)
	}

	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetUpRouter()

	r.Run()
}
