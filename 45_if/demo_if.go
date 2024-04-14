package main

import "fmt"

func main() {
	//实现功能：根据给出的学生分数，判断学生的等级：
	//>=90  -----A
	//>=80  -----B
	//>=70  -----C
	//>=60  -----D
	//<60   -----E
	score := 80
	if score >= 90 {
		fmt.Println("您的成绩为A级别")
	} else if score >= 80 {
		fmt.Println("您的成绩为B级别")
	} else if score >= 70 {
		fmt.Println("您的成绩为C级别")
	} else if score >= 60 {
		fmt.Println("您的成绩为D级别")
	} else {
		fmt.Println("您的成绩为E级别")
	}
}
