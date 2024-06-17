package router

import (
	"gin_study/module1"
	"gin_study/module2"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	module1Group := engine.Group("/module1")
	module2Group := engine.Group("/module2")

	module1.Router(module1Group)
	module2.Router(module2Group)
}
