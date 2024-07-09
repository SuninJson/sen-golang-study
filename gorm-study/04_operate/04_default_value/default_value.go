package default_value

import (
	"fmt"
	gorm_study "gorm-study"
)

func DefaultValue() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{})

	c := gorm_study.Content{}
	c.Subject = "原始内容"
	likes, likesPoint := uint(0), uint(0)
	c.Likes = likes
	// 通过指针类型可以在字段设置了默认值时使用0值
	c.LikesPoint = &likesPoint
	DB.Create(&c)
	fmt.Println(c.Likes, *c.LikesPoint)
}

// 实操中通常使用模型的创建方法来初始化默认值，不通过定义default标签的方案

const (
	defaultViews = 99
	defaultLikes = 99
)

func NewContent() gorm_study.Content {
	return gorm_study.Content{
		Likes: defaultLikes,
		Views: defaultViews,
	}
}

func DefaultValueOften() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{})

	c := NewContent()
	c.Subject = "原始内容"
	DB.Create(&c)
	fmt.Println(c.Likes, c.Views)
}
