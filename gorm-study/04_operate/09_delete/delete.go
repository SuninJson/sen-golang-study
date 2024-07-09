package _9_delete

import (
	"fmt"
	gorm_study "gorm-study"
	"log"
)

type Content gorm_study.Content

//### 主键删除
//
//将具有主键的模型作为参数传递给 `DB.Delete()` 方法，会删除该模型对应的记录。
//
//参考，基础操作删除部分。默认删除，是通过将DeleteAt字段设置为删除时间来实现的。若不存在DeleteAt字段，会执行Delete操作完成删除。
//
//示例：

func Delete() {
	DB := gorm_study.DB
	// 获取模型对象
	content := &Content{}
	if err := DB.First(content, 1).Error; err != nil {
		log.Fatal(err)
	}

	// DB.Delete() 删除
	if err := DB.Delete(content).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println("content was deleted")

	// 当然也可以
	DB.Delete(&Content{}, 1)
}

//### 条件删除
//
//通过Where子句，或Delete的内联条件，可以删除满足条件的记录。
//
//示例：

func DeleteWhere() {
	DB := gorm_study.DB
	if err := DB.Delete(&Content{}, "likes < ?", 100).Error; err != nil {
		log.Fatalln(err)
	}
	// [7.893ms] [rows:0] UPDATE `content` SET `deleted_at`='2023-04-21 18:57:13.338' WHERE likes < 100 AND `content`.`deleted_at` IS NULL
	if err := DB.Where("likes < ?", 100).Delete(&Content{}).Error; err != nil {
		log.Fatalln(err)
	}
}

//#### 查询被删除记录
//
//使用 `DB.Unscoped`能发来查询到被软删除的记录：

func FindDeleted() {
	DB := gorm_study.DB
	var c Content
	DB.Delete(&c, 13)

	if err := DB.First(&c, 13).Error; err != nil {
		log.Println(err)
	}
	//[4.604ms] [rows:0] SELECT * FROM `content` WHERE `content`.`id` = 13 AND `content`.`deleted_at` IS NULL ORDER BY `content`.`id` LIMIT 1

	if err := DB.Unscoped().First(&c, 13).Error; err != nil {
		log.Println(err)
	}
	// [3.320ms] [rows:1] SELECT * FROM `content` WHERE `content`.`id` = 13 ORDER BY `content`.`id` LIMIT 1
	fmt.Printf("%+v\n", c)
}

// #### 物理删除
//
//你可以使用 `DB.Unscoped`方法来永久删除匹配的记录

func DeleteHard() {
	DB := gorm_study.DB
	var c Content
	if err := DB.Unscoped().Delete(&c, 14).Error; err != nil {
		log.Fatalln(err)
	}
	//	[8.135ms] [rows:0] DELETE FROM `content` WHERE `content`.`id` = 14
}
