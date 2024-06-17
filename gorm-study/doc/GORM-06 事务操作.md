# 事务操作

事务，Transaction，指的是一组数据库操作组成的执行单元，要不全部的操作都成功对数据库产生影响，要不全部的操作都不会对数据库产生影响。在数据库系统中用于保证数据的完整性和一致性。其典型特征为：

- 原子性（**A**tomicity）: 事务 `要么全部完成，要么全部取消`。 如果事务崩溃，状态回到事务之前（事务回滚）。
- 隔离性（**I**solation）: 如果2个事务 T1 和 T2 同时运行，事务 T1 和 T2 最终的结果是相同的，不管 T1和T2谁先结束。
- 持久性（**D**urability）: 一旦事务提交，不管发生什么（崩溃或者出错），数据要保存在数据库中。
- 一致性（**C**onsistency）: 只有合法的数据（依照关系约束和函数约束）才能写入数据库。

## 事务方法

Gorm支持如下方法关联事务：

```go
// 开始事务
tx := DB.Begin()
// 回滚事务
tx.Rollback()
// 提交事务
tx.Commit()
```

注意，DB.Begin()方法返回开始事务的数据库对象，后续的本事务操作应该基于该对象完成，包括数据的CRUD等。

示例：

```go
type Author struct {
	gorm.Model
	Name   string
	// 积分
	Points int
}
```

```go
func TXDemo() {
	// 初始化测试数据
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a1.Points = 1600
	a2.Points = 200
	if err := DB.Create([]*Author{&a1, &a2}).Error; err != nil {
		log.Fatalln(err)
	}

	// 事务操作
	// a1 赠送 a2 2000 积分
	p := 2000
	// 开始事务
	tx := DB.Begin()
	// 有时需要考虑数据库是否支持事务的情景
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	// 执行赠送操作
	a1.Points -= p
	a2.Points += p

	// 1执行SQL，可能导致的错误
	if err := tx.Save(&a1).Error; err != nil {
		tx.Rollback()
		return
	}

	if err := tx.Save(&a2).Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return
	}

	// 2业务逻辑可能导致的错误
	// 要求author的积分不能为负数
	if a1.Points < 0 || a2.Points < 0 {
		log.Println("a1.Points < 0 || a2.Points < 0")
		// 回滚事务
		if err := tx.Rollback().Error; err != nil {
			log.Fatalln(err)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		log.Fatalln(err)
	}

	// 决定回滚还是提交，集中的处理错误风格
	//if err1 != nil || err2 != nil {
	//	tx.Rollback()
	//} else {
	//	tx.Commit()
	//}
}

```

测试，本例中，会导致回滚。

```shell
> go test -run TXDemo
2023/04/27 11:10:10 a1.Points < 0 || a2.Points < 0
2023/04/27 11:10:10 sql: transaction has already been committed or rolled back
exit status 1
FAIL    gormExample     0.147s
```

## 回调函数方式操作事务

除了手动调用 .Begin() tx.Commit() tx.Rollback() 外，GORM提供了一个回调函数的方案来执行事务，自动的完成开启事务和提交或回滚事务：

```go
func (db *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) (err error)
```

将业务逻辑代码直接由 func(tx *DB) error 函数实现即可。

同时 Transaction 也会得到具体错误。

推荐优先使用该方法完成事务。原因有2：

1. 编码重点在业务逻辑，不用关心事务的处理
2. 支持嵌套事务

示例：

```go
func TXCallback() {
	// 初始化测试数据
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a1.Points = 1600
	a2.Points = 200
	if err := DB.Create([]*Author{&a1, &a2}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(a1.ID, a2.ID)

	// 实现事务
	if err := DB.Transaction(func(tx *gorm.DB) error {
		// a1 赠送 a2 2000 积分
		p := 200
		// 执行赠送操作
		a1.Points -= p
		a2.Points += p

		// 1执行SQL，可能导致的错误
		if err := tx.Save(&a1).Error; err != nil {
			return err
		}

		if err := tx.Save(&a2).Error; err != nil {
			return err
		}

		// 2业务逻辑可能导致的错误
		// 要求author的积分不能为负数
		if a1.Points < 0 || a2.Points < 0 {
			return errors.New("a1.Points < 0 || a2.Points < 0")
		}

		// nil 的返回，会导致事务提交
		return nil
	}); err != nil {
		// 返回错误，为了后续的业务逻辑处理
		// 为了通知我们，事务成功还是失败
		// 返回错误，不影响事务的提交和回滚
		log.Println(err)
	}
}
```

Transaction方法还支持嵌套调用，用于支持嵌套事务，示例：

嵌套事务，在实操中，主要用于实现，或的逻辑，例如：a1 转给 a2 2000 积分，若失败，可以a3转给a2。此时，仅需将a1转的事务回滚即可。

示例：

```go
func TXNested() {
	// 初始化测试数据
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2, a3 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a3.Name = "欧文"
	a1.Points = 1600
	a2.Points = 200
	a3.Points = 4000
	if err := DB.Create([]*Author{&a1, &a2, &a3}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(a1.ID, a2.ID, a3.ID)

	// 实现事务
	if err := DB.Transaction(func(tx *gorm.DB) error {
		// a1 赠送 a2 2000 积分
		p := 20000

		// 执行赠送操作

		// a2 多了积分
		a2.Points += p
		if err := tx.Save(&a2).Error; err != nil {
			return err
		}

		// a1 赠送，使用嵌套事务完成
		errA1 := tx.Transaction(func(tx *gorm.DB) error {
			a1.Points -= p
			// 1执行SQL，可能导致的错误
			if err := tx.Save(&a1).Error; err != nil {
				return err
			}
			if a1.Points < 0 {
				return errors.New("a1.Points < 0")
			}
			// 没有错误成功
			return nil
		})

		// a1 发送失败，才需要a3
		if errA1 != nil {
			// a3 赠送，使用嵌套事务完成
			errA3 := DB.Transaction(func(tx *gorm.DB) error {
				a3.Points -= p
				if err := tx.Save(&a3).Error; err != nil {
					return err
				}
				if a3.Points < 0 {
					return errors.New("a3.Points < 0")
				}
				return nil
			})
			// a3 同样失败
			if errA3 != nil {
				return errors.New("a1 and a3 all send points failed")
			}
		}

		// nil 的返回，会导致事务提交
		return nil
	}); err != nil {
		// 返回错误，为了后续的业务逻辑处理
		// 为了通知我们，事务成功还是失败
		// 返回错误，不影响事务的提交和回滚
		log.Println(err)
	}
}
```

## SavePoint

GORM也提供了对事务逻辑存储点，及回到逻辑存储点的支持：

- SavePoint，定义SavePoint
- Rollbackto，回到SavePoint

示例，实现相同的a1 和 a3 给Points到a2的逻辑：

```go
func TXSavePoint() {
	// 初始化测试数据
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2, a3 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a3.Name = "欧文"
	a1.Points = 1600
	a2.Points = 200
	a3.Points = 4000
	if err := DB.Create([]*Author{&a1, &a2, &a3}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(a1.ID, a2.ID, a3.ID)

	// 事务操作
	// a1 赠送 a2 2000 积分
	p := 20000
	// 开始事务
	tx := DB.Begin()
	// 有时需要考虑数据库是否支持事务的情景
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	// 执行赠送操作
	// a2 得到积分
	a2.Points += p
	// 1执行SQL，可能导致的错误
	if err := tx.Save(&a2).Error; err != nil {
		tx.Rollback()
		return
	}

	// 逻辑记录发送points是否成功
	var flagSend bool

	// a1 先给 a2 send
	// 设置一个 savepoint
	tx.SavePoint("beforeA1")
	a1.Points -= p
	if err := tx.Save(&a1).Error; err != nil || a1.Points < 0 {
		// 回滚到 beforeA1
		tx.RollbackTo("beforeA1")

		// a3 to a2
		tx.SavePoint("beforeA3")
		a3.Points -= p
		if err := tx.Save(&a3).Error; err != nil || a3.Points < 0 {
			// 回滚到 beforeA3
			tx.RollbackTo("beforeA3")
		} else {
			flagSend = true
		}
	} else {
		flagSend = true
	}

	// 判定发送是否成功
	if flagSend {
		// 提交事务
		if err := tx.Commit().Error; err != nil {
			log.Fatalln(err)
		}
	} else {
		// 回滚事务
		tx.Rollback()
	}
}
```


## 禁用默认事务

了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。

如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升。

官网例子：

```go
// 全局禁用
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})

// 持续会话模式
tx := db.Session(&Session{SkipDefaultTransaction: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```
