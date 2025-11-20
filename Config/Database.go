package Config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     57329,
		User:     "sa",
		Password: "1234",
		DBName:   "first_go",
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%d?database=%s&instance=SQLEXPRESS",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
