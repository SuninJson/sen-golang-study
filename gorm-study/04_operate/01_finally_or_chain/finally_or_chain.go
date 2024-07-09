package finally_or_chain

// GORM 中有两种类型的方法：
// GORM 中有两种类型的方法：
//
//- 终结方法（Finisher Method），用于生成并执行当前语句的方法
//
//- ```go
//    Create, First, Take, Find, Save, Update, Delete, Scan, Row, Rows
//    ```
//
//- 链式方法（Chain Method），用于将子句（Clauses）加入当前语句的方法
//
//- ```go
//    Select, Table, Where, Joins, Group, Having, Order, Limit, Offset, Debug
//    ```

import (
	"errors"
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"log"
)

func OperatorType() {

	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.User{})

	var users []gorm_study.User

	// 分步操作
	//query := DB.Where("birthday IS NOT NULL")
	//query.Where("email like ?", "@163.com%")
	//query.Order("name DESC")
	//query.Find(&users)

	// 一步操作
	err := DB.Where("birthday IS NOT NULL").Where("email like ?", "@163.com%").Order("name DESC").Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 当 `First`、`Last`、`Take` 方法（这几个方法都是查找一条的方法）找不到记录时，GORM 会返回 `ErrRecordNotFound` 错误。这个是GORM的行为，数据库层面在没有记录是不会响应错误。
			//注意，当 `First`、`Last`、`Take` 方法存在错误时，不一定是ErrRecordNotFound类型，也有可能是其他类型。若有需要判定是否为ErrRecordNotFound类型错误，可以通过 errors.Is() 方法进行判断。
			log.Println("Record Not Found")
		} else {
			log.Fatal(err)
		}
	}
}
