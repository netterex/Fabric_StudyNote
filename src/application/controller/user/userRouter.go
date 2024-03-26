package user

import "github.com/gin-gonic/gin"

func Router(router *gin.RouterGroup) {

	router.POST("/login", Login)

	router.POST("/register", Register)
}
