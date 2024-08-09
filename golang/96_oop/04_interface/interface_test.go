package main

import (
	"fmt"
	"testing"
)

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
	//...
	subStruct    *MyStruct
	subInterface MyInterface
}

func (s *MyStruct) MyMethod() {
	//...
}

func TestInterfaceNil(t *testing.T) {

	myStruct := &MyStruct{}
	fmt.Println("myStruct == nil:", myStruct == nil)
	fmt.Println("myStruct.subStruct == nil:", myStruct.subStruct == nil)

	var subStructToInterface MyInterface = myStruct.subStruct
	fmt.Println("subStructToInterface == nil:", subStructToInterface == nil)
	fmt.Println("subStructToInterface == nil:", &subStructToInterface == nil)
	// 在Go语言中，即使接口底层的具体类型为nil，接口值本身也不会为nil。
	// 这是因为接口值包括两部分信息：底层具体类型的值和具体类型的类型信息。
	// 只有在两个部分都为nil的情况下，接口值才会被认为是nil
	// 若一个为nil的struct被指定给了它的interface，则直接通过 == nil 来判断的话，结果会为false
	fmt.Println("reflect.ValueOf(subStructToInterface).IsNil():", IsNil(subStructToInterface))

	var subInterface = myStruct.subInterface
	fmt.Println("subInterface == nil:", subInterface == nil)
	fmt.Println("subInterface == nil:", &subInterface == nil)

}
