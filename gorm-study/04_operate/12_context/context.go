package _2_context

import (
	"context"
	"fmt"
	gorm_study "gorm-study"
	"log"
	"time"
)

//## Context支持
//
//GORM支持Context：
//
//使用 DB.WithContext() 或 &Session{Context: ctx} 字段进行配置。
//
//示例，控制执行时间的Context，在预设的时间没有执行完毕的话，DB会返回错误

type Content gorm_study.Content

func ContextTOCancel() {
	DB := gorm_study.DB
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	var cs []Content
	if err := DB.WithContext(ctx).Limit(10).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cs)
}
