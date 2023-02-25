package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"simple-video-server/config"
	"simple-video-server/models"
	"time"
)

//var MysqlDB *gorm.DB

func SetupMysql() *gorm.DB {
	//logMode := logger.Info
	//if config.Env.Debug {
	//	logMode = logger.Error
	//}

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			//TablePrefix:   "",
		},
		//Logger: logger.Default.LogMode(logMode),
	})

	db.AutoMigrate(
		&models.User{},
		&models.Video{},
	)

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(10) // 最大空闲链接 TODO: 配置化
	sqlDb.SetMaxOpenConns(10) // 最大打开链接 TODO: 配置化
	sqlDb.SetConnMaxLifetime(time.Hour)

	if config.Env.Debug {
		//MysqlDB = db.Debug()
		return db.Debug()
	}

	//MysqlDB = db
	return db
}
