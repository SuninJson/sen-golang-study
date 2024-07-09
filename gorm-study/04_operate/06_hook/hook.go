package _6_hook

import (
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// 钩子，Hook，在特定执行时间点执行的方法。类似于事件驱动。CRUD都有对应的钩子方法。
// 如果我们定义了钩子方法，那么在操作时就会执行对应的钩子方法。
// 创建时，可用的钩子方法有：
//	// 开始事务
//	BeforeSave()
//	BeforeCreate()
//	Create()
//	AfterCreate()
//	AfterSave()
//	// 提交或回滚事务
// 依据上面的顺序执行。其中BeforeSave和AfterSave是创建和更新操作通用的。

// BeforeCreate 示例：
func (c *Content) BeforeCreate(db *gorm.DB) error {
	// 业务代码
	if c.PublishTime == nil {
		now := time.Now()
		c.PublishTime = &now
	}

	// 配置代码
	db.Statement.AddClause(clause.OnConflict{UpdateAll: true})

	return nil
}

type Content struct {
	gorm.Model

	Subject string
	// 通过default标签来设置默认值，当字段为类型零值时，触发使用默认值
	Likes       uint  `gorm:"default 100"`
	LikesPoint  *uint `gorm:"default 99"`
	Views       uint
	PublishTime *time.Time
}

func CreateUseHook() {
	DB := gorm_study.DB
	DB.AutoMigrate(&Content{})
	c1 := Content{}

	err := DB.Create(&c1).Error
	if err != nil {
		log.Fatal(err)
	}
	//INSERT INTO `msb_content` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`views`,`publish_time`) VALUES ('2023-04-11 18:44:56.62','2023-04-11 18:44:56.62',NULL,'',0,0,'2023-04-11 18:44:56.62') ON DUPLICATE KEY UPDATE `updated_at`='2023-04-11 18:44:56.62',`deleted_at`=VALUES(`deleted_at`),`subject`=VALUES(`subject`),`likes`=VALUES(`likes`),`views`=VALUES(`views`),`publish_time`=VALUES(`publish_time`)

}
