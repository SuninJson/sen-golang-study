package _3_replace_association

import (
	"fmt"
	gormStudy "gorm-study"
	"log"
)

type Author gormStudy.Author
type Essay gormStudy.Essay
type Tag gormStudy.Tag
type EssayMate gormStudy.EssayMate

//## 替换关联
//
//使用新的关联关系，替换旧的关系。使用方法：.Replace() 完成
//
//主要用在对多的关系上。
//
//示例：

func AssocReplace() {
	DB := gormStudy.DB
	// A. 替换
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
	if err := DB.Model(&a).Association("Essays").Replace([]Essay{e1, e3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))
	// 基于当前的基础上，添加关联
	if err := DB.Model(&a).Association("Essays").Replace([]Essay{e2, e3}); err != nil {
		log.Println(err)
	}
	fmt.Println(len(a.Essays))

}
