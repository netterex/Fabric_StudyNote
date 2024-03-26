package police

import "github.com/gin-gonic/gin"

func Router(router *gin.RouterGroup) {

	router.GET("/query", Query)

	router.POST("/set", Set)
}
