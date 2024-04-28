package main

import "fmt"

//在Golang中，如果一个struct嵌套了另一个匿名结构体，
//那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。

type Animal struct {
	Age    int
	Weight float32
}

func (an *Animal) Shout() {
	fmt.Println("喊叫")
}

func (an *Animal) ShowInfo() {
	fmt.Printf("动物的年龄是：%v,动物的体重是：%v\n", an.Age, an.Weight)
}

type Cat struct {
	//为了复用性，体现继承思维，嵌入匿名结构体：将Animal中的字段和方法都达到复用
	Animal
	//结构体的匿名字段可以是基本数据类型
	int
	//结构体的字段可以是结构体类型的。（组合模式）
	ear Ear
}

type Ear struct {
	Shape string
	Color string
}

// 对Cat绑定特有的方法：
func (c *Cat) catchMouse() {
	fmt.Println("抓老鼠")
}

func (c *Cat) Shout() {
	fmt.Println("喵喵叫")
}

func main() {
	cat := &Cat{}
	// cat.Age ---> cat对应的结构体中找是否有Age字段，如果有直接使用，如果没有就去找嵌入的结构体类型中的Age
	cat.Age = 3
	cat.Weight = 10.6

	// 如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
	cat.Animal.Shout()
	cat.Shout()

	cat.ShowInfo()
	cat.catchMouse()

	fmt.Println("结构体中类型为基本数据类型的匿名字段：", cat.int)

	// 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	cat2 := Cat{
		Animal{
			Weight: 9,
			Age:    8,
		},
		3,
		Ear{
			Shape: "尖尖的",
			Color: "白色",
		},
	}

	cat2.ShowInfo()
	fmt.Println(cat2.ear.Shape)

}
