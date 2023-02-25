package video

type VideoAdd struct {
	Title string `json:"title" form:"title" binding:"required,max=50"`
	//Age   int    `json:"age" form:"age"`
	//File  any    `json:"file" form:"file"`
}

type VideoUpdate struct {
	Title string `json:"title" form:"title"`
}

//func init() {
//	err := global.MysqlDB.AutoMigrate(
//		&Video{},
//	)
//
//	if err != nil {
//		panic(err)
//	}
//}
