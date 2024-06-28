package config

import (
	"fmt"
	"os"
	"strconv"
)

var HostName = "127.0.0.1"
var Port = 9000
var CorsAllowOrigin = "http://localhost:3000"
var DBHostName = "db"
var DBPort = 3306
var DBName = "training"
var DBUser = "root"
var DBPassword = ""
var JwtSecret = ""

func init() {
	if v := os.Getenv("HOSTNAME"); v != "" {
		HostName = v
	}
	if v, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64); err == nil {
		Port = int(v)
	}
	if v := os.Getenv("CORS_ALLOW_ORIGIN"); v != "" {
		CorsAllowOrigin = v
	}
	if v := os.Getenv("DB_HOSTNAME"); v != "" {
		DBHostName = v
	}
	if v, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64); err == nil {
		DBPort = int(v)
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DBName = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		DBUser = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		DBPassword = v
	}
	if v := os.Getenv("JWT_SECRET"); v != "" {
		JwtSecret = v
	}
	//環境変数を確認
	fmt.Println("HOSTNAME: ", HostName)
	fmt.Println("PORT: ", Port)
	fmt.Println("CORS_ALLOW_ORIGIN: ", CorsAllowOrigin)
	fmt.Println("DB_HOSTNAME: ", DBHostName)
	fmt.Println("DB_PORT: ", DBPort)
	fmt.Println("DB_NAME: ", DBName)
	fmt.Println("DB_USER: ", DBUser)
	fmt.Println("DB_PASSWORD: ", DBPassword)
	fmt.Println("JWT_SECRET: ", JwtSecret)
}
