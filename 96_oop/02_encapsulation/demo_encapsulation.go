package main

import "fmt"

// 【1】什么是封装：
// 封装(02_encapsulation)就是把抽象出的字段和对字段的操作封装在一起，
// 数据被保护在内部,程序的其它包只有通过被授权的操作方法，才能对字段进行操作。
// 【2】封装的好处：
// 1) 隐藏实现细节
// 2) 提可以对数据进行验证，保证安全合理
func main() {
	p := NewPerson("Jack")
	p.SetAge(166)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
}

type person struct {
	Name string
	// age属性首字母小写，其它包不能直接访问
	age int
}

// NewPerson 定义工厂模式的函数，相当于构造器
func NewPerson(name string) *person {
	return &person{Name: name}
}

// SetAge 定义set和get方法，为age字段对外提供可操作的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("你传入的年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}
