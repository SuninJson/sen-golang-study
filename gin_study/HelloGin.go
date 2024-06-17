package main

import (
	"fmt"
	"gin_study/middleware"
	"gin_study/router"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default()返回的是一个引擎Engine，Engine是Gin框架非常重要的数据结构，是框架的入口
	engine := gin.Default()
	// 注册自定义模板函数
	engine.SetFuncMap(template.FuncMap{"add": Add})

	// 使用中间件
	engine.Use(middleware.CustomMiddleware01, middleware.CustomMiddleware02, middleware.CustomMiddleware03)
	// 通过engine使用中间件，会在全局的路由中生效
	// 若我们希望中间件只在某个路由组生效可通过使用RouterGroup.Use函数
	// 若我们希望中间件只在具体的某一个路由生效，我们可以在具体路由和函数之间的参数加上中间件
	engine.GET("/hello", middleware.HelloMiddleware, helloGin)

	// 通过HTML模板文件渲染
	engine.LoadHTMLGlob("templates/*")
	// 使用静态文件
	engine.Static("/static", "static")
	engine.GET("/helloHtml", helloLoadHtml)

	// 将后端的值传给前端
	// 渲染字符串
	engine.GET("/helloName", helloName)
	// 渲染结构体
	engine.GET("/helloStruct", helloStruct)
	// 渲染数组
	engine.GET("/helloArray", helloArray)
	// 渲染Map中的结构体
	engine.GET("/helloMapStruct", helloMapStruct)

	// 将前端的值传递给后端
	// 获取路径参数的值
	// 使用‘:’作为占位符，必须在路径给值才可以匹配到路径，否则在浏览器会报404
	// 使用‘*’作为占位符，是否给路径给值都会匹配的路径
	engine.GET("/helloPathValue1/:id", helloPathValue)
	engine.GET("/helloPathValue2/*id", helloPathValue)
	// 通过键值对的形式传递参数
	engine.GET("/helloQueryValue", helloQueryValue)

	// 获取POST请求数据
	engine.GET("/helloPost", helloPost)
	engine.POST("/getUserInfo", getUserInfo)

	// 使用Ajax
	engine.GET("/helloAjax", helloAjax)
	engine.POST("/validateUserName", validateUserName)

	// 文件上传
	engine.GET("/helloUpload", helloUpload)
	engine.POST("/upload", upload)

	// 上传多个文件
	engine.GET("/helloUploadFiles", helloUploadFiles)
	engine.POST("/uploadFiles", uploadFiles)

	// 响应重定向
	engine.GET("/helloRedirect1", helloRedirect1)
	engine.GET("/helloRedirect2", helloRedirect2)

	// 模板函数
	engine.GET("/helloTemplateFunc", helloTemplateFunc)

	// 路由组：将不同的路由按照版本、模块进行不同的分组，利于维护，方便管理。
	v1 := engine.Group("/v1")
	{
		// 表单数据绑定
		v1.GET("/helloBindForm", helloBindForm)
		v1.POST("/doBindForm", doBindForm)
	}

	v2 := engine.Group("/v2")
	{
		// JSON数据绑定
		v2.GET("/helloBindFormJSON", helloBindFormJSON)
		v2.POST("/doBindFormJSON", doBindFormJSON)
	}

	// 按业务模块路由分组
	router.Router(engine)

	err := engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

type UserJSON struct {
	Uname string `json:"uname"`
	Age   int    `json:"age"`
}

func doBindFormJSON(context *gin.Context) {
	var user UserJSON
	err := context.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		context.JSON(404, gin.H{"msg": "绑定失败"})
	} else {
		context.JSON(200, gin.H{"msg": "绑定成功"})
	}
}

func helloBindFormJSON(context *gin.Context) {
	context.HTML(200, "helloBindFormJSON.html", nil)
}

func doBindForm(context *gin.Context) {
	var user User //数据绑定：
	err := context.ShouldBind(&user)
	//打印结构体对象的内容：
	fmt.Println(user)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}

type User struct {
	Username string `form:"username"`
	Password string `form:"pwd"`
}

func helloBindForm(context *gin.Context) {
	context.HTML(200, "helloBindForm.html", nil)
}

func Add(num1 int, num2 int) int { return num1 + num2 }

func helloTemplateFunc(context *gin.Context) {
	dataMap := make(map[string]interface{})

	arr := []string{"1", "2", "3"}
	dataMap["arr"] = arr
	dataMap["nowTime"] = time.Now()
	context.HTML(200, "helloTemplateFunc.html", dataMap)
}

func helloRedirect2(context *gin.Context) {
	context.String(http.StatusOK, "重定向成功，helloRedirect2")
}

func helloRedirect1(context *gin.Context) {
	fmt.Println("helloRedirect1")
	context.Redirect(http.StatusFound, "/helloRedirect2")
}

func uploadFiles(context *gin.Context) {
	// 获取form表单
	form, _ := context.MultipartForm()
	//在form表单中获取name相同的文件
	files := form.File["myfile"]
	// 遍历处理
	for _, file := range files {
		timeInt := time.Now().Unix()
		timeStr := strconv.FormatInt(timeInt, 10) //10:十进制
		err := context.SaveUploadedFile(file, "d://temp/"+timeStr+file.Filename)
		if err != nil {
			log.Println(err)
		}
	}
	//todo 响应上传结果
}

func helloUploadFiles(context *gin.Context) {
	context.HTML(200, "helloUploadFiles.html", nil)
}

func upload(context *gin.Context) {
	file, _ := context.FormFile("uploadedFile")
	fmt.Println(file.Filename)
	timeInt := time.Now().Unix()
	timeStr := strconv.FormatInt(timeInt, 10) //10:十进制
	err := context.SaveUploadedFile(file, "d://"+timeStr+file.Filename)
	if err != nil {
		log.Fatalln(err)
	}
	context.String(200, "文件上传成功")
}

func helloUpload(context *gin.Context) {
	context.HTML(200, "helloUpload.html", nil)
}

func validateUserName(context *gin.Context) {
	uname := context.PostForm("uname")
	fmt.Println(uname)
	fmt.Println(uname == "Sen")
	if uname == "Sen" {
		context.JSON(200, gin.H{"msg": "用户名重复了！"})
	} else {
		context.JSON(200, gin.H{"msg": ""})
	}
}

func helloAjax(context *gin.Context) {
	context.HTML(200, "helloAjax.html", nil)
}

func getUserInfo(context *gin.Context) {
	//获取post请求的参数：
	//PostForm方法：作用：通过key得到value数据
	uname := context.PostForm("username")
	pwd := context.PostForm("pwd")
	//DefaultPostForm方法:
	//作用：当页面中未定义表单元素进行提交给出默认值，
	//如果页面定义了元素但是提交没有提交数据，那么不会有默认值，会认为是没有提交数据
	age := context.DefaultPostForm("age", "18")
	//PostFormArray方法：
	//作用：如果前端value数据过多可以用数组接收：
	loveLanguage := context.PostFormArray("loveLanguage")
	//PostFormMap方法:
	//作用：获取map的数据,参数需要注意：传入的是整个map（而不是具体的key）
	userMap := context.PostFormMap("user")
	fmt.Println(uname)
	fmt.Println(pwd)
	fmt.Println(age)
	fmt.Println(loveLanguage)
	fmt.Println(userMap)
}

func helloPost(context *gin.Context) {
	context.HTML(200, "helloPost.html", nil)
}

func helloQueryValue(context *gin.Context) {
	id := context.Query("id")
	name := context.Query("name")
	context.String(200, "获取键值对参数的值，%s:%s", id, name)
}

func helloPathValue(context *gin.Context) {
	id := context.Param("id")
	context.String(200, "获取路径参数的值,%s", id)
}

func helloMapStruct(context *gin.Context) {
	structMap := map[string]StructDemo{
		"No1": {
			Id:   1,
			Name: "Sen1",
		},
		"No2": {
			Id:   2,
			Name: "Sen2",
		},
	}
	context.HTML(200, "helloMapStruct.html", structMap)
}

func helloArray(context *gin.Context) {
	arr := []string{"1", "2", "3"}
	context.HTML(200, "helloArray.html", arr)
}

func helloStruct(context *gin.Context) {
	structDemo := StructDemo{
		Id:   1,
		Name: "Sen",
	}
	context.HTML(200, "helloStruct.html", structDemo)
}

func helloName(context *gin.Context) {
	context.HTML(200, "hello.html", "My name")
}

func helloLoadHtml(context *gin.Context) {
	context.HTML(200, "hello.html", nil)
}

func helloGin(ctx *gin.Context) { // Gin把请求和响应都封装到gin.Context上下文环境中了
	ctx.String(200, "Hello Gin")
}

type StructDemo struct {
	Id   int8
	Name string
}
