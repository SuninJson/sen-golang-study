package _5_control_field

import (
	"fmt"
	gormStudy "gorm-study"
	"log"
)

// `-` 标签标示忽略该字段的数据库操作。
//
// 例如，网络传递的字段是 url，但数据表中存储的字段是：schema, host, path, query_string。
// 就是将URL拆解成多个部分进行存储，但网络传输时使用的是url字段。
// 那么该url字段，就应该忽略数据库的读写（包括迁移）操作，示例：

type Service struct {
	Schema      string
	Host        string
	Path        string
	QueryString string
	Url         string `gorm:"-"`
}

func CreateServiceTable() {
	if err := gormStudy.DB.AutoMigrate(&Service{}); err != nil {
		log.Fatalln(err)
	}
}

func ServiceCRUD() {
	s := &Service{}
	s.Schema = "https"
	s.Url = "https://www.sen.com/study?key=value"

	if !queryService() {
		gormStudy.DB.Create(s)
		fmt.Printf("system service obj:%+v\n", s)
		queryService()
	} else {
		fmt.Printf("system service obj:%+v\n", s)
	}
}

func queryService() bool {
	dbS := &Service{}
	if err := gormStudy.DB.First(dbS, "`Schema` = 'https'").Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("db service obj:%+v\n", dbS)

	return dbS.Schema != ""
}
