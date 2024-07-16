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

//- Author 和 Author间，在Author的角度是一对多，在Author的角度是多对一
//- Author和Tag间，是多对多
//- Author和AuthorMate间，既可以是一对多，也可以做一对一，看业务逻辑，本例中我们采用一对一
//
//在GORM中，可以在模型中定义关联的方式，实现以上的对应的关系：
//
//- 使用模型类型，表示对应一个的关系
//- 使用模型切片类型，表示对应多个的关系
//- 使用tag，many2many表示多对多关系，需要制定关联表名
//- 需要使用外键字段确保关联。默认的关联字段是模型+ID的形式。
//  - 例如Author一对多关联Author，那么Author中就应该有AuthorID作为关联字段
//  - 允许自定义

// Author Author模型
type Author struct {
	gorm.Model
	Status int
	Name   string
	Email  string

	// 拥有多个论文内容
	Essays []Essay
}

// Essay 论文内容
type Essay struct {
	gorm.Model
	Subject string
	Content string

	// 外键字段
	AuthorID *uint

	// 属于某个作者
	Author Author

	// 拥有一个论文元信息
	EssayMate EssayMate

	// 拥有多个Tag
	Tags []Tag `gorm:"many2many:essay_tag"`
}

// EssayMate 论文元信息
type EssayMate struct {
	gorm.Model
	Keyword     string
	Description string

	// 外键字段
	EssayID *uint

	// 属于一个论文内容，比较少用
	//Essay *Essay
}

type Tag struct {
	gorm.Model
	Title string

	// 拥有多个Essay
	Essays []Essay `gorm:"many2many:essay_tag"`
}
