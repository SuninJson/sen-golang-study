package module2

import "github.com/gin-gonic/gin"

func Router(group *gin.RouterGroup) {
	group.GET("/hello1", Hello1)
	group.GET("/hello2", Hello2)
}

func Hello2(context *gin.Context) {
	context.String(200, "module 2 hello2")
}

func Hello1(context *gin.Context) {
	context.String(200, "module 2 hello1")
}
