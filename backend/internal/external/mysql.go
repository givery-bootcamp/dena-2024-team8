package external

import (
	"fmt"
	"myapp/internal/config"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Database Setup
// !!! You have to call this function after config setup
func SetupDB() {
	host := config.DBHostName
	port := config.DBPort
	dbname := config.DBName
	dbpass := config.DBPassword
	dbuser := config.DBUser
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, host, port, dbname)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	DB = db.Debug()
}
