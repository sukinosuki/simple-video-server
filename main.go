package main

import (
	"simple-video-server/db"
	"simple-video-server/pkg/global"
)

// @title simple video server222
// @version 0.0.1
// @description 描述111
func main() {

	//global.MysqlDB = db.SetupMysql()
	global.MysqlDB = db.GetOrmDB()

	global.RDB = db.GetRedisDB()

	SetupRouter()
}
