package _1_create_association

import (
	gorm_study "gorm-study"
	"log"
)

type Author gorm_study.Author
type Essay gorm_study.Essay
type Tag gorm_study.Tag
type EssayMate gorm_study.EssayMate

func StdAssocModel() {
	DB := gorm_study.DB
	// 利用migrate创建表
	// 以及多对多的关联表
	// 以及外键约束
	if err := DB.AutoMigrate(&Author{}, &Essay{}, &Tag{}, &EssayMate{}); err != nil {
		log.Fatalln(err)
	}
	// CREATE TABLE `author` (
	//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	//  `created_at` datetime(3) DEFAULT NULL,
	//  `updated_at` datetime(3) DEFAULT NULL,
	//  `deleted_at` datetime(3) DEFAULT NULL,
	//  `status` bigint DEFAULT NULL,
	//  `name` longtext,
	//  `email` longtext,
	//  PRIMARY KEY (`id`),
	//  KEY `idx_author_deleted_at` (`deleted_at`)
	//) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

	// CREATE TABLE `essay` (
	//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	//  `created_at` datetime(3) DEFAULT NULL,
	//  `updated_at` datetime(3) DEFAULT NULL,
	//  `deleted_at` datetime(3) DEFAULT NULL,
	//  `subject` longtext,
	//  `content` longtext,
	//  `author_id` bigint unsigned DEFAULT NULL,
	//  PRIMARY KEY (`id`),
	//  KEY `idx_essay_deleted_at` (`deleted_at`),
	//  KEY `fk_author_essays` (`author_id`),
	//  CONSTRAINT `fk_author_essays` FOREIGN KEY (`author_id`) REFERENCES `author` (`id`)
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

	// CREATE TABLE `essay_mate` (
	//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	//  `created_at` datetime(3) DEFAULT NULL,
	//  `updated_at` datetime(3) DEFAULT NULL,
	//  `deleted_at` datetime(3) DEFAULT NULL,
	//  `keyword` longtext,
	//  `description` longtext,
	//  `essay_id` bigint unsigned DEFAULT NULL,
	//  PRIMARY KEY (`id`),
	//  KEY `idx_essay_mate_deleted_at` (`deleted_at`),
	//  KEY `fk_essay_essay_mate` (`essay_id`),
	//  CONSTRAINT `fk_essay_essay_mate` FOREIGN KEY (`essay_id`) REFERENCES `essay` (`id`)
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

	// CREATE TABLE `tag` (
	//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	//  `created_at` datetime(3) DEFAULT NULL,
	//  `updated_at` datetime(3) DEFAULT NULL,
	//  `deleted_at` datetime(3) DEFAULT NULL,
	//  `title` longtext,
	//  PRIMARY KEY (`id`),
	//  KEY `idx_tag_deleted_at` (`deleted_at`)
	//) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

	// CREATE TABLE `essay_tag` (
	//  `tag_id` bigint unsigned NOT NULL,
	//  `essay_id` bigint unsigned NOT NULL,
	//  PRIMARY KEY (`tag_id`,`essay_id`),
	//  KEY `fk_essay_tag_essay` (`essay_id`),
	//  CONSTRAINT `fk_essay_tag_essay` FOREIGN KEY (`essay_id`) REFERENCES `essay` (`id`),
	//  CONSTRAINT `fk_essay_tag_tag` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`)
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	log.Println("migrate successful")
}

//在GORM中，将模型关联分为四种：
//
//- Has Many，一对多关系中，一端有多个多端，外键定义在多端
//- Belongs To，一对多关系中，多端属于一端，外键定义在多端
//- Many to Many，多对多，外键定义在关联表中
//- Has One，一对多关系中，一端有一个多端，外键定义在多端
//
//注意Author和AuthorMate的关系定义：
//
//- 当前结构中，可以表示一对多，也可以表示一对一
//- 本例中，选择了一对一
//- 若需要一对多，那么增加Author中的关联定义 `AuthorMates []AuthorMate`

//操作关联时，使用方法
//
//```go
//db.Model(&model).Association("Association")
//```
//
//完成关联的建立。参数是模型中定义的关联字段，具体的关联类型取决于模型的定义。
//
//要求model的主键不能为空。
//
//关联建立后，即可完成关联的管理。
