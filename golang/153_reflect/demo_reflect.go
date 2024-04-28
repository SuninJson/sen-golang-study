package main

import (
	"fmt"
	"reflect"
)

func main() {
	println("\n 对基本数据类型反射")
	testReflect(10)

	println("\n 对结构体反射")
	student := Student{"Jack", 18}
	testReflect(student)
	testReflectStruct(student)
	testReflectStruct(&student)
}

func testReflect(i interface{}) {
	// Type和 Kind 的区别：
	// Type是类型, Kind是类别,Type和Kind 可能是相同的，也可能是不同的.
	// 比如:var num int = 10 num的Type是int , Kind也是int
	// 比如:var stu Student的 Type是 pkg1.Student , Kind是struct
	println("\n 通过反射获取数据类型")
	inParamType := reflect.TypeOf(i)
	fmt.Printf("传入参数的具体类型：%v，reflect.TypeOf返回的参数的类别：%v \n", inParamType, inParamType.Kind())

	println("\n 通过反射获取值")
	inParamValue := reflect.ValueOf(i)
	fmt.Printf("传入参数的值：%v，reflect.ValueOf返回的参数的类别：%v \n", inParamValue, inParamValue.Kind())

}

func testReflectStruct(i interface{}) {
	inParamValue := reflect.ValueOf(i)

	println("\n 获取字段的数量")
	isPtr := inParamValue.Kind() == reflect.Pointer
	var numField int
	if isPtr {
		numField = inParamValue.Elem().NumField()
	} else {
		numField = inParamValue.NumField()

	}
	fmt.Println(numField)

	println("\n 通过遍历获取具体的字段")
	for i := 0; i < numField; i++ {
		if isPtr {
			fmt.Printf("第%d个字段的值是：%v\n", i, inParamValue.Elem().Field(i))
		} else {
			fmt.Printf("第%d个字段的值是：%v\n", i, inParamValue.Field(i))
		}
	}
	println()

	if isPtr {
		println("\n 通过反射修改结构体中的变量")
		nameField := inParamValue.Elem().FieldByName("Name")
		if nameField.Kind() == reflect.String {
			fmt.Printf("%v的第一个字段，修改前的值：%v\n", inParamValue.Type(), nameField.String())
			nameField.SetString("张三")
			println("修改后的值：", nameField.String())
		}
	}

	println("\n 通过reflect.Value类型操作结构体内部的方法")
	numMethod := inParamValue.NumMethod()
	fmt.Println("传入参数的类型含有方法的数量：", numMethod)
	// 调用方法，方法的首字母必须大写才能有对应的反射的访问权限
	// 方法的顺序按照ASCII的顺序排列的，a,b,c...索引：0,1,2...
	println("\n 调用没有入参的方法")
	inParamValue.Method(0).Call(nil)

	println("\n 调用有入参的方法")
	var params []reflect.Value
	params = append(params, reflect.ValueOf("Jason"))
	params = append(params, reflect.ValueOf(20))
	result := inParamValue.Method(1).Call(params)
	fmt.Println("Student.BSet方法的返回值为：", result[0])
}

type Student struct {
	Name string
	Age  int
}

func (s Student) APrint() {
	println("Student.APrint")
	fmt.Printf("学生的名字：%v，年龄：%v\n", s.Name, s.Age)
}

func (s Student) BSet(name string, age int) Student {
	s.Name = name
	s.Age = age
	return s
}
