package create

import (
	"fmt"
	gorm_study "gorm-study"
	"log"
	"time"
)

func BasicCreate() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{})

	c1 := gorm_study.Content{}
	c1.Subject = "GORM的使用"

	result1 := DB.Create(&c1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println(c1.ID, result1.RowsAffected)
}

func UseMapCreate() {
	DB := gorm_study.DB
	// map 指定数据
	//设置map 的values
	values := map[string]any{
		"Subject":     "Map指定值",
		"PublishTime": time.Now(),
	}
	// create
	result2 := DB.Model(&gorm_study.Content{}).Create(values)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	// 测试输出
	fmt.Println(result2.RowsAffected)
}

func MultiCreate() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{})

	// model
	cs := []gorm_study.Content{
		{Subject: "批量插入标题1"},
		{Subject: "批量插入标题2"},
		{Subject: "批量插入标题3"},
	}
	result1 := DB.Create(&cs)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println(result1.RowsAffected)
	for _, c := range cs {
		fmt.Println(c.ID)
	}

	// map
	vs := []map[string]any{
		{"Subject": "批量插入标题4"},
		{"Subject": "批量插入标题5"},
		{"Subject": "批量插入标题6"},
	}
	result2 := DB.Model(&gorm_study.Content{}).Create(vs)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println(result2.RowsAffected)

	// 当需要执行大量的记录插入时，推荐使用分批插入。
	// 也就是不是一次性全部插入，而是每次插入固定的条数，直到全部插入完毕。
	// 分批插入的优势可以避免单条SQL过长问题，也会避免当某条记录插入失败而需要全部重新插入的问题。
	// model
	csInBatch := []gorm_study.Content{
		{Subject: "分批插入标题1"},
		{Subject: "分批插入标题2"},
		{Subject: "分批插入标题3"},
		{Subject: "分批插入标题4"},
		{Subject: "分批插入标题5"},
	}
	inBatchResult1 := DB.CreateInBatches(&csInBatch, 2)
	if result1.Error != nil {
		log.Fatal(inBatchResult1.Error)
	}
	fmt.Println(inBatchResult1.RowsAffected)
	for _, c := range csInBatch {
		fmt.Println(c.ID)
	}
}
