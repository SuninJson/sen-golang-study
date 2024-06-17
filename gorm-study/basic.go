package gorm_study

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func BasicUse() {
	// 定义DSN（Data Source Name）
	// [MySQL DSN 说明](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"

	// 通过连接池打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connecting DB success!")

	// 通过模型来创建表结构
	if err := db.AutoMigrate(&Article{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Migrate article table success!")

}

const (
	GlobalLogger = 1
	FileLogger   = 2
)

func gormConfig(initType int) *gorm.Config {
	var loggerInterface logger.Interface
	switch initType {
	case GlobalLogger:
		// 可以在 `gorm.Open` 时设置全局日志级别，
		// 或者通过 `db.Config.Logger`完成配置：DB.Config.Logger = logger.Default.LogMode(logger.Info)
		loggerInterface = logger.Default.LogMode(logger.Info)
	case FileLogger:
		// 0644的意思：
		// （1）开头的0是一个约定，用来明确后面跟随的是一个八进制数。它本身并不直接代表任何文件权限，而是作为一个指示符，告诉解析这个数值的系统（或人）后面的数字应当被解释为八进制形式。
		// （2）第2位6：4表示读权限（r）2表示写权限（w）因此，6是4+2，意味着文件所有者有读和写权限。
		// （3）第3位4：表示所属组用户只有读权限。
		// （4）第4位4：表示其他用户的只有读权限
		logWriter, _ := os.OpenFile("./sql.log", os.O_CREATE|os.O_APPEND, 0644)
		// LstdFlags是以下两个常用标志的组合：
		// log.Ldate：这个标志表示在每次日志输出时包含日期信息
		// log.Ltime：这个标志表示在每次日志输出时包含时间信息
		loggerInterface = logger.New(log.New(logWriter, "\n", log.LstdFlags),
			logger.Config{
				// 慢查询阈值 200ms
				SlowThreshold: 200 * time.Millisecond,
				// 日志级别
				LogLevel: logger.Info,
				// 是否忽略记录不存在的错误
				IgnoreRecordNotFoundError: false,
				// 不彩色化
				Colorful: false,
			})
	default:
		loggerInterface = logger.Default
	}

	return &gorm.Config{
		Logger: loggerInterface,
	}
}

func init() {
	// 定义DSN
	const dsn = "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接服务器（池）
	db, err := gorm.Open(mysql.Open(dsn), gormConfig(FileLogger))

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func Create() {
	// 构建Article类型数据
	article := &Article{
		Subject:     "GORM 的 CRUD 基础操作",
		Likes:       0,
		Published:   true,
		PublishTime: time.Now(),
		AuthorID:    42,
	}

	// DB.Create 完成数据库的insert
	if err := DB.Create(article).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}

func Retrieve(id uint) {
	// 初始化Article模型，零值
	article := &Article{}

	// 选择查询单个还是多个
	//- Find() 多个
	//- First() 单个
	if err := DB.First(article, id).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}

func Update() {
	// 获取需要更新的对象
	article := &Article{}
	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}

	// 更新对象字段
	article.AuthorID = 23
	article.Likes = 101
	article.Subject = "新的文章标题"

	// 存储，DB.Save()
	if err := DB.Save(article).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}

func Delete() {
	// DB.Delete() 删除
	if err := DB.Debug().Delete(&Article{}, 1).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println("article was deleted")
}

// Article 定义模型，以文章的数据结构为例
type Article struct {
	// 通常情况下，要嵌入 gorm.Model，用于保有核心字段
	// gorm.Model.DeletedAt 用来记录删除时间，可用于逻辑删除的功能
	gorm.Model

	Subject     string
	Likes       uint
	Published   bool
	PublishTime time.Time
	AuthorID    uint
}
