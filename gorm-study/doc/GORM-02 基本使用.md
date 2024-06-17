# 基本使用

## 启动MySQL服务器等待使用

采用任意方式都可以，保证有可用的MySQL服务器即可。

本课程采用docker-compose方式管理MySQL服务器。

在案例目录创建文件：docker-compose.yml

键入如下内容：

```yml
services:
  db:
    container_name: gormExampleMySQL
    image: mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: gormExample
    ports:
      - "3306:3306"
    volumes:
      - ./mysql/data:/var/lib/mysql
```

启动Docker.

docker compose

```shell
# 启动
docker-compose up -d
# 关闭
docker-compose down

> docker-compose up -d
[+] Running 2/2
 - Network gormexample_default  Created                                                          0.7s
 - Container gormExampleMySQL   Started
```

测试：

```shell
# 输入密码
docker exec -it gormExampleMySQL mysql -p

# 行内密码
docker exec -it gormExampleMySQL mysql -psecret
```

## 使用流程

1. 安装GORM和驱动
2. 连接数据库服务器
3. 定义模型
4. 使用迁移管理表
5. 完成操作
6. 调试、监控

## 抽象层和数据库驱动

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1680502356013/db45129fd85e41feab26038078873c4d.png)

抽象层提供操作接口，具体的数据库操作实现由数据库对应的驱动实现。

```go
import (
  // 驱动
  "gorm.io/driver/mysql"
  // 抽象层
  "gorm.io/gorm"
)
```

gorm 为了方便使用，整合了以下驱动：

```go
gorm.io/driver/sqlite
gorm.io/driver/mysql
gorm.io/driver/postgres
gorm.io/driver/sqlserver
```

因此使用GROM时，需要安装gorm和对应的驱动。

## 安装GORM和MySQL驱动

我们创建一个示例module，初始化mod后，安装gorm和驱动：

```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 连接数据库服务器

示例代码：

```go
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func BasicUsage() {
	dsn := "root:secret@tcp(127.0.0.1:3306)/gormExample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db:", db)
}
```

GORM使用连接池技术管理连接。

方法：

```go
func Open(dialector Dialector, opts ...Option) (db *DB, err error)
```

用于初始化数据库回话，基于拨号器 Dialector,和选项。返回 `*gorm.DB` 对象和错误。

Dialector，是通过驱动的Open方法创建的，以MySQL为例：

```go
func Open(dsn string) gorm.Dialector
```

需要提供DSN参数。

DSN，Data Source Name，数据源名称，用于描述在哪里找到数据。MySQL的DSN信息：

[MySQL DSN 说明](https://github.com/go-sql-driver/mysql#dsn-data-source-name)

## 定义模型

模型，就是一个struct类型，示例代码为：

```go
type Article struct {
	gorm.Model

	Subject     string
	Likes       uint
	Published   bool
	PublishTime time.Time
	AuthorID    uint
}
```

通常情况下，要嵌入 gorm.Model，用于保有核心字段。

gorm.Model

```go
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}
```

## 迁移表结构

Migrate，迁移指的是通过模型来确定表结构。通常使用 AutoMigrate()完成迁移：

```go
// 迁移
db.AutoMigrate(&Article{})
```

执行以上代码后，对应的表结构就自动创建出来了。 AutoMigrate 会创建表、缺失的外键、约束、列和索引。 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。 出于保护您数据的目的，它 **不会** 删除未使用的列。

使用 mysql 客户端，查看表结构的变化：

```
mysql> show create table articles;
```

## 基本CRUD操作

基于模型对象，完成CRUD，模型对象，也就是Article类型的数据示例。Article{}

### 初始化DB对象

```go
var DB *gorm.DB

func init() {
	// 定义DSN
	const dsn = "root:secret@tcp(127.0.0.1:3306)/gormExample?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接服务器（池）
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
```

### Create，创建

步骤：

- 创建模型对象
- db.Create() 完成insert操作

示例：

```go
func Create() {
	// 构建Article类型数据
	article := &Article{
		Subject:     "GORM 的 CRUD 基础操作",
		Likes:       0,
		Published:   true,
		PublishTime: time.Now(),
		AuthorID:    42,
	}

	// DB.Create 完成数据库的insert
	if err := DB.Create(article).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}

```

测试执行

```shell
func TestCreate(t *testing.T) {
	Create()
}

go test -run Create  
&{{1 2023-04-03 19:22:13.46 +0800 CST 2023-04-03 19:22:13.46 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}}
 GORM 的使用 0 true 2023-04-03 19:22:13.4586223 +0800 CST m=+0.015389201 42}
PASS
ok      gormExample     0.056s

```

在数据库中查看：

```mysql
SELECT * FROM articles\G
```

注意，gorm.Model 嵌入的字段：

- ID，auto_increment
- created_at和updated_at为当前时间
- deleted_at为null

### Retrieve，获取

步骤：

- 给定查询条件，例如PK
- 选择查询单个还是多个
  - Find() 多个
  - First() 单个

示例：

```go
func Retrieve(id uint) {
	// 初始化Article模型，零值
	article := &Article{}

	// DB.First()
	if err := DB.First(article, id).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}
```

测试：

```shell
func TestRetrieve(t *testing.T) {
	Retrieve(1)
}

go test -run Retrieve
&{{1 2023-04-03 19:22:13.46 +0800 CST 2023-04-03 19:22:13.46 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}}
 GORM 的使用 0 true 2023-04-03 19:22:13.459 +0800 CST 42}
PASS
ok      gormExample     0.052s
```

此处仅仅展示根据主键查询。

使用错误的id，测试查询出错的情况。

```shell
func TestRetrieve(t *testing.T) {
	Retrieve(3)
}

go test -run Retrieve

2023/04/03 19:34:46 D:/apps/courses/gormExample/basic.go:62 record not found
[3.809ms] [rows:0] SELECT * FROM `articles` WHERE `articles`.`id` = 3 AND `articles`.`deleted_at` IS NULL ORD
ER BY `articles`.`id` LIMIT 1
2023/04/03 19:34:46 record not found
exit status 1   
FAIL    gormExample     0.057s
```

### Update, 更新

步骤：

- 先确定更新的对象
- 设置对象属性字段
- 将对象存储

示例：

```go
func Update() {
	// 获取需要更新的对象
	article := &Article{}
	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}

	// 更新对象字段
	article.AuthorID = 23
	article.Likes = 101
	article.Subject = "新的文章标题"

	// 存储，DB.Save()
	if err := DB.Save(article).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(article)
}
```

测试：

```shell
func TestUpdate(t *testing.T) {
	Update()
}

go test -run Update
&{{1 2023-04-03 19:22:13.46 +0800 CST 2023-04-03 19:38:02.379 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}
} 新的文字标题 100 true 2023-04-03 19:22:13.459 +0800 CST 42}
PASS
ok      gormExample     0.066s
```

本例子中，体现的典型的ORM更新。语法上更新对象。gorm同样支持基于条件的更新。

### Delete, 删除

步骤：

- 确定删除的模型对象
- 删除

示例：

```go
func Delete() {
	// 获取模型对象
	article := &Article{}
	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}

	// DB.Delete() 删除
	if err := DB.Delete(article).Error; err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println("article was deleted")
}


// 当然也可以
DB.Delete(&Article{}, 1)
```

测试：

```shell
func TestDelete(t *testing.T) {
	Delete()
}

go test -run Delete  
article was deleted
PASS
ok      gormExample     0.069s
```

也可以通过主键ID删除，但本例子主要体现ORM的概念。同时在实际业务逻辑中，删除前，往往要对数据做额外的处理，通常也会先查询到的。

查看数据表，会发现记录的deleted_at字段设置了时间，表示该记录被删除。

```shell
mysql> select * from articles\G
*************************** 1. row ***************************
          id: 1
  created_at: 2023-04-03 19:22:13.460
  updated_at: 2023-04-03 19:38:02.379
  deleted_at: 2023-04-03 19:42:39.757
     subject: 新的文字标题
       likes: 100
   published: 1
publish_time: 2023-04-03 19:22:13.459
   author_id: 42
1 row in set (0.00 sec)
```

## Debug和日志

### db.Debug

`db.Debug()`方法用于将当前操作的log级别调整为 info 级别，就是可以获取当前执行的SQL：

```go
func Debug() {
	article := &Article{
		Subject:     "Article Subject",
		PublishTime: time.Now(),
	}
	if err := DB.Debug().Create(article).Error; err != nil {
		log.Fatal(err)
	}

	if err := DB.Debug().First(article, 1).Error; err != nil {
		log.Fatal(err)
	}
}
```

测试：

```powershell
func TestDebug(t *testing.T) {
	Debug()
}

> go test -run Debug

2023/04/04 13:16:59 D:/apps/courses/gormExample/basic.go:100
[8.999ms] [rows:1] INSERT INTO `articles` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`publishe
d`,`publish_time`,`author_id`) VALUES ('2023-04-04 13:16:59.102','2023-04-04 13:16:59.102',NULL,'Article Subj
ect',0,false,'2023-04-04 13:16:59.1',0)

2023/04/04 13:16:59 D:/apps/courses/gormExample/basic.go:104
[3.000ms] [rows:1] SELECT * FROM `articles` WHERE `articles`.`id` = 4 AND `articles`.`deleted_at` IS NULL AND
 `articles`.`id` = 4 ORDER BY `articles`.`id` LIMIT 1
PASS
ok      gormExample     0.055s

```

### 全局配置日志级别

可以在 `gorm.Open` 时设置全局日志级别，或者通过 `db.Config.Logger`完成配置：

```go
// gorm.Open
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})

// DB.Config.Logger
DB.Config.Logger = logger.Default.LogMode(logger.Info)
```

配置完成后，后续的操作都会使用Info级别的Log。

示例，更新 init() 方法：

```go
func init() {
	// 定义DSN
	const dsn = "root:secret@tcp(127.0.0.1:3306)/gormExample?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接服务器（池）
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置Info级别的默认日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
```

之前之前的CRUD操作，都会存在SQL输出了。

### 日志级别

GORM定义了4个级别：

- Info，logger.Info，全部消息
- Warn，logger.Warn，默认，警告
- Error, logger.Error，错误
- Silent, logger.Silent，静默

### 配置日志选项

GORM有默认的Logger实现，如下：

```go
Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
```

**我们通过自定义Logger，即可控制选项：**

```go
// 自定义日志
var logWriter io.Writer

func init() {
	// 定义DSN
	const dsn = "root:secret@tcp(127.0.0.1:3306)/gormExample?charset=utf8mb4&parseTime=True&loc=Local"
	// 初始化logWriter
	logWriter, _ = os.OpenFile("./sql.log", os.O_CREATE|os.O_APPEND, 0644)
	customLogger := logger.New(log.New(logWriter, "\n", log.LstdFlags),
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
	// 连接服务器（池）
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置为自定义的日志
		Logger: customLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
```

后续执行操作时，日志被记录在./sql.log文件中:

2023/04/04 15:11:24 D:/apps/mashibing/gormExample/basic.go:105
[12.441ms] [rows:1] INSERT INTO `articles` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`published`,`publish_time`,`author_id`) VALUES ('2023-04-04 15:11:24.438','2023-04-04 15:11:24.438',NULL,'GORM 的 CRUD 基础操作',0,true,'2023-04-04 15:11:24.435',42)

2023/04/04 15:12:05 D:/apps/mashibing/gormExample/basic.go:105
[11.999ms] [rows:1] INSERT INTO `articles` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`published`,`publish_time`,`author_id`) VALUES ('2023-04-04 15:12:05.032','2023-04-04 15:12:05.032',NULL,'GORM 的 CRUD 基础操作',0,true,'2023-04-04 15:12:05.029',42)

2023/04/04 15:12:27 D:/apps/mashibing/gormExample/basic.go:105
[17.603ms] [rows:1] INSERT INTO `articles` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`published`,`publish_time`,`author_id`) VALUES ('2023-04-04 15:12:27.809','2023-04-04 15:12:27.809',NULL,'GORM 的 CRUD 基础操作',0,true,'2023-04-04 15:12:27.806',42)
