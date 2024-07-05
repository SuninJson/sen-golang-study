package gorm_study

import (
	"gorm-study/03_model_def/01_table_name"
	"testing"
)

func TestBasicUse(t *testing.T) {
	BasicUse()
}

func TestCreate(t *testing.T) {
	Create()
}

func TestRetrieve(t *testing.T) {
	Retrieve(1)
}

func TestUpdate(t *testing.T) {
	Update()
}

func TestDelete(t *testing.T) {
	Delete()
}

func TestMigrate(t *testing.T) {
	_1_table_name.CreateBoxTable()
}
