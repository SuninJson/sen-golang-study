package study_serializer

import (
	"errors"
	"fmt"
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"log"
)

type TestSerializerStruct struct {
	gorm.Model
	Subject string
	// 若字符串数组类型不添加标签来序列化为JSON字符串，
	// 则进行数据库操作时会报错 [error] unsupported data type: &[]
	Tags []string `gorm:"serializer:json"`
}

func (TestSerializerStruct) TableName() string {
	return "test_serializer_struct"
}

func SerializerCurd() {
	if err := gorm_study.DB.AutoMigrate(&TestSerializerStruct{}); err != nil {
		log.Fatalln(err)
	}

	// 常规操作
	newTestSerializer := &TestSerializerStruct{}
	// DB.First方法会执行反序列化工作
	if result := gorm_study.DB.First(newTestSerializer, "id = ?", "1"); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到匹配的记录
			fmt.Println("未找到匹配的记录，进行创建")
		}
		testSerializer := &TestSerializerStruct{}
		testSerializer.Subject = "使用Serializer操作Tags字段"
		testSerializer.Tags = []string{"Go", "Serializer", "Gorm", "MySQL"}
		// create 会执行序列化工作,serialize
		if err := gorm_study.DB.Create(testSerializer).Error; err != nil {
			log.Fatal(err)
		}

		// 未找到匹配的记录，创建后，再次查询
		gorm_study.DB.First(newTestSerializer, "id = ?", "1")
	}

	fmt.Printf("%+v\n", newTestSerializer)

}
