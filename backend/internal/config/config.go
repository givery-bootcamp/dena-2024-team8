package config

import (
	"os"
	"strconv"
)

var HostName = "127.0.0.1"
var Port = 9000
var CorsAllowOrigin = "http://localhost:3000"
var DBHostName = "db"
var DBPort = 3306
var DBName = "training"

//ローカル環境か判定する関数
func IsLocal() bool {
	return os.Getenv("ENV") == "local"
}
//コミットするためだけのコメント

// ローカルの場合と本番の場合でDBのユーザー名とパスワードを変更する
var isLocal = IsLocal()
var DBUser = "admin"
var DBPassword = "dena-training-2024-team-8"

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
	if isLocal {
		DBUser = "root"
		DBPassword = ""
	}
	// if v := os.Getenv("DB_USER"); v != "" {
	// 	DBUser = v
	// }
	// if v := os.Getenv("DB_PASSWORD"); v != "" {
	// 	DBPassword = v
	// }
}
