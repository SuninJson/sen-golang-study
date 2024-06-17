# 模型定义

模型，model，在ORM中与关系表映射的应用程序对象，不同语言使用不同数据类型实现，通常为对象类型，在Go语言中是struct结构体类型。具体的struct类型实例就是具体的某个表的模型对象，ORM中与表中记录映射。

```go
type Article struct {
    gorm.Model
	Subject     string
}
```

## 表名定义

### 表名约定

在默认情况下，GORM有约定，使用小写+下划线（蛇形命名）的复数形式作为表名，例如：

| Model        | Table           |
| ------------ | --------------- |
| Post         | posts           |
| Category     | categories      |
| Box          | boxes           |
| PostCategory | post_categories |

示例：

```go
type Post struct {
	gorm.Model
}
type Category struct {
	gorm.Model
}
type PostCategory struct {
	gorm.Model
}
type Box struct {
	gorm.Model
}

func Migrate() {
	if err := DB.Debug().AutoMigrate(&Post{}, &Category{}, &PostCategory{}, &Box{}); err != nil {
		log.Fatal(err)
	}
}

// =============================
// *_test.go
func TestMigrate(t *testing.T) {
	Migrate()
}
> go test -run Migrate

```

查看数据表：

```mysql
mysql> show tables;
+-----------------------+
| Tables_in_gormexample |
+-----------------------+
| boxes                 |
| categories            |
| post_categories       |
| posts                 |
+-----------------------+
4 rows in set (0.01 sec)
```

### 命名策略自定义

GORM的约定，就是使用内置的命名策略来实现的。命名策略是实现了Namer接口的类型：

```go
// gorm.io/gorm/schema

// Namer namer interface
type Namer interface {
	TableName(table string) string
	SchemaName(table string) string
	ColumnName(table, column string) string
	JoinTableName(joinTable string) string
	RelationshipFKName(Relationship) string
	CheckerName(table, column string) string
	IndexName(table, column string) string
}
```

默认的策略（约定）的TableName实现如下：

```go
// gorm.io/gorm/schema

// NamingStrategy tables, columns naming strategy
type NamingStrategy struct {
    // 表名前缀
	TablePrefix   string
    // 是否单数表名
	SingularTable bool
    // 替换器，用于替换特定字符串
	NameReplacer  Replacer
    // 是否为sname_casing形式
	NoLowerCase   bool
}

// 将模型名转为表名
func (ns NamingStrategy) TableName(str string) string {
	// 单数
    if ns.SingularTable {
		return ns.TablePrefix + ns.toDBName(str)
	}
    // 复数
	return ns.TablePrefix + inflection.Plural(ns.toDBName(str))
}

// ColumnName convert string to column name
func (ns NamingStrategy) ColumnName(table, column string) string {
	return ns.toDBName(column)
}
```

### 表名前缀

使用默认的命名策略，来增加表名前缀。

通过修改gorm.Open时的配置实现：

```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    NamingStrategy: schema.NamingStrategy{
        TablePrefix:   "msb_",
        SingularTable: true,
        NameReplacer:  nil,
        NoLowerCase:   true,
    },
})
```

测试生成的数据表：

```mysql
mysql> show tables;
+-----------------------+
| Tables_in_gormexample |
+-----------------------+
| boxes                 |
| categories            |
| post_categories       |
| posts                 |
| msb_Category          |
| msb_Post              |
| msb_PostCategory      |
| msb_Box               |
+-----------------------+
```

实际项目中，比较喜欢使用小写表名的。所以 NoLowerCase: false比较常见。

> 旧版v1的修改表名前缀的方案，是：
>
> ```go
> gorm.DefaultTableNameHanlder
> ```

### 表名自定义

若需要使用自定义规则表名，模型需要实现 Tabler 接口，Tabler接口：

```go
// gorm.io/gorm/schema
type Tabler interface {
	TableName() string
}
```

示例：

```go
type Box struct {
	gorm.Model
}

func (Box) TableName() string {
	return "my_box"
}
```

得到的表名就是 my_box。

### 临时指定表名

方法：

```go
func (db *DB) Table(name string, args ...interface{}) (tx *DB)DB.Table
```

用于在一个执行周期内，指定临时表名。若配合Migrate使用，可以设置所迁移的表的名字。例如：

```go
DB.Table("temp_articles").AutoMigrate(&Article{})
```

Table方法常用于SQL的执行中。

## 字段类型映射

模型的字段类型可为：

- Go基本数据类型，典型的：`bool, int, uint, float32/64, string, time.Time, []byte`
- Go基本数据类型的指针类型，典型的：`*int, *uint, *float32/64, *string, *time.Time`
- 实现了Scanner和Valuer接口的自定义类型，典型的database/sql包定义的sql.NullType系列类型

示例：

```go
type TypeMap struct {
	gorm.Model

	FInt       int
	FUInt      uint
	FFloat32   float32
	FFloat64   float64
	FString    string
	FTime      time.Time
	FByteSlice []byte

	FIntP     *int
	FUIntP    *uint
	FFloat32P *float32
	FFloat64P *float64
	FStringP  *string
	FTimeP    *time.Time
}
```

基于以上模型，迁移生成的表结构：

```mysql
mysql> desc type_maps;
+--------------+-----------------+------+-----+---------+----------------+
| Field        | Type            | Null | Key | Default | Extra          |
+--------------+-----------------+------+-----+---------+----------------+
| id           | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at   | datetime(3)     | YES  |     | NULL    |                |
| updated_at   | datetime(3)     | YES  |     | NULL    |                |
| deleted_at   | datetime(3)     | YES  | MUL | NULL    |                |
| f_int        | bigint          | YES  |     | NULL    |                |
| f_uint       | bigint unsigned | YES  |     | NULL    |                |
| f_float32    | float           | YES  |     | NULL    |                |
| f_float64    | double          | YES  |     | NULL    |                |
| f_string     | longtext        | YES  |     | NULL    |                |
| f_time       | datetime(3)     | YES  |     | NULL    |                |
| f_byte_slice | longblob        | YES  |     | NULL    |                |
| f_int_p      | bigint          | YES  |     | NULL    |                |
| f_uint_p     | bigint unsigned | YES  |     | NULL    |                |
| f_float32_p  | float           | YES  |     | NULL    |                |
| f_float64_p  | double          | YES  |     | NULL    |                |
| f_string_p   | longtext        | YES  |     | NULL    |                |
| f_time_p     | datetime(3)     | YES  |     | NULL    |                |
+--------------+-----------------+------+-----+---------+----------------+
```

以上就是MySQL数据对应的类型。注意不同数据库有不同的实现，但以上类型是大多数数据库支持的通用类型。

## 指针类型和非指针类型的区别

上面的案例中，*T和T对应的MySQL类型是一致的，这是因为在MySQL的数据字段类型层面，没有指针的概念。

但在Go的类型中，存在区别：

- T，不能表示NULL，如果数据库中字段的值为NULL，那么映射到模型对象的字段为类型零值。同理，无法设置记录字段的值为NULL。
- *T，可以使用nil表示NULL，意味着如果数据库中字段的值为NULL，那么映射到模型对象的字段为nil。同理，可以通过将字段设置为nil，来设置字段为NULL。

插入一条ID:1的记录作为测试：

```mysql
mysql> INSERT INTO `msb_type_map` (`id`) VALUES (1);
```

示例，使用TypeMap模型，完成查询：

```go
func PointerDiff() {
	tm := &TypeMap{}
	fmt.Printf("%+v\n", tm)

	DB.First(tm, 1)
	fmt.Printf("%+v\n", tm)
}
```

比较模型零值时，不同字段的查询。

比较表中记录字段为Null时，模型字段的差异。

```go
&{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time
:0001-01-01 00:00:00 +0000 UTC Valid:false}} FInt:0 FUInt:0 FFloat32:0 FFloat64:0 FString: FTime:0001-01-01 0
0:00:00 +0000 UTC FByteSlice:[] FIntP:<nil> FUIntP:<nil> FFloat32P:<nil> FFloat64P:<nil> FStringP:<nil> FTime
P:<nil>}
&{Model:{ID:1 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time
:0001-01-01 00:00:00 +0000 UTC Valid:false}} FInt:0 FUInt:0 FFloat32:0 FFloat64:0 FString: FTime:0001-01-01 0
0:00:00 +0000 UTC FByteSlice:[] FIntP:<nil> FUIntP:<nil> FFloat32P:<nil> FFloat64P:<nil> FStringP:<nil> FTime
P:<nil>}

```

**结论：若表中字段可以为NULL，那么应该使用*T指针类型，映射字段，nil表示NULL。**

## 自定义字段类型

除标准Go类型外，还可以使用实现database/sql包中Scanner和database/sql/driver包中Valuer接口的自定义类型，以便让 GORM 知道如何将该类型接收、保存到数据库。其中：

Scanner接口：

```go
// database/sql/sql.go

// Scanner is an interface used by Scan.
type Scanner interface {
	// Scan 从数据库中分配一个值，用于查询时设置字段值
	Scan(src any) error
}
```

Valuer接口：

```go
// database/sql/driver/types.go

type Valuer interface {
	// Value 用于获取模型字段的值
	Value() (Value, error)
}
```

以 sql.NullTime 为例：

```go
// database/sql/sql.go

// NullTime represents a time.Time that may be null.
// NullTime implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (n *NullTime) Scan(value any) error {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Time, value)
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
```

**典型的sql.NullType**

```go
sql.NullTime
sql.NullByte
sql.NullBool
sql.NullFloat64
sql.NullInt16
sql.NullInt64
sql.NullString
sql.NullInt32
```

结构体字段 Valid:false 表示，数据表值是NULL。使用模型时，可以根据对应的Valid字段，判定数据表中数据是否为NULL。

**第三方自定义类型**

```go
github.com/google/uuid
```

install：

```shell
go get github.com/google/uuid
```

**示例**：

```go
type CustomTypeModel struct {
	gorm.Model

	FTime     time.Time
	FNullTime sql.NullTime

	FString     string
	FNullString sql.NullString

	FUUID     uuid.UUID
	FNullUUID uuid.NullUUID
}

func CustomType() {
	//id := uuid.UUID{}
	//id.Scan()  // Scanner
	//id.Value() // Valuer

	// 初始化模型
	ctm := &CustomTypeModel{}
	// 迁移数据表
	DB.AutoMigrate(ctm)

	// 创建
	ctm.FTime = time.Now()         // 当前时间
	ctm.FNullTime = sql.NullTime{} // 零值，Valid默认为false
	ctm.FString = ""
	ctm.FNullString = sql.NullString{}

	ctm.FUUID = uuid.New()
	ctm.FNullUUID = uuid.NullUUID{}

	DB.Create(ctm)

	// 查询
	DB.First(ctm, ctm.ID)

	// 判定字段是否为NULL
	if ctm.FString == "" {
		fmt.Println("FString is NULL")
	} else {
		fmt.Println("FString is NOT NULL")
	}

	if ctm.FNullString.Valid == false {
		fmt.Println("FNullString is NULL")
	} else {
		fmt.Println("FNullString is NOT NULL")
	}
}
```

测试结果：

```shell
#// 测试函数
func TestCustomType(t *testing.T) {
	CustomType()
}

# 运行测试
> go test -run UUIDCreate
&{Model:{ID:4 CreatedAt:2023-04-06 11:45:28.752 +0800 CST UpdatedAt:2023-04-06 11:45:28.752 +0800 CST Deleted
At:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} FTime:2023-04-06 11:45:28.751 +0800 CST FString: FNullTi
me:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false} FNullString:{String: Valid:false} FUUID:04c608f2-6de3-4f0
5-a86b-c699b3560a6f}
字段为NULL
PASS  
ok      gormExample     0.096s


# mysql 结果
mysql> select * from msb_custom_type_model\G
*************************** 1. row ***************************
           id: 1
   created_at: 2023-04-06 11:37:26.584
   updated_at: 2023-04-06 11:37:26.584
   deleted_at: NULL
       f_time: 2023-04-06 11:37:26.583
     f_string:
  f_null_time: NULL
f_null_string: NULL
       f_uuid: af5a0505-0f07-495e-ac03-250cd5ccc8bf
1 row in set (0.00 sec)
```

## 字段标签设置字段属性

结构体字段的标签 gorm，可用来对GROM的行为进行设置。

标签语法

```go
type ModelType struct {
	Field Type `gorm:"key:value;key;"`
}
```

常用的字段标签如下表：

| key                    | 功能                                                                                      | 说明                                                                                                       |
| ---------------------- | ----------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| column                 | 列名                                                                                      |                                                                                                            |
| type                   | 类型                                                                                      | 根据是需求设置字段类型，推荐兼容性好的类型。所有数据库都支持 bool、int、uint、float、string、time、bytes。 |
| not null               | 设置为NOT NULL，默认为NULL，不需要指定                                                    |                                                                                                            |
| default                | 默认值                                                                                    | 根据需求设置                                                                                               |
| autoCreateTime         | 创建时追踪当前时间，支持Time和int类型，int类型表示时间戳，通常使用gorm.Model中的定义      | autoCreateTime;autoUpdateTime:milli;autoCreateTime:nano;分别表示秒级，毫秒级，纳秒级                       |
| autoUpdateTime         | 创建/更新时追踪当前时间，支持Time和int类型，int类型表示时间戳，通常使用gorm.Model中的定义 | autoCreateTime;autoUpdateTime:milli;autoCreateTime:nano;分别表示秒级，毫秒级，纳秒级                       |
| autoIncrement          | autoUpdateTime自动增长                                                                    |                                                                                                            |
| autoIncrementIncrement | 自动增长的步长                                                                            |                                                                                                            |
| comment                | 注释                                                                                      |                                                                                                            |
|                        |                                                                                           |                                                                                                            |
| size                   | 列长度，通常在type中指定，例如 varchar(255)                                               |                                                                                                            |
| precision              | 精度，通常在type中指定，例如 decimal(10, 2) 中的10                                        |                                                                                                            |
| scale                  | 小数位数，通常在type中指定，例如 decimal(10, 2)中的2                                      |                                                                                                            |

示例：

```go
type FieldTag struct {
	gorm.Model

	FTypeChar    string `gorm:"type:char(127)"`
	FTypeVarchar string `gorm:"type:varchar(255)"`
	FTypeText    string `gorm:"type:text"`
	FTypeBlob    []byte `gorm:"type:blob"`
	FTypeEnum    string `gorm:"type:enum('Go', 'GORM', 'MySQL')"`
	FTypeSet     string `gorm:"type:set('Go', 'GORM', 'MySQL')"`

    FColName    string `gorm:"column:custom_field"`
	FTypeNotNull string `gorm:"type:varchar(255);not null"`
	FTypeDefault string `gorm:"type:varchar(255);not null;default:hello gorm!"`
	FTypeComment string `gorm:"type:varchar(255);not null;default:hello gorm!;comment:some comment"`
}
```

通过migrate测试创建的表的字段属性。

```mysql
mysql> set names utf8;
Query OK, 0 rows affected, 1 warning (0.01 sec)

mysql> show create table msb_field_tag\G
*************************** 1. row ***************************
       Table: msb_field_tag
Create Table: CREATE TABLE `msb_field_tag` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `f_string_default` longtext,
  `f_type_char` char(32) DEFAULT NULL,
  `f_type_varchar` varchar(255) DEFAULT NULL,
  `f_type_text` text,
  `f_type_blob` blob,
  `f_type_enum` enum('Go','GORM','MySQL') DEFAULT NULL,
  `f_type_set` set('Go','GORM','MySQL') DEFAULT NULL,
  `custom_column_name` longtext,
  `f_col_not_null` varchar(255) NOT NULL,
  `f_col_default` varchar(255) NOT NULL DEFAULT 'gorm middle ware',
  `f_col_comment` varchar(255) DEFAULT NULL COMMENT '带有注释的字段',
  PRIMARY KEY (`id`),
  KEY `idx_msb_field_tag_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)
```

注意，类型与结构之间的合理性，下面的标签也可以设置成功，但操作时会有很大局限性：

```go
 FTypeUn1 string `gorm:"type:int"`
 FTypeUn2 int    `gorm:"type:tinyint"`
```

## 索引和约束

- 索引，Index，快速检索数据
  - Index
  - UniqueIndex
- 约束，Constraint，表中数据的限制条件，大部分约束都是基于索引实现的。
  - 索引约束, Index
  - 主键约束，PK
  - 外键约束，FK
  - 检查约束，check
  - NOT NULL
  - DEFAULT

MySQL中使用 Key 实现了约束和索引的功能：

索引和约束是通过字段标签来定义的，常用的标签如下：

| key         | 功能     | 说明 |
| ----------- | -------- | ---- |
| primaryKey  | 主键     |      |
| unique      | 唯一键   |      |
| index       | 索引     |      |
| uniqueIndex | 唯一索引 |      |
| check       | 检查约束 |      |

支持创建复合索引，通过名字识别。复合索引中，默认基于模型字段顺序确定索引字段优先级，支持使用priority选项定义。

支持索引选项，index:索引名字,key1:value1,key2:value2 的方式指定选项：

- sort:desc 降序关键字，默认升序
- length:N 前缀N作为关键字
- comment 索引注释
- type:btree 索引类型
- where:CONDITION 过滤条件

示例：

```go
type IAndC struct {
	ID          uint    `gorm:"primaryKey"`
	Email       string  `gorm:"type:varchar(255);unique"`
	Age         int8    `gorm:"index;check:age >= 18 AND email is not null"`
	FirstName   string  `gorm:"index:name"`
	LastName    string  `gorm:"index:name"`
	FirstName1  string  `gorm:"index:name1,priority:2"`
	LastName1   string  `gorm:"index:name1,priority:1"`
	Height      float32 `gorm:"index:,sort:desc"`
	AddressHash string  `gorm:"type:varchar(42);index:,length:12"`
	Telephone   string  `gorm:"type:varchar(16);uniqueIndex:,comment:电话号码唯一索引"`
}
```

利用Migrate测试表结构：

```mysql
mysql> show create table msb_i_and_c\G
*************************** 1. row ***************************
       Table: msb_i_and_c
Create Table: CREATE TABLE `msb_i_and_c` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `age` tinyint DEFAULT NULL,
  `first_name` varchar(191) DEFAULT NULL,
  `last_name` varchar(191) DEFAULT NULL,
  `first_name1` varchar(191) DEFAULT NULL,
  `last_name1` varchar(191) DEFAULT NULL,
  `address_hash` varchar(42) DEFAULT NULL,
  `height` float DEFAULT NULL,
  `telephone` varchar(16) DEFAULT NULL,

  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `telephone` (`telephone`),
  UNIQUE KEY `idx_msb_i_and_c_telephone` (`telephone`) COMMENT '电话号码唯一索引',
  KEY `idx_msb_i_and_c_address_hash` (`address_hash`(12)),
  KEY `idx_msb_i_and_c_age` (`age`),
  KEY `name` (`first_name`,`last_name`),
  KEY `name1` (`last_name1`,`first_name1`),
  KEY `idx_msb_i_and_c_height` (`height` DESC),
  CONSTRAINT `name_checker` CHECK ((`name` <> _utf8mb4'jinzhu'))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)
```

写数据时，就要满足以上的约束了。

测试不满足约束的插入：

```go
func IAndCCreate() {
	iac := &IAndC{}
	//iac.Age = 18
	if err := DB.Create(iac).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", iac)
}
```

测试：

```shell
func TestIAndCCreate(t *testing.T) {
	IAndCCreate()
}

> go test -run IAndCCreate
2023/04/06 16:17:22 Error 3819 (HY000): Check constraint 'chk_msb_i_and_c_age' is violated.
exit status 1   
FAIL    gormExample     0.063s
```

## 字段操作控制

### 忽略字段

`-` 标签标示忽略该字段的数据库操作。

例如，网络传递的字段是 url，但数据表中存储的字段是：schema, host, path, query_string。就是将URL拆解成多个部分进行存储，但网络传输时使用的是url字段。那么该url字段，就应该忽略数据库的读写（包括迁移）操作，示例：

```go
type Service struct {
	Schema      string
	Host        string
	Path        string
	QueryString string
	Url         string `gorm:"-"`
}
```

以上模型在迁移生成表时，会忽略Url字段，同样在CRUD时，也会被忽略。

迁移测试：

```mysql
mysql> desc msb_service;
+--------------+----------+------+-----+---------+-------+
| Field        | Type     | Null | Key | Default | Extra |
+--------------+----------+------+-----+---------+-------+
| schema       | longtext | YES  |     | NULL    |       |
| host         | longtext | YES  |     | NULL    |       |
| path         | longtext | YES  |     | NULL    |       |
| query_string | longtext | YES  |     | NULL    |       |
+--------------+----------+------+-----+---------+-------+
4 rows in set (0.01 sec)
```

忽略字段可以做更细致的忽略限定：

```
-:migration // 忽略迁移，不忽略CRUD
```

测试，手动增加url字段：

```mysql
mysql> alter table msb_service add column url longtext;
```

测试读写操作：

```go
func ServiceCRUD() {
	s := &Service{}
	s.Schema = "https"
	s.Url = "https://www.mashibing.com/study?key=value"
	DB.Create(s)
	fmt.Printf("%+v\n", s)
}
```

基于tag的不同：

```
// - 不能修改
// -:migration 可以修改
```

### 权限控制

使用tag，可以控制某个字段的读写权限：

```
// 写权限控制
<-:false 无写入
<-:create 仅创建
<-:update 仅更新

// 读权限控制
->:false 无读取
```

示例：

```go
type Service struct {
	Schema      string
	Host        string
	Path        string `gorm:"->:false"`
	QueryString string `gorm:"<-:update"`
	Url string `gorm:"-"`
}
```

测试CRUD，可以看到对应的操作会被控制。

## 字段自动编解码

当我们需要处理数据库不能直接处理的数据时，通常要自定义处理过程。典型的自定义方案有：

- 使用序列化器
- 使用实现Scanner和Valuer接口的自定义类型

### 序列化器 Serializer

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1680606521071/e89858df32a54e7f94e77a123066e58e.png)

优先推荐使用GORM提供的序列化器，完成字段的序列化与反序列化。

GORM 提供了一些默认的序列化器：json、gob、unixtime。

使用 serializer 标签进行设置。

示例，处理为JSON编码：

```go
type Paper struct {
	gorm.Model

	Subject string
	//Tags    []string
	// 使用 json 序列化器进行处理
	Tags []string `gorm:"serializer:json"`
}

func PaperCrud() {
	if err := DB.AutoMigrate(&Paper{}); err != nil {
		log.Fatal(err)
	}

	// 常规操作
	paper := &Paper{}
	paper.Subject = "使用Serializer操作Tags字段"
	paper.Tags = []string{"Go", "Serializer", "Gorm", "MySQL"}
	// create 会执行序列化工作,serialize
	if err := DB.Create(paper).Error; err != nil {
		log.Fatal(err)
	}

	// 查询
	newPaper := &Paper{}
	// First会执行反序列化工作，unserialize
	DB.First(newPaper, 5)
	fmt.Printf("%+v\n", newPaper)
}

```

测试，发现可以处理[]string类型的字段，对应的数据库内容是JSON编码内容：

```mysql
*************************** 1. row ***************************
        id: 1
created_at: 2023-04-07 16:09:12.993
updated_at: 2023-04-07 16:09:12.993
deleted_at: NULL
   subject: 使用Serializer操作Tags字段
categories: NULL
      tags: ["Go","Serializer","Gorm","MySQL"]
5 rows in set (0.00 sec)
```

而模型字段是[]类型：

```go
Tags:[Go Serializer Gorm MySQL]
```

### 自定义编解码器

编解码器需要实现下面两个接口：

```go
// import "gorm.io/gorm/schema"

type SerializerInterface interface {
    Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error
    SerializerValuerInterface
}

type SerializerValuerInterface interface {
    Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error)
}
```

以上接口的方法，与 Valuer和Scanner接口几乎一致：

```go
// database/sql/sql.go
// Scanner is an interface used by Scan.
type Scanner interface {
	// Scan 从数据库中分配一个值，用于查询时设置字段值
	Scan(src any) error
}

// database/sql/driver/types.go
type Valuer interface {
	// Value 用于获取模型字段的值
	Value() (Value, error)
}
```

可见，思路一致的。

自定义编码器步骤：

1. 定义实现编码器接口的类型
2. 注册编码器
3. 在模型tag中使用

示例，实现自定义编码器，同样处理[]string类型为CSV格式（Comma-Separated Values，逗号分隔值）：

```go
// 1.定义实现了序列化器接口的类型
type CSVSerializer struct{}

// 实现Scan，unserialize时执行
// ctx Context对象
// field 模型的字段对应的类型
// dst 目标值（最终结果赋值到dst）
// dbValue 从数据库读取的值
// 错误
func (CSVSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	// 初始化一个用来存储字段值的变量
	var fieldValue []string
	// 一:解析读取到的数据表的数据
	if dbValue != nil { // 不是 NULL
		// 支持解析的只有string和[]byte
		// 使用类型检测进行判定
		var str string
		switch v := dbValue.(type) {
		case string:
			str = v
		case []byte:
			str = string(v)
		default:
			return fmt.Errorf("failed to unmarshal CSV value: %#v", dbValue)
		}
		// 二：核心：将数据表中的字段使用逗号分割，形成 []string
		fieldValue = strings.Split(str, ",")
	}

	// 三，将处理好的数据，设置到dst上
	field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(fieldValue))
	return nil
}

// 实现Value, serialize时执行
// fieldValue 模型的的字段值
func (CSVSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	// 将字段值转换为可存储的CSV结构
	return strings.Join(fieldValue.([]string), ","), nil
}
```

测试，csv编码器的使用：

```go
type Paper struct {
	gorm.Model

	Subject string
	//Tags    []string
	// 使用 json 序列化器进行处理
	Tags []string `gorm:"serializer:json"`
	// 使用自定义的编码器
	Categories []string `gorm:"serializer:csv"`
}

// 2.注册到GORM中
// 3.测试
func CustomSerializer() {
	// 注册序列化器
	schema.RegisterSerializer("csv", CSVSerializer{})

	// 测试
	if err := DB.AutoMigrate(&Paper{}); err != nil {
		log.Fatal(err)
	}

	// 常规操作
	paper := &Paper{}
	paper.Subject = "使用自定义的Serializer操作Categories字段"
	paper.Tags = []string{"Go", "Serializer", "Gorm", "MySQL"}
	paper.Categories = []string{"Go", "Serializer", "Gorm", "MySQL"}
	// create 会执行序列化工作,serialize
	if err := DB.Create(paper).Error; err != nil {
		log.Fatal(err)
	}

	// 查询
	newPaper := &Paper{}
	// First会执行反序列化工作，unserialize
	DB.First(newPaper, paper.ID)
	fmt.Printf("%+v\n", newPaper)
}
```

unit test:

```shell
func TestCustomSerializer(t *testing.T) {
	CustomSerializer()
}

> go test -run CustomSerializer
&{Model:{ID:16 CreatedAt:2023-04-07 19:01:53.513 +0800 CST UpdatedAt:2023-04-07 19:01:53.513 +0800 CST DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Subject:使用自定义的Serializer操作Categories字段 Tags:[Go Serializer Gorm MySQL] Categories:[Go Serializer Gorm MySQL]}
PASS
ok      gormExample     0.082s


# mysql
*************************** 16. row ***************************
        id: 16
created_at: 2023-04-07 19:01:53.513
updated_at: 2023-04-07 19:01:53.513
deleted_at: NULL
   subject: 使用自定义的Serializer操作Categories字段
categories: Go,Serializer,Gorm,MySQL
      tags: ["Go","Serializer","Gorm","MySQL"]
16 rows in set (0.00 sec)
```


## 嵌入结构体和gorm.Model

### gorm.Model

GORM 定义一个 `gorm.Model` 结构体，其包括字段 `ID`、`CreatedAt`、`UpdatedAt`、`DeletedAt`：

```go
type Model struct {
    // Primary Key
	ID        uint `gorm:"primarykey"`
    // 创建时间
	CreatedAt time.Time
    // 创建或更新时间
	UpdatedAt time.Time
    // 删除时间
	DeletedAt DeletedAt `gorm:"index"`
}
```

其中，DeleteAt类型的定义：

```go
type DeletedAt sql.NullTime
```

其中 sql.NullTime 类型指的是可以为Null的Time类型，定义如下：

```go
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}
```

若需要使用自定义名字的创建和更新时间，可以使用tag：autoCreateTime 和 autoUpdateTime。

强烈推荐，实体模型都要嵌入 gorm.Model。

### 嵌入结构体

除了gorm.Model, 其他结构体也可以嵌入。

结构体字段标签 embeddedPrefix 用于设置嵌入结构体字段映射到DB中的前缀。

示例：

```go
type Blog struct {
	gorm.Model

	BlogBasic `gorm:""`
	Author    `gorm:"embeddedPrefix:author_"`
}

type BlogBasic struct {
	Subject string
	Summary string
	Content string
}

type Author struct {
	Name  string
	Email string
}

```

迁移生成的表：

```mysql
mysql> desc msb_blog;
+---------------+-----------------+------+-----+---------+----------------+
| Field         | Type            | Null | Key | Default | Extra          |
+---------------+-----------------+------+-----+---------+----------------+
| id            | bigint unsigned | NO   | PRI | NULL    | auto_increment |
| created_at    | datetime(3)     | YES  |     | NULL    |                |
| updated_at    | datetime(3)     | YES  |     | NULL    |                |
| deleted_at    | datetime(3)     | YES  | MUL | NULL    |                |
| subject       | longtext        | YES  |     | NULL    |                |
| summary       | longtext        | YES  |     | NULL    |                |
| content       | longtext        | YES  |     | NULL    |                |
| author_name   | longtext        | YES  |     | NULL    |                |
| author_email  | longtext        | YES  |     | NULL    |                |
+---------------+-----------------+------+-----+---------+----------------+
```

### 实际开发时模型结构体通常有三个部分

```go
type Blog struct {
	// 一：基础结构
	gorm.Model

	// 二：实体字段
	BlogBasic `gorm:""`
	Author    `gorm:"embeddedPrefix:author_"`

	// 三：关联关系
	User user.User
}
```


## 小结

核心内容：

- 表名管理
- 类型管理
  - 默认的映射类型
  - 指针和非指针类型
  - 自定义字段类型
  - 字段的编码解码
- 字段属性管理
  - NOT NULL
  - DEFAULT
  - COMMENT
- 索引和约束的管理
  - 主键约束
  - 唯一键，唯一索引
  - 普通索引
  - 复合索引
  - check约束
- 嵌入结构体的使用
