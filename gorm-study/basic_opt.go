package gorm_study

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

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
