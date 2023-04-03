package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os/signal"
	"simple-video-server/app_server"
	"simple-video-server/config"
	_ "simple-video-server/docs"
	"syscall"
	"time"
)

var AppRouter *gin.RouterGroup

var AdminRouter *gin.RouterGroup

func Run(router *gin.Engine) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	port := fmt.Sprintf(":%d", config.Env.Port)

	fmt.Println("run000 port ", port)

	go func() {
		err := router.Run(port)
		fmt.Println("run1111 err ", err)

		if err != nil {
			panic(err)
		}

	}()

	fmt.Println("before ctx done")
	<-ctx.Done()
	fmt.Println("ctx done!!")

	fmt.Println("准备关闭应用")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancelFunc()

	<-ctx.Done()

	fmt.Println("关闭应用")

}

func SetupRouter() {

	var router *gin.Engine

	if !config.Env.Debug {
		gin.SetMode(gin.ReleaseMode)
		//router = gin.Default()
		router = gin.New()
		router.Use(gin.Recovery())
	} else {
		//gin.SetMode(gin.ReleaseMode)
		router = gin.Default()
	}

	SetupSwaggerRoute(router)

	AppRouter = router.Group("/app")

	AdminRouter = router.Group("/admin")

	//
	app_server.SetupRoutes(AppRouter)

	//

	Run(router)
}

func SetupSwaggerRoute(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
