package _4_delete_association

import (
	"fmt"
	gormStudy "gorm-study"
	"log"
)

//## 删除关联
//
//删除与某模型间的关联关系：使用方法：.Delete() 完成
//
//- 多对一、一对多，删除关联字段
//- 多对多，删除关联记录
//- 对应的实体记录不会删除
// 若要删除全部关联：使用方法 Association.Clear() 完成

type Author gormStudy.Author
type Essay gormStudy.Essay
type Tag gormStudy.Tag
type EssayMate gormStudy.EssayMate

func AssocDelete() {
	DB := gormStudy.DB
	// B. 删除，外键的
	// 创建测试数据
	var a Author
	a.Name = "一位作者"
	if err := DB.Create(&a).Error; err != nil {
		log.Println(err)
	}
	log.Println("a:", a.ID)

	var e1, e2, e3 Essay
	e1.Subject = "一篇内容"
	e2.Subject = "另一篇内容"
	e3.Subject = "第三篇内容"
	if err := DB.Create([]*Essay{&e1, &e2, &e3}).Error; err != nil {
		log.Println(err)
	}
	log.Println("e1, e2, e3: ", e1.ID, e2.ID, e3.ID)

	// 添加关联
	if err := DB.Model(&a).Association("Essays").Replace([]Essay{e1, e2, e3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))

	if err := DB.Model(&a).Association("Essays").Delete([]Essay{e1, e3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))
	fmt.Println("------------------------")

	// B. 删除，多对多，关联表
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
	if err := DB.Model(&e1).Association("Tags").Append([]Tag{t1, t2, t3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(e1.Tags))

	if err := DB.Model(&e1).Association("Tags").Delete([]Tag{t1, t3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(e1.Tags))

	// C. 清空关联
	if err := DB.Model(&e1).Association("Tags").Clear(); err != nil {
		log.Println(err)
	}
	fmt.Println(len(e1.Tags))
}
