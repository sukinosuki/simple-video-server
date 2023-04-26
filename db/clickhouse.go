package db

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	click "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"simple-video-server/config"
	"simple-video-server/models"
	"time"
)

var clickhouseDB *gorm.DB

func GetClickhouseDB() *gorm.DB {
	return clickhouseDB
}

func init() {
	addr := fmt.Sprintf("%s:%d", config.Clickhouse.Host, config.Clickhouse.Port)
	fmt.Println("clickhouse addr ", addr)

	sqlDB := clickhouse.OpenDB(&clickhouse.Options{
		//Addr: []string{"192.168.10.100:9000"},
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: "simple_video_server",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		//Compression: &clickhouse.Compression{
		//	clickhouse.CompressionLZ4,
		//},
		Debug: config.Env.Debug,
	})

	db, err := gorm.Open(click.New(click.Config{
		Conn: sqlDB,
	}))

	if err != nil {
		fmt.Println("open clickhouse err ", err)
		return
		//panic(err)
	}

	err = db.AutoMigrate(
		&models.RequestLog{},
		&models.OperationLog{},
	)
	if err != nil {
		panic(err)
	}
	clickhouseDB = db
}
