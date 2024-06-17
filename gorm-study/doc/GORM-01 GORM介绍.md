# GORM介绍

## ORM介绍

应用程序操作关系型数据系统的典型两种方案：

1. SQL，DAO，数据访问对象，通过编写SQL完成数据库的操作
2. ORM，对象关系映射，通过使用Object的语法完成关系型数据库的操作。其实底层也要转换为SQL。

其中，SQL的方案，我们通过 database/sql 提供的方法，可以完成操作。参考数据库操作课程。

ORM，Object Relational Mapping，对象关系映射。为了解决面向对象语言操作关系型数据库系统时，数据类型不匹配的一种技术。典型的语法特征是不用直接编辑SQL，直接通过对象的方法即可完成数据的典型操作。

典型的映射方案如下：

![image.png](https://fynotefile.oss-cn-zhangjiakou.aliyuncs.com/fynote/fyfile/13080/1680502356013/cb26f25dc2ee495dbfbe70d4bcc1341b.png)

| 对象                     | 关系           | ORM         |
| ------------------------ | -------------- | ----------- |
| 类，struct类型           | 表，table      | 模型，model |
| 对象，struct类型数据实例 | 记录，row      |             |
| 属性，struct实例的字段   | 字段，field    |             |
| 方法                     | 记录的CRUD操作 |             |
| 关系模型                 | 关联关系       |             |

## 常用的ORM

- GORM：https://github.com/go-gorm/gorm，32k
- XORM：https://github.com/go-xorm/xorm，6.6k

## GORM的特点

官网：https://gorm.io/

特点如下：

- GORM 官方支持的数据库类型有：MySQL, PostgreSQL, SQLite, SQL Server 和 TiDB
- 全功能 ORM
- 关联 (拥有一个，拥有多个，属于，多对多，多态，单表继承)
- Create，Save，Update，Delete，Find 中钩子方法
- 支持 Preload、Joins 的预加载
- 事务，嵌套事务，Save Point，Rollback To to Saved Point
- Context、预编译模式、DryRun 模式
- 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD
- SQL 构建器，Upsert，锁，Optimizer/Index/Comment Hint，命名参数，子查询
- 复合主键，索引，约束
- 自动迁移
- 自定义 Logger
- 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…
- 每个特性都经过了测试的重重考验
- 开发者友好
