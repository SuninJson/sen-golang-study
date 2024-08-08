package gorm_study

import (
	"gorm.io/gorm"
	"log"
)

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
