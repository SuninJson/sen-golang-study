package _7_preload_association

import (
	gormStudy "gorm-study"
	"gorm.io/gorm/clause"
	"log"
)

type Author gormStudy.Author
type Essay gormStudy.Essay
type Tag gormStudy.Tag
type EssayMate gormStudy.EssayMate

//## 预加载
//
//预加载，指的是在查询当前模型时，自动查询关联的模型，使用方法：
//
//```go
//.Preload("Association", conds)
//```
//
//来实现。
//
//支持多次链式调用，来预加载多个关联。
//
//支持指定关联查询条件。
//
//示例：

func AssocPreload() {
	DB := gormStudy.DB
	// A.直接一步查询Author对应的Essays
	a := Author{}
	if err := DB.
		Preload("Essays").
		First(&a, 1).Error; err != nil {
		log.Fatalln(err)
	}
	// [3.840ms] [rows:2] SELECT * FROM `essay` WHERE `essay`.`author_id` = 1 AND `essay`.`deleted_at` IS NULL
	// [13.014ms] [rows:1] SELECT * FROM `author` WHERE `author`.`id` = 1 AND `author`.`deleted_at` IS NULL ORDER BY `author`.`id` LIMIT 1
	log.Println(a.Essays)
	log.Println("--------------------")

	// B.支持条件过滤
	if err := DB.
		Preload("Essays", "id IN ?", []uint{2, 3, 4}).
		First(&a, 1).Error; err != nil {
		log.Fatalln(err)
	}
	// [3.217ms] [rows:1] SELECT * FROM `essay` WHERE `essay`.`author_id` = 1 AND id IN (2,3,4) AND `essay`.`deleted_at` IS NULL
	log.Println(a.Essays)
	log.Println("-----------------------")

	// C. 支持多次链式调用，同时预加载多个关联
	e := Essay{}
	if err := DB.
		Preload("Author").
		Preload("EssayMate").
		Preload("Tags").
		First(&e, 1).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(e)
	// [2.776ms] [rows:1] SELECT * FROM `author` WHERE `author`.`id` = 1 AND `author`.`deleted_at` IS NULL
	// [10.398ms] [rows:0] SELECT * FROM `essay_mate` WHERE `essay_mate`.`essay_id` = 1 AND `essay_mate`.`deleted_at` IS NULL
	// [3.260ms] [rows:2] SELECT * FROM `essay_tag` WHERE `essay_tag`.`essay_id` = 1
	// [3.264ms] [rows:2] SELECT * FROM `tag` WHERE `tag`.`id` IN (1,3) AND `tag`.`deleted_at` IS NULL
	// [28.067ms] [rows:1] SELECT * FROM `essay` WHERE `essay`.`id` = 1 AND `essay`.`deleted_at` IS NULL ORDER BY `essay`.`id` LIMIT 1

	// 预加载全部，`clause.Associations`不会预加载层级的关联，可以配合多级预加载一起使用
	e2 := Essay{}
	if err := DB.
		Preload(clause.Associations).
		First(&e2, 1).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(e)
}

//### 多级预加载
//
//.Preload()的参数支持层级语法：
//
//```go
//.Preload("Association1.Association2.Assocaition3", conds)
//```
//
//默认情况下，GORM仅仅会加载一级的关联。使用多级语法，可以预加载多级关联数据。
//
//示例：

func AssocLevelPreload() {
	DB := gormStudy.DB
	a := Author{}
	if err := DB.
		//Preload("Essays").
		// 多级关联
		Preload("Essays.Tags").
		First(&a, 1).Error; err != nil {
		log.Fatalln(err)
	}
	// [3.843ms] [rows:5] SELECT * FROM `essay_tag` WHERE `essay_tag`.`essay_id` IN (1,2)
	// [3.284ms] [rows:3] SELECT * FROM `tag` WHERE `tag`.`id` IN (1,3,2) AND `tag`.`deleted_at` IS NULL
	// [10.396ms] [rows:2] SELECT * FROM `essay` WHERE `essay`.`author_id` = 1 AND `essay`.`deleted_at` IS NULL
	// [17.609ms] [rows:1] SELECT * FROM `author` WHERE `author`.`id` = 1 AND `author`.`deleted_at` IS NULL ORDER BY `author`.`id` LIMIT 1

	log.Println(a.Essays[0].Tags)
	log.Println(a.Essays[1].Tags)
}
