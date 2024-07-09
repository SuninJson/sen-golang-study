package operate

import (
	finally_or_chain "gorm-study/04_operate/01_finally_or_chain"
	create "gorm-study/04_operate/02_create"
	update_insert "gorm-study/04_operate/03_update_insert"
	default_value "gorm-study/04_operate/04_default_value"
	_5_select_omit "gorm-study/04_operate/05_select_omit"
	_6_hook "gorm-study/04_operate/06_hook"
	_7_select "gorm-study/04_operate/07_select"
	"testing"
)

func TestOperatorType(t *testing.T) {
	finally_or_chain.OperatorType()
}

func TestCreateBasic(t *testing.T) {
	create.BasicCreate()
}

func TestUseMapCreate(t *testing.T) {
	create.UseMapCreate()
}

func TestMultiCreate(t *testing.T) {
	create.MultiCreate()
}

func TestUpSert(t *testing.T) {
	update_insert.UpSert()
}

func TestDefaultValue(t *testing.T) {
	default_value.DefaultValue()
	default_value.DefaultValueOften()
}

func TestSelectAndOmit(t *testing.T) {
	_5_select_omit.SelectCol()
}

func TestCreateUseHook(t *testing.T) {
	_6_hook.CreateUseHook()
}

func TestGetByPK(t *testing.T) {
	_7_select.GetByPk()
}

func TestGetOne(t *testing.T) {
	_7_select.GetOne()
}

func TestGetToMap(t *testing.T) {
	_7_select.GetToMap()
}

func TestGetPluck(t *testing.T) {
	_7_select.GetPluck()
}

func TestGetSelect(t *testing.T) {
	_7_select.GetSelect()
}

func TestGetDistinct(t *testing.T) {
	_7_select.GetDistinct()
}

func TestWhereMethod(t *testing.T) {
	_7_select.WhereMethod(0, "")
}

func TestWhereType(t *testing.T) {
	_7_select.WhereType()
}

func TestPlaceHolder(t *testing.T) {
	_7_select.PlaceHolder()
}

func TestOrderBy(t *testing.T) {
	_7_select.OrderBy()
}

func TestPagination(t *testing.T) {
	pager := _7_select.Pager{
		Page:     1,
		PageSize: 10,
	}
	_7_select.Pagination(pager)
}

func TestPaginationScope(t *testing.T) {
	request := _7_select.Pager{Page: 3, PageSize: 15}
	_7_select.PaginationScope(request)
}

func TestGroupHaving(t *testing.T) {
	_7_select.GroupHaving()
}

func TestLocking(t *testing.T) {
	_7_select.Locking()
}

func TestSubQuery(t *testing.T) {
	_7_select.SubQuery()
}

func TestFindHook(t *testing.T) {
	_7_select.FindHook()
}
