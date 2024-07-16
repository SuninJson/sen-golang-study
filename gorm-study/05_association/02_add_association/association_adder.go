package _2_add_association

import (
	"fmt"
	gorm_study "gorm-study"
	"log"
)

type Author gorm_study.Author
type Essay gorm_study.Essay
type Tag gorm_study.Tag
type EssayMate gorm_study.EssayMate

//## 添加关联
//
//.Append() 方法添加关联。
//
//参数为需要关联的模型，或模型切片。取决于是对一还是对多。
//
//其中：
//
//- `many to many`、`has many` 添加新的关联
//- `has one`, `belongs to` 替换当前的关联
//
//示例：

func AssocAppend() {
	DB := gorm_study.DB
	// A：一对多的关系, Author 1:n Essay
	// 创建测试数据
	var a Author
	a.Name = "一位作者"
	if err := DB.Create(&a).Error; err != nil {
		log.Println(err)
	}
	log.Println("a:", a.ID)
	var e1, e2 Essay
	e1.Subject = "一篇内容"
	//e1.AuthorID = a.ID
	e2.Subject = "另一篇内容"
	if err := DB.Create([]*Essay{&e1, &e2}).Error; err != nil {
		log.Println(err)
	}
	log.Println("e1, e2: ", e1.ID, e2.ID)

	// 添加关联
	if err := DB.Model(&a).Association("Essays").Append([]Essay{e1}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))
	// 基于当前的基础上，添加关联
	if err := DB.Model(&a).Association("Essays").Append([]Essay{e2}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))
	// 添加后，a模型对象的Essays字段，自动包含了关联的Essay模型
	//fmt.Println(a.Essays)

	// B: Essay M:N TAg
	var t1, t2, t3 Tag
	t1.Title = "Go"
	t2.Title = "GORM"
	t3.Title = "Ma"
	if err := DB.Create([]*Tag{&t1, &t2, &t3}).Error; err != nil {
		log.Println(err)
	}
	log.Println("t1, t2, t3: ", t1.ID, t2.ID, t3.ID)

	// e1 t1, t3
	// e2 t1, t2, t3
	if err := DB.Model(&e1).Association("Tags").Append([]Tag{t1, t3}); err != nil {
		log.Println(err)
	}

	if err := DB.Model(&e2).Association("Tags").Append([]Tag{t1, t2, t3}); err != nil {
		log.Println(err)
	}

	// 关联表查看
	// mysql> select * from essay_tag;
	//+--------+----------+
	//| tag_id | essay_id |
	//+--------+----------+
	//|      1 |       12 |
	//|      3 |       12 |
	//|      1 |       13 |
	//|      2 |       13 |
	//|      3 |       13 |
	//+--------+----------+

	// C, Belongs To. Essay N:1 Author
	var e3 Essay
	e3.Subject = "第三篇内容"
	if err := DB.Create([]*Essay{&e3}).Error; err != nil {
		log.Println(err)
	}
	log.Println("e3: ", e3.ID)

	log.Println(e3.Author)
	// 关联
	if err := DB.Model(&e3).Association("Author").Append(&a); err != nil {
		log.Println(err)
	}
	log.Println(e3.Author.ID)

	// 对一的关联，会导致关联被更新
	var a2 Author
	a2.Name = "另一位作者"
	if err := DB.Create(&a2).Error; err != nil {
		log.Println(err)
	}
	log.Println("a2:", a2.ID)
	if err := DB.Model(&e3).Association("Author").Append(&a2); err != nil {
		log.Println(err)
	}
	log.Println(e3.Author.ID)

}
