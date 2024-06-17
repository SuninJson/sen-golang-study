package gorm_study

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

const (
	GlobalLogger = 1
	FileLogger   = 2
)

var DB *gorm.DB

func init() {
	// 定义DSN
	const dsn = "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接服务器（池）
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         customLogger(FileLogger),
		NamingStrategy: namingStrategy(),
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func customLogger(initLoggerType int) logger.Interface {
	var loggerInterface logger.Interface
	switch initLoggerType {
	case GlobalLogger:
		// 可以在 `gorm.Open` 时设置全局日志级别，
		// 或者通过 `db.Config.Logger`完成配置：DB.Config.Logger = logger.Default.LogMode(logger.Info)
		loggerInterface = logger.Default.LogMode(logger.Info)
	case FileLogger:
		// 0644的意思：
		// （1）开头的0是一个约定，用来明确后面跟随的是一个八进制数。它本身并不直接代表任何文件权限，而是作为一个指示符，告诉解析这个数值的系统（或人）后面的数字应当被解释为八进制形式。
		// （2）第2位6：4表示读权限（r）2表示写权限（w）因此，6是4+2，意味着文件所有者有读和写权限。
		// （3）第3位4：表示所属组用户只有读权限。
		// （4）第4位4：表示其他用户的只有读权限
		logWriter, _ := os.OpenFile("./sql.log", os.O_CREATE|os.O_APPEND, 0644)
		// LstdFlags是以下两个常用标志的组合：
		// log.Ldate：这个标志表示在每次日志输出时包含日期信息
		// log.Ltime：这个标志表示在每次日志输出时包含时间信息
		loggerInterface = logger.New(log.New(logWriter, "\n", log.LstdFlags),
			logger.Config{
				// 慢查询阈值 200ms
				SlowThreshold: 200 * time.Millisecond,
				// 日志级别
				LogLevel: logger.Info,
				// 是否忽略记录不存在的错误
				IgnoreRecordNotFoundError: false,
				// 不彩色化
				Colorful: false,
			})
	default:
		loggerInterface = logger.Default
	}

	return loggerInterface
}

func namingStrategy() schema.NamingStrategy {
	return schema.NamingStrategy{
		// 表名前缀
		TablePrefix: "gorm_study_",
		// 是否单数表名，默认为复数表名
		SingularTable: true,
		// 替换器，用于替换特定字符串
		NameReplacer: nil,
		// 是否为name_casing形式
		NoLowerCase: true,
	}
}
