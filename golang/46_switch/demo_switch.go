package main

import "fmt"

func main() {
	//给出一个学生分数：
	var score int = 88

	//switch后是一个表达式(即:常量值、变量、一个有返回值的函数等都可以)
	switch score / 10 {
	//case后面的值如果是常量值(字面量)，则要求不能重复
	//case后的各个值的数据类型，必须和 switch 的表达式数据类型一致
	//case后面可以带多个值，使用逗号间隔。比如 case 值1,值2...
	//不用像java一样每个分支后面都加入break关键字
	case 10, 9:
		fmt.Println("您的等级为A级")
	case 8:
		fmt.Println("您的等级为B级")
		//switch穿透，利用fallthrough关键字，如果在case语句块后增加fallthrough ,则会继续执行下一个case,也叫switch穿透
		fallthrough
	case 7:
		fmt.Println("您的等级为C级")
	case 6:
		fmt.Println("您的等级为D级")
	case 5:
		fmt.Println("您的等级为E级")
	case 4:
		fmt.Println("您的等级为E级")
	case 3:
		fmt.Println("您的等级为E级")
	case 2:
		fmt.Println("您的等级为E级")
	case 1:
		fmt.Println("您的等级为E级")
	case 0:
		fmt.Println("您的等级为E级")
	//default是用来“兜底”的一个分支，其它case分支都不走的情况下就会走default分支
	//default分支可以放在任意位置上，不一定非要放在最后。
	default:
		fmt.Println("您的成绩有误")
	}
}
