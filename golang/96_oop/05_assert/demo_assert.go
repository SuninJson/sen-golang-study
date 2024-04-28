package main

import "fmt"

// Go语言里面有一个语法，可以直接判断是否是该类型的变量：
// value, ok := element.(T)，
// 这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
func main() {
	c := Chinese{}
	a := American{}
	greet(c)
	greet(a)
}

// SayHello 接口的定义：定义规则、定义规范，定义某种能力
type SayHello interface{ sayHello() }
type Chinese struct{ name string }

func (person Chinese) sayHello() { fmt.Println("你好") }
func (person Chinese) eat()      { fmt.Println("用筷子吃饭") }

type American struct{ name string }

func (person American) sayHello() { fmt.Println("hi") }

func greet(s SayHello) {
	s.sayHello()
	// 通过断言判断接口是否转换成功，转换成功后再执行转换后的类型独有的方法
	ch, ok := s.(Chinese)
	if ok {
		ch.eat()
	}

}
