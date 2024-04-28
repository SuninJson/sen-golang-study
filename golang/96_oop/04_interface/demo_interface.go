package main

import "fmt"

//（1）接口中可以定义一组方法，但不需要实现，不需要方法体。并且接口中不能包含任何变量。
//到某个自定义类型要使用的时候（实现接口的时候）,再根据具体情况把这些方法具体实现出来。
//（2）实现接口要实现所有的方法才是实现。
//（3）Golang中的接口不需要显式的实现接口。Golang中没有implement关键字。
//（Golang中实现接口是基于方法的，不是基于接口的）

// SayHello 接口的定义：定义规则、定义规范，定义某种能力：
type SayHello interface {
	//声明没有实现的方法
	sayHello()
}

// Eat 结构体可以实现多个接口
type Eat interface {
	eat()
}

type Chinese struct{}

// 实现接口的方法 ---> 具体的实现
func (person Chinese) sayHello() { fmt.Println("你好") }

func (person Chinese) eat() { fmt.Println("用筷子吃饭") }

type American struct{}

func (person American) sayHello() { fmt.Println("hi") }

func (person American) eat() { fmt.Println("eat") }

// 定义一个函数：专门用来各国人打招呼的函数，接收具备SayHello接口的能力的变量
func greet(s SayHello) { s.sayHello() }

func eat(e Eat) {
	e.eat()
}

type integer int

func (i integer) sayHello() {
	fmt.Println("Say hi ", i)
}

type CInterface interface{ c() }
type BInterface interface{ b() }

// AInterface 一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)，
// 这时如果要实现A接口,也必须将B,C接口的方法也全部实现。
// 否则在使用时会编译器会提示：
// Cannot use 'stu' (type Stu) as the type AInterface Type does not implement 'AInterface' as some methods are missing:b() c()
type AInterface interface {
	BInterface
	CInterface
	a()
}

type Stu struct{}

func (s Stu) a() { fmt.Println("a") }
func (s Stu) b() { fmt.Println("b") }
func (s Stu) c() { fmt.Println("c") }

// EmptyInterface 空接口没有任何方法,所以可以理解为所有类型都实现了空接口，
// 所以可以把任何一个变量赋给空接口
type EmptyInterface interface {
}

func main() {
	c := Chinese{}
	a := American{}
	println("\n在Golang中多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。 ")
	greet(c)
	greet(a)
	eat(c)
	eat(a)

	println("\n接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量")
	var s SayHello = c
	s.sayHello()

	println("\n只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。")
	var i integer = 10
	s = i
	s.sayHello()

	println("\n一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)")
	var stu Stu
	var aInterface AInterface = stu
	aInterface.a()
	aInterface.b()
	aInterface.c()

	println("\n空接口没有任何方法，所以可以理解为所有类型都实现了空接口，可以把任何一个变量赋给空接口")
	var e1 EmptyInterface = a
	fmt.Println(e1)
	var e2 EmptyInterface = c
	fmt.Println(e2)
	num := 6.6
	var e3 EmptyInterface = num
	fmt.Println(e3)

}
