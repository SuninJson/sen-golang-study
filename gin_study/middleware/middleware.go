package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloMiddleware(ctx *gin.Context) {
	fmt.Println("只对 /hello 路由生效，hello中间件-开始")
	fmt.Println("只对 /hello 路由生效，hello中间件-结束")
}

func CustomMiddleware01(ctx *gin.Context) {
	fmt.Println("自定义中间件，处理一些统一逻辑，如鉴权，日志，限流等")
	fmt.Println("中间件1-开始")
	// gin.Context.Next 函数作用：走中间件链中的下一个中间件
	ctx.Next()
	fmt.Println("中间件1-结束")
}

func CustomMiddleware02(ctx *gin.Context) {
	fmt.Println("中间件2-开始")
	// 可以使用 gin.Context.Abort函数，来处理符合某条件后，主动终止中间件链的逻辑
	//if conditionIsTrue() {
	//	ctx.Abort()
	//}
	fmt.Println("中间件2-结束")
}

func conditionIsTrue() bool {
	return true
}

func CustomMiddleware03(ctx *gin.Context) {
	fmt.Println("中间件3-开始")
	fmt.Println("中间件3-结束")
}
