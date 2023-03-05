package db

import (
	"database/sql"
	"fmt"
	"simple-video-server/config"
)

var _sqlDB *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	_sqlDB = db
}

func SqlExecLog[T any](f func(sql string, args ...any) (T, error), sql string, args ...any) (T, error) {
	fmt.Println("sql: ", sql, args)

	t, err := f(sql, args...)

	return t, err
}

func SqlSelectLog(f func(dest any, sql string, args ...any) error, dest any, sql string, args ...any) error {
	fmt.Println("sql: ", sql, args)

	err := f(dest, sql, args...)

	return err
}
