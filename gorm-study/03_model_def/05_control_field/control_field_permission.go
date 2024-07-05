package _5_control_field

// 写权限控制
// (1) <-:false 无写入
// (2) <-:create 仅创建
// (3) <-:update 仅更新
//
// 读权限控制
// (1) ->:false 无读取

type TestControlFieldStruct struct {
	Schema      string
	Host        string `gorm:"<-:false"`
	Path        string `gorm:"->:false"`
	Url         string `gorm:"->:create"`
	QueryString string `gorm:"<-:update"`
}
