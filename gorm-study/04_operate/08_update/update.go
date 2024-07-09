package _8_update

import (
	"fmt"
	gorm_study "gorm-study"
	"gorm.io/gorm"
	"log"
)

type Content gorm_study.Content

//### 主键更新
//
//模型更新的典型方法 `Save()`，用来存储模型的字段。会基于模型**主键**是否存在有效值（非零值）决定执行Insert或Update操作。
//
//示例：

func UpdatePK() {
	DB := gorm_study.DB
	c := Content{}
	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)

	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)
}

//### 条件更新及更新行数
//
//可以使用struct或map[string]any表示字段和值的对应，来更新满足条件的记录。通常使用：
//
//- Updates() ，执行update操作
//- Where(), 设置where条件子句
//- Model(), 设置from表名子句
//
//完成更新。
//
//示例：

func UpdateWhere() {
	DB := gorm_study.DB
	values := map[string]any{
		"subject": "new subject",
		"likes":   10001,
	}
	result := DB.
		Model(&Content{}).
		Where("likes = ?", 0).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 条件更新时，通常需要知道更新的记录数量，通过 result.RowsAffected 来获取
	log.Println("updated rows num: ", result.RowsAffected)
}

// ### 阻止无条件的更新
//
//若条件更新时未指定条件，那么GORM不会更新记录，同时会返回 `ErrMissingWhereClause`错误。
//
//此行为的目的，是为了保护失误情况下的全局更新。
//
//示例：

func UpdateNoWhere() {
	DB := gorm_study.DB
	// map结构
	values := map[string]any{
		"subject": "new subject",
		"likes":   10001,
	}
	result := DB.
		Model(&Content{}).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
		// WHERE conditions required
	}
	log.Println("updated rows num: ", result.RowsAffected)
}

//### 使用表达式值更新
//
//若带更新字段的值为表达式，例如+10。需要使用clause.Expr{}类型进行表示。推荐使用gorm.Expr()方法来构建该类型：
//
//示例：

func UpdateExpr() {
	DB := gorm_study.DB
	// 更新的字段值数据
	// map推荐
	values := map[string]any{
		"subject": "Where Update Row",
		// 值为表达式计算的结果时，使用Expr类型
		"likes": gorm.Expr("likes + ?", 10),
		//"likes": "likes + 10",
		// Incorrect integer value: 'likes + 10' for column 'likes' at row 1
	}

	// 执行带有条件的更新
	result := DB.Model(&Content{}).
		Where("likes > ?", 100).
		Updates(values)
	// [17.011ms] [rows:51] UPDATE `msb_content` SET `likes`=likes + 10,`subject`='Where Update Row',`updated_at`='2023-04-21 17:28:45.498' WHERE likes > 100 AND `msb_content`.`deleted_at` IS NULL

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	// 获取更新结果，更新的记录数量（受影响的记录数）
	// 指的是修改的记录数，而不是满足条件的记录数
	log.Println("updated rows num: ", result.RowsAffected)
}

// ### 更新Hook
//
//类似创建，更新支持4个Hook：
// 执行顺序
// 启动事务
//BeforeSave,
//BeforeUpdate,
//Updates()
//AfterUpdate
//AfterSave,
// 提交事务
//更新操作的Hook中，两个特殊的操作比较常用：
//
//- db.Statement.SetColumn，修改某个特定字段的值，用于before钩子方法。通常可以使用模型直接操作。
//- tx.Statement.Changed, 判定某些字段是否发生变化，用于before钩子方法，在使用update或updates方法时起作用。通过与model的字段值比较，判定是否变化。
