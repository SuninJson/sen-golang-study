package main

import "fmt"

// Golang语言面向对象编程说明：
// （1）Golang也支持面向对象编程(OOP)，但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言。所以我们说Golang支持面向对象编程特性是比较准确的。
// （2）Golang没有类(class)，Golang的结构体(01_struct)和其它编程语言的类(class)有同等的地位，可以理解Golang是基于struct来实现OOP特性的。
// （3）Golang面向对象编程非常简洁，去掉了传统OOP语言的方法重载、构造函数和析构函数、隐藏的this指针等等
// （4）Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它OOP语言不一样，比如继承:Golang没有extends 关键字，继承是通过匿名字段来实现。

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	Age  int
}

func (student *Student) SetName(name string) {
	student.Name = name
}

func (student *Student) String() string {
	str := fmt.Sprintf("学生姓名 = %v ， 学生年龄 = %v", student.Name, student.Age)
	return str
}

func main() {
	println("结构体的定义和赋值")
	var person1 = Person{
		Name: "Jack",
		Age:  19,
	}
	fmt.Println("未赋值时默认值：", person1)
	fmt.Println(person1)

	println("\n创建结构体时赋值")
	person2 := Person{"Jason", 19}
	fmt.Println(person2)

	println("\n通过new内置函数创建结构体")
	// 通过new函数创建的结构体返回的是 *Type 即结构体的指针
	var person3 = new(Person)
	// 通过 * 来根据地址取值
	(*person3).Name = "Jackson"
	// Golang简化了赋值方式，下面语句Golang编译器对 person3.Age = 20 转换为了 (*person3).Age = 20
	person3.Age = 20
	fmt.Println(*person3)

	println("\n结构体的内存")
	fmt.Println("person3的内存地址：", &person3)
	fmt.Println("person3的Name的内存地址：", &person3.Name)
	fmt.Println("person3的Age的内存地址：", &person3.Age)

	println("\n结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类型) ")
	student1 := Student(person1)
	fmt.Println(student1)

	println("\n结构体的方法")
	fmt.Println("student1调用方法前：", student1)
	student1.SetName("Jack Up")
	fmt.Println("执行student1.SetName(\"Jack Up\")后：", student1)

	println("\n如果一个类型实现了String()这个方法，那么通过fmt.Println(&变量名)会调用这个变量的String()进行输出")
	fmt.Println(&student1)
}
