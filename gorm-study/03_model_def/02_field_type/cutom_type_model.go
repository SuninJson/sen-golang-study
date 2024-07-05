package _2_field_type

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm-study"
	"gorm.io/gorm"
	"time"
)

// 自定义字段类型

type CustomTypeModel struct {
	gorm.Model

	FTime     time.Time
	FNullTime sql.NullTime

	FString     string
	FNullString sql.NullString

	FUUID     uuid.UUID
	FNullUUID uuid.NullUUID
}

func CustomType() {
	// 如何实现自定义字段类型，可参考uuid.UUID的实现
	//id := uuid.UUID{}
	//id.Scan()  // 实现Scanner接口
	//id.Value() // 实现Valuer接口

	// 初始化模型
	ctm := &CustomTypeModel{}
	// 迁移数据表
	gorm_study.DB.AutoMigrate(ctm)

	// 创建
	ctm.FTime = time.Now()         // 当前时间
	ctm.FNullTime = sql.NullTime{} // 零值，Valid默认为false
	ctm.FString = ""
	ctm.FNullString = sql.NullString{}

	ctm.FUUID = uuid.New()
	ctm.FNullUUID = uuid.NullUUID{}

	gorm_study.DB.Create(ctm)

	// 查询
	gorm_study.DB.First(ctm, ctm.ID)

	// 判定字段是否为NULL
	if ctm.FString == "" {
		fmt.Println("FString is NULL")
	} else {
		fmt.Println("FString is NOT NULL")
	}

	if ctm.FNullString.Valid == false {
		fmt.Println("FNullString is NULL")
	} else {
		fmt.Println("FNullString is NOT NULL")
	}
}
