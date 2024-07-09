package _0_execute_sql

import (
	gorm_study "gorm-study"
	"log"
)

// ## 原生SQL
//
//原生SQL，指的是我们使用 标准SQL语句完成DB操作。
//
//当需要执行原生的SQL的时，将SQL中的参数使用占位符占位，之后提供变量，拼凑构造成完整的SQL进行执行。
//
//查询时，使用各种子句方法，其实就是再构造SQL的各个部分。
//
//执行SQL，分为两种：
//
//- 查询类型，存在返回数据的，典型就是select, show, desc等。利用DB.Raw()和DB.Scan()完成，通常需要定义响应结果结构。
//- 非查询类，没有返回数据的， insert， update，delete，DDL等。利用DB.Exec()完成。
//
// ### sql.Row和sql.Rows
//
//若需要在原生SQL查询时，使用标准结构sql.Row和sql.Rows时，调用DB.Row()和DB.Rows()方法：
//func (db *DB) Rows() (*sql.Rows, error)
//func (db *DB) Row() *sql.Row
//获取Row或Rows后，需要扫描到结果变量来使用：
//
//- 将Row扫描到单独变量，row.Scan
//- 将Row扫描到整体结构体类型，DB.ScanRow()
//- 判断Rows中是否存在记录，rows.Next()
//示例：

func RawSelect() {
	DB := gorm_study.DB
	// 结果类型
	type Result struct {
		ID           uint
		Subject      string
		Likes, Views int
	}
	var rs []Result

	// SQL
	sql := "SELECT `id`, `subject`, `likes`, `views` FROM `content` WHERE `likes` > ? ORDER BY `likes` DESC LIMIT ?"

	// 执行SQL，并扫描结果
	if err := DB.Raw(sql, 99, 12).Scan(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	// [8.298ms] [rows:12] SELECT `id`, `subject`, `likes`, `views` FROM `content` WHERE `likes` > 99 ORDER BY `likes` DESC LIMIT 12

	log.Println(rs)
}

func RawExec() {
	DB := gorm_study.DB
	// SQL
	sql := "UPDATE `content` SET `subject` = CONCAT(`subject`, '-new postfix') WHERE `id` BETWEEN ? AND ?"

	// 执行，获取结果
	result := DB.Exec(sql, 30, 40)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// [13.369ms] [rows:10] UPDATE `content` SET `subject` = CONCAT(`subject`, '-new postfix') WHERE `id` BETWEEN 30 AND 40

	log.Println(result.RowsAffected)

}
