package _5_find_association

import (
	gormStudy "gorm-study"
	"log"
)

type Author gormStudy.Author
type Essay gormStudy.Essay
type Tag gormStudy.Tag
type EssayMate gormStudy.EssayMate

//## 关联查询
//
//使用Find方法，可以查找关联。查找的结果通常是关联的模型或模型切片，支持子句过滤，例如条件，排序，Limit等：

func AssocFind() {
	DB := gormStudy.DB
	e := Essay{}
	DB.First(&e, 18)

	// 查询关联的tags
	//var ts []Tag
	if err := DB.Model(&e).Association("Tags").Find(&e.Tags); err != nil {
		log.Println(err)
	}
	log.Println(e.Tags)

	// 子句，要写在Association()方法前面
	if err := DB.Model(&e).
		Where("tag_id > ?", 7).
		Order("tag_id DESC").
		Association("Tags").Find(&e.Tags); err != nil {
		log.Println(err)
	}
	log.Println(e.Tags)

	// 查询关联的模型的数量
	// .Count()方法可以返回关联的数量，不用查询到全部的关联实体。
	count := DB.Model(&e).Association("Tags").Count()
	log.Println("count:", count)

}
