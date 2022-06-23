package main

import (
	"Gin_Day1/config"
	"Gin_Day1/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 默认为 debug 模式，设置为发布模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化引擎
	engine := gin.Default()

	// 为之前初始化的引擎设置路由
	router.InitRouter(engine)

	// 启动引擎 传入端口
	engine.Run(config.PORT)
}

