package _5_association

import (
	_5_find_association "gorm-study/05_association/05_find_association"
	_6_save_association "gorm-study/05_association/06_save_association"
	"testing"
)

func TestFind(t *testing.T) {
	_5_find_association.AssocFind()
}

func TestSave(t *testing.T) {
	_6_save_association.AssocSave()
}
