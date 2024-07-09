package _7_select

import (
	"database/sql"
	"fmt"
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// // 查询一条
// db.First(&model, PK)
// // 模型的主键字段存在值，自动构建基于主键的查询
// model := Model{ID:10}
// db.First(&model)
//
// // 查询多条
// db.Find(&[]model, []PK{PK1, PK2, ...})

// 若主键为string类型，需要使用条件表达式：
// // 查询一条
// db.First(&model, "pk = ?", "stringPK")
//
// // 查询多条
// db.Find(&[]model, "pk IN ?", []PK{"stringPK1", "stringPK2", ...})

// 示例代码：

type ContentStrPK struct {
	ID          string `gorm:"primaryKey"`
	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

func GetByPk() {
	DB := gorm_study.DB
	DB.AutoMigrate(&gorm_study.Content{}, &ContentStrPK{})

	c := gorm_study.Content{}
	if err := DB.First(&c, 10).Error; err != nil {
		log.Println(err)
	}

	cStrPk := ContentStrPK{}
	if err := DB.First(&cStrPk, "id=?", "some id").Error; err != nil {
		log.Println(err)
	}

	var cs []gorm_study.Content
	if err := DB.Find(&cs, []uint{10, 11, 12}).Error; err != nil {
		log.Println(err)
	}

	var cStrPks []ContentStrPK
	if err := DB.Find(&cStrPks, "id IN ?", []string{"some", "id"}).Error; err != nil {
		log.Println(err)
	}
}

// 查询单条可以使用以上三个方法，区别为：
//
//- db.First，主键升序排序的第一条
//- db.Last，主键降序排序的第一条
//- db.Take，不拼凑排序子句的第一条，数据库的默认返回顺序
//- 带有Limit的Find，若Find的结果传递为单模型的引用，也可以查询单条记录。但一定要配合Limit使用，否者会出现查询到多条，然后从中选择一条的严重查询性能Bug。
//
// 可见，区别在于使用非主键查询时，有差异。若使用主键查询，结果永远是确定的那条记录。
// 示例代码：

func GetOne() {
	DB := gorm_study.DB
	c := gorm_study.Content{}
	if err := DB.First(&c, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}

	o := gorm_study.Content{}
	if err := DB.Last(&o, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}

	n := gorm_study.Content{}
	if err := DB.Take(&n, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}

	f := gorm_study.Content{}
	if err := DB.Limit(1).Find(&f, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}

	fs := gorm_study.Content{}
	if err := DB.Find(&fs, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
}

// ### 扫描查询结果至Map映射表
//	Gorm除了允许将结果扫描至model或[]model外，还支持将结果扫描至map:
// // 单条
//	map[string]any
// // 多条
//	[]map[string]any
// 示例代码：

type Content gorm_study.Content

func GetToMap() {
	DB := gorm_study.DB
	// 单条
	c := map[string]any{}
	if err := DB.Model(&Content{}).First(&c, 1).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(c)
	// 需要接口类型断言，才能继续处理
	if c["ID"].(uint) == 13 {
		fmt.Println("id bingo")
	}
	// time类型的处理
	fmt.Println(c["CreatedAt"])
	t, err := time.Parse("2006-01-02 15:04:05.000 -0700 CST", "2023-04-10 22:00:11.582 +0800 CST")
	if err != nil {
		log.Println(err)
	}
	if c["CreatedAt"].(time.Time) == t {
		fmt.Println("created_at bingo")
	}

	// 多条
	var cs []map[string]any
	if err := DB.Model(&Content{}).Find(&cs, []uint{1, 2, 3}).Error; err != nil {
		log.Println(err)
	}
	for _, c := range cs {
		fmt.Println(c["ID"].(uint), c["Subject"].(string), c["CreatedAt"].(time.Time))
	}
}

// ###查询单列Pluck()
// 除了查询记录，还可以查询单列，使用方法DB.Pluck()实现。
// 需要使用Model确定映射的表名。
// 查询的结果是切片类型。
// 示例：

func GetPluck() {
	DB := gorm_study.DB
	var subjects []sql.NullString
	if err := DB.Model(&Content{}).Where("id > ?", 30).Pluck("subject", &subjects).Error; err != nil {
		log.Println(err)
	}

	for _, subject := range subjects {
		if subject.Valid {
			fmt.Println(subject.String)
		} else {
			fmt.Println("NULL")
		}
	}
}

// ### select字段选择子句
//	查询时，使用Select()方法指定需要从数据库查询的字段，默认为*全部字段
//	示例：

func GetSelect() {
	DB := gorm_study.DB
	var c Content
	if err := DB.Select("subject", "likes").First(&c, 1).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)

	// 字段还可以使用表达式代替
	if err := DB.Select("subject", "concat(subject, views) as sv").First(&c, 1).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)
}

// ### distinct去重子句
//	查询时去掉重复的行，主要配合Find()使用
// 示例：

func GetDistinct() {
	DB := gorm_study.DB
	var cs []Content

	if err := DB.Distinct("*").Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

func WhereMethod(likeCondition int, subjectCondition string) {
	var cs []Content
	DB := gorm_study.DB

	// 使用 Find 方法中condition字段
	// SELECT * FROM `content` WHERE likes > 100 AND subject like 'gorm%' AND `content`.`deleted_at` IS NULL
	if err := DB.Find(&cs, "likes > ? and subject like ?", 100, "gorm%").Error; err != nil {
		log.Fatalln(err)
	}

	// 通过where来根据字段是否为空来动态的控制条件查询
	// SELECT * FROM `content` WHERE likes > 100 AND subject like 'gorm%' AND `content`.`deleted_at` IS NULL

	var query1 = DB
	if likeCondition != 0 {
		query1 = DB.Where("likes > ?", 100)
	}

	if subjectCondition != "" {
		query1.Where("subject like ?", "gorm%")
	}
	if err := query1.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

	// 如何在where中使用 OR 条件
	// SELECT * FROM `content` WHERE (likes > 100 OR subject like 'gorm%') AND `content`.`deleted_at` IS NULL
	query2 := DB.Where("likes > ?", 100)
	query2.Or("subject like ?", "gorm%")
	if err := query2.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

	// 如何在where中使用 NOT 条件
	// SELECT * FROM `content` WHERE NOT likes > 100 AND subject like 'gorm%' AND `content`.`deleted_at` IS NULL
	query3 := DB.Not("likes > ?", 100)
	query3.Where("subject like ?", "gorm%")
	if err := query3.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

// #### 条件表述语法
//
// 使用如下的类型来表述方式条件：
//
//- **字符串，配合占位符（匿名和具名占位符）构建条件，最典型的结构，推荐。**
//- *gorm.DB类型，分组条件，用于构建复杂的逻辑运算。应该从初始的DB对象进行构建。
//- number，主键匹配
//- slice, In条件的值，未指定字段，则使用主键
//- map，key为字段，value为字段值，通常是=，in运算
//- Struct，field为字段，value为字段值，为=运算，零值不视作条件
//示例：

func WhereType() {
	DB := gorm_study.DB
	var cs []Content
	//
	//query := DB.Where("likes > ? and subject like ?", 100, "gorm%")
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	// 通过 *gorm.DB类型，分组条件，用于构建复杂的逻辑运算
	// 例如 where (条件1 or 条件2) and (条件3 and (条件4 or 条件5))
	// cond1 := DB.Where("likes > ?", 100).Or("likes < ?", 1000)
	// cond2 := DB.Where("views > ?", 2000).Or("views < ?", 20000)
	// cond3 := DB.Where("subject like ?", "gorm%").Where(cond2)
	// query := DB.Where(cond1).Where(cond3)
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	// 通过slice构建条件, 其中In条件的值，未指定字段，则使用主键
	//query := DB.Where("likes = ? AND views IN ?", 100, []uint{1, 2, 3})
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	// 通过map构建条件，其中key为字段，value为字段值，通常是=，in运算
	query := DB.Where(map[string]any{
		"likes": 100,
		"views": []uint{1, 2, 3},
	})
	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

	// 通过Struct构建条件，其中field为字段，value为字段值，为=运算，零值不视作条件
	//query := DB.Where(Content{
	//	Likes: 100,
	//	Views: 1000,
	//})
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}
}

// #### 占位符
//
//- ? 匿名占位符，通过索引顺序与数据绑定
//- @someName 具名占位符，通过名字与数据绑定
//
//示例：

func PlaceHolder() {
	var cs []Content
	DB := gorm_study.DB

	// 通过 sql.Named 将数据与具名占位符绑定
	query := DB.Where("Likes = @like AND Views IN @view", sql.Named("view", []int{1, 2}), sql.Named("like", 100))

	// 也可通过 map[string]any 将数据与具名占位符绑定
	//query := DB.Where("likes = @like AND views IN @view", map[string]any{
	//	"view": []uint{1, 2, 3},
	//	"name": 100,
	//})

	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

// ### order排序子句
//使用Db.Order()完成排序。
//常用的参数为字符串类型，设置order by子句。
//支持连续调用，设置多个排序字段。（多个排序字段拼凑成一个字符串也可以）
//还支持子句类型gorm.Clause.OrderBy{}类型，用于构建带有表达式的排序子句。
//示例，按照某个值list进行排序：

func OrderBy() {
	var cs []Content
	DB := gorm_study.DB

	query := DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "field(id, ?)",
			Vars:               []any{[]uint{2, 1, 3}},
			WithoutParentheses: true,
		},
	})
	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

//### limit结果限制子句
//Gorm使用：
//- Limit(n) 限定查询的结果数量，n若为-1，表示不使用限制子句。
//- Offset(n) 控制偏移，n若为-1，表示不使用偏移子句。
//
//示例（官网例子）：
// db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;
//
// db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;
//
// db.Limit(3).Offset(3).Find(&users)
// SELECT * FROM users LIMIT 3 OFFSET 3;
//典型的应用场景为分页：
//基于用户请求的页码数和每页需要的记录数量（这个也可以后端控制），来确定limit和offset的参数。
//示例：

// Pager 定义分页必要数据结构
type Pager struct {
	Page, PageSize int
}

// 默认的值
const (
	DefaultPage     = 1
	DefaultPageSize = 12
)

// Pagination 翻页程序
func Pagination(pager Pager) {
	DB := gorm_study.DB
	// 确定page, offset 和 pageSize
	page := DefaultPage
	if pager.Page != 0 {
		page = pager.Page
	}

	pageSize := DefaultPageSize
	if pager.PageSize != 0 {
		pageSize = pager.PageSize
	}

	// 计算offset
	// page, pageSize, offset
	// 1, 10, 0
	// 2, 10, 10
	// 3, 10, 20
	offset := pageSize * (page - 1)

	var cs []Content
	// SELECT * FROM `content` WHERE `content`.`deleted_at` IS NULL LIMIT 15 OFFSET 30
	if err := DB.Offset(offset).Limit(pageSize).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

//#### 使用Scope重用翻页代码
//由于分页是非常典型的查询列表操作。因此可以将翻页逻辑重用。
//使用GORM提供的Scopes，可以重用。
//作用域允许你复用通用的逻辑，这种共享逻辑需要定义为类型 `func(*gorm.DB) *gorm.DB`。
//需要通过DB.Scope(func(*gorm.DB) *gorm.DB)来复用。
//示例：

// 用于得到func(db *gorm.DB) *gorm.DB类型函数
// 为什么不直接定义函数，因为需要func(db *gorm.DB) *gorm.DB与分页信息产生联系。
func Paginate(pager Pager) func(db *gorm.DB) *gorm.DB {
	// 计算page
	page := DefaultPage
	if pager.Page != 0 {
		page = pager.Page
	}

	// 计算pagesize
	pagesize := DefaultPageSize
	if pager.PageSize != 0 {
		pagesize = pager.PageSize
	}

	// 计算offset
	// page, pagesize, offset
	// 1, 10, 0
	// 2, 10, 10
	// 3, 10, 20
	offset := pagesize * (page - 1)

	return func(db *gorm.DB) *gorm.DB {
		// 使用闭包的变量，实现翻页的业务逻辑
		return db.Offset(offset).Limit(pagesize)
	}
}

// PaginationScope 测试重用的分页查询
func PaginationScope(pager Pager) {
	DB := gorm_study.DB

	var cs []Content
	// SELECT * FROM `content` WHERE `content`.`deleted_at` IS NULL LIMIT 15 OFFSET 30
	if err := DB.Scopes(Paginate(pager)).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

	var users []gorm_study.User
	// SELECT * FROM `msb_post` WHERE `content`.`deleted_at` IS NULL LIMIT 15 OFFSET 30
	if err := DB.Scopes(Paginate(pager)).Find(&users).Error; err != nil {
		log.Fatalln(err)
	}
}

//### group和having分组和过滤子句
//
//- group 子句，将结果进行分组
//- Having子句，基于分组的结果进行过滤。通常分组后会执行合计操作，例如，count，max，avg等，基于这些操作的结果进行过滤。
//
//分组合计过滤后，得到的数据通常就不是典型的模型或模型集合了，而是自定义的结构体类型或map类型。因此在Find时，通常给的都是自定义的结构体切片。

func GroupHaving() {
	DB := gorm_study.DB
	DB.AutoMigrate(&Content{})

	type Result struct {
		UserID     uint
		TotalLikes int
		TotalViews int
		AvgViews   int
	}

	var rs []Result
	if err := DB.Model(&Content{}).Select("author_id", "SUM(likes) as total_likes", "SUM(views) as total_views", "AVG(views) as avg_views").
		Group("author_id").Having("total_views > ?", 99).
		Find(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	// SELECT `author_id`,SUM(likes) as total_likes,SUM(views) as total_views,AVG(views) as avg_views FROM `content` WHERE `content`.`deleted_at` IS NULL GROUP BY `author_id` HAVING total_views > 99

}

//### 锁子句
//
//GORM支持在查询时加锁，使用子句 clause.Locking实现：
//
//Locking结构如下：
//type Locking struct {
//    // 锁强度（类型），典型的SHARE，UPDATE
//	Strength string
//    // 对应的表
//	Table    Table
//    // 选项，例如NOWAIT非阻塞
//	Options  string
//}

// 示例：

func Locking() {
	DB := gorm_study.DB
	var cs []Content
	if err := DB.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&cs, "likes > ?", 10).Error; err != nil {
		log.Fatalln(err)
	}
	//[4.904ms] [rows:19] SELECT * FROM `content` WHERE likes > 10 AND `content`.`deleted_at` IS NULL FOR UPDATE

	if err := DB.Clauses(clause.Locking{Strength: "SHARE"}).Find(&cs, "likes > ?", 10).Error; err != nil {
		log.Fatalln(err)
	}
	// [2.663ms] [rows:19] SELECT * FROM `content` WHERE likes > 10 AND `content`.`deleted_at` IS NULL FOR SHARE
}

type Author gorm_study.Author

func SubQuery() {
	DB := gorm_study.DB
	DB.AutoMigrate(&Author{}, &Content{})

	authorIDs := DB.Model(&Author{}).Select("id").Where("status=?", 0)
	var cs []Content
	if err := DB.
		Where("author_id IN (?)", authorIDs).
		Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	// [2.128ms] [rows:0] SELECT * FROM `msb_content` WHERE author_id IN (SELECT `id` FROM `msb_author` WHERE status=0 AND `msb_author`.`deleted_at` IS NULL) AND `msb_content`.`deleted_at` IS NULL

	type Result struct {
		Subject string
		Likes   int
	}
	var rs []Result
	fromQuery := DB.Model(&Content{}).Select("subject", "likes").Where("publish_time is null")
	if err := DB.Table("(?) as temp", fromQuery).
		Where("likes > ?", 10).
		Find(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	// [3.278ms] [rows:17] SELECT * FROM (SELECT `subject`,`likes` FROM `msb_content` WHERE publish_time is null AND `msb_content`.`deleted_at` IS NULL) as temp WHERE likes > 10
}

//### 查询操作的钩子方法
//
//查询操作支持一个钩子方法： AfterFind(tx *gorm.DB) (err error)
//在查询后执行。先Find，在AfterFind，最后返回。

func (c *Content) AfterFind(db *gorm.DB) error {
	// 业务代码
	if c.AuthorID == 0 {
		c.AuthorID = 1 // 1 is default author
	}

	return nil
}

func FindHook() {
	DB := gorm_study.DB
	var c Content
	if err := DB.First(&c, 1).Error; err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c.AuthorID)
}
