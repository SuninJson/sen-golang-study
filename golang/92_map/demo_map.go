package main

import "fmt"

func main() {
	//map的特点：
	//（1）map集合在使用前一定要make
	//（2）map的key-value是无序的
	//（3）key是不可以重复的，如果遇到重复，后一个value会替换前一个value
	//（4）value可以重复的

	println("map的创建方式")
	//声明的基本语法：var 变量名 map[键的数据类型]值的数据类型
	var map1 map[int]string
	//只声明map内存是没有分配空间
	//必须通过make函数进行初始化，才会分配空间
	map1 = make(map[int]string, 10)

	map1[20095452] = "张三"
	map1[20095387] = "李四"
	map1[20097291] = "王五"
	fmt.Println(map1)

	map2 := make(map[int]string)
	map2[20095452] = "张三"
	map2[20095387] = "李四"
	fmt.Println(map2)

	map3 := map[int]string{20095452: "张三", 20098765: "李四"}
	fmt.Println(map3)

	println("\nmap的增加和更新操作")
	println("map[key]= value  ---> 如果key还没有，就是增加，如果key存在就是修改")
	fmt.Println("增加和更新前:", map1)
	map1[20097292] = "朱六"
	map1[20095387] = "张四"
	fmt.Println("增加和更新后:", map1)

	println("\nmap的删除操作")
	fmt.Println("删除前:", map1)
	delete(map1, 20095387)
	fmt.Println("删除后:", map1)

	println("\nmap的查找操作")
	value, flag := map1[20095452]
	fmt.Println("查到到的值", value)
	fmt.Println("是否查找到：", flag)

	println("\n遍历map")
	for k, v := range map1 {
		fmt.Printf("key:%v value:%v \n", k, v)
	}

	println("\n将map的值类型再定义为map")
	mapValueMap := make(map[string]map[int]string)
	mapValueMap["班级1"] = make(map[int]string, 3)
	mapValueMap["班级1"][20096677] = "露露"
	mapValueMap["班级1"][20098833] = "丽丽"
	mapValueMap["班级1"][20097722] = "菲菲"
	mapValueMap["班级2"] = make(map[int]string, 3)
	mapValueMap["班级2"][20089911] = "小明"
	mapValueMap["班级2"][20085533] = "小龙"
	mapValueMap["班级2"][20087244] = "小飞"

	for k1, v1 := range mapValueMap {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学生学号：%v 学生姓名：%v \n", k2, v2)
		}
		fmt.Println()
	}
}
