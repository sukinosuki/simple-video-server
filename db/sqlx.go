package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"simple-video-server/config"
)

var _sqlxDB *sqlx.DB

func GetSqlxDB() *sqlx.DB {
	return _sqlxDB
}

func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database)

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	_sqlxDB = db
}
