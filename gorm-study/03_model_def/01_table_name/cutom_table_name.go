package _1_table_name

import (
	"gorm-study"
	"gorm.io/gorm"
	"log"
)

type Box struct {
	gorm.Model
}

// TableName 自定义表名示例
func (Box) TableName() string {
	return "custom_box"
}

func CreateBoxTable() {
	if err := gorm_study.DB.Debug().AutoMigrate(&Box{}); err != nil {
		log.Fatal(err)
	}

	//临时指定表名
	if err := gorm_study.DB.Debug().Table("temp_box").AutoMigrate(&Box{}); err != nil {
		log.Fatal(err)
	}
}
