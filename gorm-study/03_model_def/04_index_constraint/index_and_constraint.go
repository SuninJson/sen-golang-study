package _4_index_constraint

import (
	gorm_study "gorm-study"
	"log"
)

type IAndC struct {
	ID          uint    `gorm:"primaryKey"`
	Email       string  `gorm:"type:varchar(255);unique"`
	Age         int8    `gorm:"index;check:age >= 18 AND email is not null"`
	FirstName   string  `gorm:"index:name"`
	LastName    string  `gorm:"index:name"`
	FirstName1  string  `gorm:"index:name1,priority:2"`
	LastName1   string  `gorm:"index:name1,priority:1"`
	Height      float32 `gorm:"index:,sort:desc"`
	AddressHash string  `gorm:"type:varchar(42);index:,length:12"`
	Telephone   string  `gorm:"type:varchar(16);uniqueIndex:,comment:电话号码唯一索引"`
}

func CreateIAndCTable() {
	// 在MySQL中执行命令 show create table gorm_study_iandc; 来查看生成的表结构
	if err := gorm_study.DB.AutoMigrate(&IAndC{}); err != nil {
		log.Fatalln(err)
	}
}
