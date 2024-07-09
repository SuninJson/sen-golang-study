package update_insert

import (
	"fmt"
	gorm_study "gorm-study"
	"gorm.io/gorm/clause"
	"log"
)

// UpSert, Update Insert的缩写。逻辑是当插入冲突时（主键或唯一键已经存在），执行更新操作。判定是否冲突的要素通常是主键。
// GROM通过 clause.OnConflict{} 类型实现控制冲突后的行为
// 使用DB.Clauses()子句，传入以上类型实例来配置Create()冲突后的操作

func UpSert() {
	DB := gorm_study.DB
	c := gorm_study.Content{}
	c.Subject = "原始标题"
	c.Likes = 12
	DB.Create(&c)
	fmt.Println(c)

	c2 := gorm_study.Content{}
	c2.ID = c.ID
	c2.Subject = "新标题"
	c2.Likes = 20
	if err := DB.Create(&c2).Error; err != nil {
		log.Println(err)
		//Error 1062 (23000): Duplicate entry '13' for key 'msb_content.PRIMARY'
	}

	c3 := gorm_study.Content{}
	c3.ID = c.ID
	c3.Subject = "新标题UpdateAll"
	c3.Likes = 20
	if err := DB.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&c3).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("UpdateAll:", c3)

	c4 := gorm_study.Content{}
	c4.ID = c.ID
	c4.Subject = "新标题DoUpdates"
	c4.Likes = 20
	if err := DB.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"likes"})}).Create(&c4).Error; err != nil {
		log.Fatal(err)
	}
	log.Println("clause.OnConflict{DoUpdates}", c4)
}
