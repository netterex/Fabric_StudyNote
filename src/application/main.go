package main

import (
	"fabric-houserental/application/allrouter"
	"fabric-houserental/application/middle_ware"
	"github.com/gin-gonic/gin"
)

/*
	Gin 框架入口
*/

func main() {
	rou := gin.Default()

	rou.Use(middle_ware.CorsMiddleWare)
	allrouter.InitRouter(rou)

	rou.Run()
}
