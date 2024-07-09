package _1_session

import (
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"log"
)

//### 会话模式的基本使用
//
//GORM支持链式操作，意味着前边的操作会对后影响后边的调用，这在很多时候很好用。
//但当我们需要连续执行多次查询时，就可能出问题，导致子句的重叠，
//使用Session会话，可以让某些子句重用

type Content gorm_study.Content

func SessionNew() {
	DB := gorm_study.DB

	// 使用Session方法启动了新的Session，意味着，db对象可以保持会话开启前的状态。继续使用db对象时，到执行终结方法前，都是以该会话状态为初始化状态的。这就可以保证，会话中的子句可以重用
	// 将Session方法前的配置，记录到了当前的会话中
	// 后边再次调用db的方法直到终结方法，会保持会话中的子句选项
	// 执行完终结方法后，再次调用db的方法到终结方法，可以重用会话中的子句选项。
	db := DB.Model(&Content{}).Where("views > ?", 100).Session(&gorm.Session{})

	// 连续执行查询
	// 1
	// Where("views > ?", 100).Where("likes > ?", 9)
	var cs1 []Content
	db.Where("likes > ?", 9).Find(&cs1)
	// [4.633ms] [rows:0] SELECT * FROM `content` WHERE views > 100 AND likes > 9 AND `content`.`deleted_at` IS NULL

	// 2,找到likes<5
	// Where("views > ?", 100).Where("likes < ?", 5)
	var cs2 []Content
	db.Where("likes < ?", 5).Find(&cs2)
	// [3.846ms] [rows:0] SELECT * FROM `content` WHERE views > 100 AND likes < 5 AND `content`.`deleted_at` IS NULL
}

//通过配置Session来禁用Hook

func SessionOptionDisableHook() {
	DB := gorm_study.DB
	db := DB.Model(&Content{}).Session(&gorm.Session{
		// 通过修改 SkipHooks: true， 可以看到是否有输出content before create hook。
		SkipHooks: true,
	})
	db.Save(&Content{Subject: "no hooks"})
}

func (c *Content) BeforeCreate(db *gorm.DB) error {
	log.Println("content before create hook")
	return nil
}

//#### 预编译模式
//
//在执行 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
// 当连续执行结构一致，但数据不同的SQL时，可以利用预编译的SQL缓存，提升效率。

func SessionOptionPrepareStmt() {
	DB := gorm_study.DB
	// prepare
	db := DB.Model(&Content{}).Session(&gorm.Session{
		PrepareStmt: true,
	})

	stmtManger, ok := db.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		log.Fatalln("*gorm.PreparedStmtDB assert failed")
	}
	log.Println(stmtManger.PreparedSQL)

	var c1 Content
	db.First(&c1, 13)
	log.Println(stmtManger.PreparedSQL)
	var c2 Content
	db.First(&c2, 13)
	var c3 Content
	db.First(&c3, 13)
}

//#### Debug()
//
//Debug 利用将日志级别更改为logger.Info来实现。
//
//func (db *DB) Debug() (tx *DB) {
//
//	return db.Session(&Session{
//		Logger: db.Logger.LogMode(logger.Info),
//	})
//}
