package study_serializer

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
)

// 自定义编码器步骤：
//
// 1. 定义实现编码器接口的类型
// 编解码器需要实现两个接口 SerializerInterface和SerializerValuerInterface
// 2. 注册编码器
// 3. 在模型tag中使用

type TestCSVSerializerStruct struct {
	gorm.Model
	Test []string `gorm:"serializer:csv"`
}

type CSVSerializer struct {
}

// Scan 实现SerializerInterface.Scan方法
// field 模型的字段对应类型
// dst 目标值：最终的结果需要赋值到dst字段中
// dbValue 从数据库读取的值
// 返回 error：是否存在错误
func (CSVSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	// 初始化用来存储结果的变量
	var fieldValue []string
	if dbValue != nil {
		// 解析读取到的数据表的数据
		var str string
		switch v := dbValue.(type) {
		case string:
			str = v
		case []byte:
			str = string(v)
		default:
			// 支持解析的只有string和[]byte
			return fmt.Errorf("failed to unmarshar CSV value: %#v", dbValue)
		}

		// 将数据表中的字段使用逗号分割
		fieldValue = strings.Split(str, ",")
	}

	// 将处理好的数据，设置到dst上
	field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(fieldValue))
	return nil
}

func (CSVSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	// 将字段值转换为可存储的CSV结构
	return strings.Join(fieldValue.([]string), ","), nil
}

func CustomSerializer() {
	// 注册CSV序列化器
	schema.RegisterSerializer("csv", CSVSerializer{})
}
