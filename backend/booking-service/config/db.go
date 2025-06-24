package config

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
  var err error
  dsn := "root:@tcp(127.0.0.1:3306)/user_service?charset=utf8mb4&parseTime=True&loc=Local"
  DB, err = sql.Open("mysql", dsn)
  if err != nil {
    log.Fatal("DB connection failed:", err)
  }

  if err = DB.Ping(); err != nil {
    log.Fatal("DB ping error:", err)
  }
  log.Println("Database connected!")
}
