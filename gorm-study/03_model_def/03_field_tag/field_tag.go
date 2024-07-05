package _3_field_tag

import (
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"log"
)

type FieldTag struct {
	gorm.Model

	FTypeChar    string `gorm:"type:char(127)"`
	FTypeVarchar string `gorm:"type:varchar(255)"`
	FTypeText    string `gorm:"type:text"`
	FTypeBlob    []byte `gorm:"type:blob"`
	FTypeEnum    string `gorm:"type:enum('Go', 'GORM', 'MySQL')"`
	FTypeSet     string `gorm:"type:set('Go', 'GORM', 'MySQL')"`

	FColName     string `gorm:"column:custom_field"`
	FTypeNotNull string `gorm:"type:varchar(255);not null"`
	FTypeDefault string `gorm:"type:varchar(255);not null;default:hello gorm!"`
	FTypeComment string `gorm:"type:varchar(255);not null;default:hello gorm!;comment:some comment"`
}

func CreateFieldTagTable() {
	// 在MySQL中执行命令 show create table gorm_study_fieldtag; 来查看生成的表结构
	if err := gorm_study.DB.AutoMigrate(&FieldTag{}); err != nil {
		log.Fatalln(err)
	}
}
