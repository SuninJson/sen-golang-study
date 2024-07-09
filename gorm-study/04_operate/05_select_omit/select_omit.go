package _5_select_omit

import (
	gorm_study "gorm-study"
	"time"
)

func SelectCol() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{})

	c := gorm_study.Content{}
	c.Views = 99
	c.Likes = 7
	c.Subject = "标题"
	now := time.Now()
	c.PublishTime = &now

	// 选择字段
	DB.Select("Subject", "Views", "CreatedAt").Create(&c)
	// INSERT INTO `msb_content` (`created_at`,`updated_at`,`subject`,`views`) VALUES ('2023-04-11 17:51:39.895','2023-04-11 17:51:39.895','标题',99)

	// 忽略字段
	DB.Omit("Subject", "Views", "CreatedAt").Create(&c)
	// INSERT INTO `msb_content` (`updated_at`,`deleted_at`,`likes`,`publish_time`) VALUES ('2023-04-11 17:52:29.034',NULL,7,'2023-04-11 17:52:29.032')
}
