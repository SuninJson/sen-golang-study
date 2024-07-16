package _6_save_association

import (
	gormStudy "gorm-study"
	"log"
)

//## 自动存储关联
//
//在创建或更新模型时，如果关联模型存在，GORM会自动存储关联：
//
//示例：

func AssocSave() {
	DB := gormStudy.DB
	var t1 gormStudy.Tag
	DB.First(&t1, 10)

	e := gormStudy.Essay{
		Subject: "一个组合的Save",
		Author:  gormStudy.Author{Name: "Sen"},
		Tags: []gormStudy.Tag{
			t1,
			{Title: "自动存储关联"},
			{Title: "GORM"},
		},
	}

	if err := DB.Save(&e).Error; err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", e)

}
