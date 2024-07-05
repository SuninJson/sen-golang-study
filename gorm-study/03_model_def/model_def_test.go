package _3_model_def

import (
	"errors"
	"fmt"
	gorm_study "gorm-study"
	"gorm-study/03_model_def/02_field_type"
	"gorm-study/03_model_def/03_field_tag"
	"gorm-study/03_model_def/04_index_constraint"
	"gorm-study/03_model_def/05_control_field"
	studySerializer "gorm-study/03_model_def/06_serializer"
	"gorm.io/gorm"
	"log"
	"reflect"
	"testing"
)

func TestCustomType(t *testing.T) {
	_2_field_type.CustomType()
}

func TestCreateFieldTagTable(t *testing.T) {
	_3_field_tag.CreateFieldTagTable()
}

func TestCreateIAndCTable(t *testing.T) {
	_4_index_constraint.CreateIAndCTable()
}

func TestCreateServiceTable(t *testing.T) {
	_5_control_field.CreateServiceTable()
}

func TestServiceCRUD(t *testing.T) {
	_5_control_field.ServiceCRUD()
}

func TestSerializerCurd(t *testing.T) {
	studySerializer.SerializerCurd()
}

func TestCSVSerializer(t *testing.T) {
	testTableCRUD(&studySerializer.TestCSVSerializerStruct{})
}

func testTableCRUD(tableStruct interface{}) {
	if err := gorm_study.DB.AutoMigrate(&tableStruct); err != nil {
		log.Fatalln(err)
	}

	// 常规操作
	tableStructType := reflect.TypeOf(tableStruct)
	newTestSerializer := new(tableStructType)
	// DB.First方法会执行反序列化工作
	if result := gorm_study.DB.First(newTestSerializer, "id = ?", "1"); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到匹配的记录
			fmt.Println("未找到匹配的记录，进行创建")
		}

		tableStructReflect := reflect.ValueOf(tableStruct).Elem()
		testField := tableStructReflect.FieldByName("Test")
		setValueForField(testField)
		testSerializer := tableStruct
		// create 会执行序列化工作,serialize
		if err := gorm_study.DB.Create(testSerializer).Error; err != nil {
			log.Fatal(err)
		}

		// 未找到匹配的记录，创建后，再次查询
		gorm_study.DB.First(newTestSerializer, "id = ?", "1")
	}

	fmt.Printf("%+v\n", newTestSerializer)
}

func setValueForField(field reflect.Value) {
	if field.IsValid() && field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.String {
		// 创建一个字符串切片并赋值给字段
		value := reflect.ValueOf([]string{"value1", "value2", "value3"})
		field.Set(value)
	} else {
		fmt.Println("Field is not a []string")
	}

}
