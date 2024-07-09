package gorm_study

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Username string
	Name     string
	Email    string
	Birthday *time.Time
}

type Content struct {
	gorm.Model

	Subject string
	// 通过default标签来设置默认值，当字段为类型零值时，触发使用默认值
	Likes       uint  `gorm:"default 100"`
	LikesPoint  *uint `gorm:"default 99"`
	Views       uint
	PublishTime *time.Time
	AuthorID    uint `gorm:"default 1"`
}

type Author struct {
	gorm.Model
	Status int

	Name  string
	Email string
}
