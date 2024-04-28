package main

import (
	"fmt"
	"strconv"
)

// Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存。
// Golang中字符使用UTF-8编码
// 可从【http://www.mytju.com/classcode/tools/encode_utf8.asp 】查看UTF-8编码表
func main() {
	//定义字符类型的数据：
	var c1 byte = 'a'
	fmt.Println(c1) //97
	var c2 byte = '6'
	fmt.Println(c2) //54
	var c3 byte = '('
	fmt.Println(c3 + 20) //40
	// 字符类型，本质上就是一个整数，也可以直接参与运算，
	//输出字符的时候，会将对应的码值做一个输出
	//字母，数字，标点等字符，底层是按照ASCII进行存储。
	var c4 int = '中'
	fmt.Println(c4)
	//汉字字符，底层对应的是Unicode码值
	//对应的码值为20013，byte类型溢出，能存储的范围：可以用int
	//总结：Golang的字符对应的使用的是UTF-8编码（Unicode是对应的字符集，UTF-8是Unicode的其中的一种编码方案）
	var c5 byte = 'A'
	//想显示对应的字符，必须采用格式化输出
	fmt.Printf("c5对应的具体的字符为：%c", c5)

	//练习转义字符：
	//\n  换行
	fmt.Println("aaa\nbbb")
	//\b 退格
	fmt.Println("aaa\bbbb")
	//\r 光标回到本行的开头，后续输入就会替换原有的字符
	fmt.Println("aaaaa\rbbb")
	//\t 制表符
	fmt.Println("aaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("aaaaaaaa\tbbbbb")
	//\"
	fmt.Println("\"Golang\"")

	//字符串的使用：
	//1.定义一个字符串：
	var defS1 string = "你好全面拥抱Golang"
	fmt.Println("defS1:", defS1)
	//2.字符串是不可变的：指的是字符串一旦定义好，其中的字符的值不能改变
	var defS2 string = "abc"
	//s2 = "def"
	//s2[0] = 't'
	fmt.Println("defS2:", defS2)
	//3.字符串的表示形式：
	//（1）如果字符串中没有特殊字符，字符串的表示形式用双引号
	var defS3 string = "asdfasdfasdf"
	fmt.Println("defS3:", defS3)
	//（2）如果字符串中有特殊字符，字符串的表示形式用反引号 ``
	var defS4 string = `       
		package main        
		import "fmt"                
		func main(){                
			//测试布尔类型的数值：               
			var flag01 bool = true                
			fmt.Println(flag01)                        
			var flag02 bool = false                
			fmt.Println(flag02)                        
			var flag03 bool = 5 < 9                
			fmt.Println(flag03)        }        `
	fmt.Println("defS4:", defS4)
	//4.字符串的拼接效果：
	var defS5 string = "abc" + "def"
	defS5 += "hijk"
	fmt.Println("defS5:", defS5)
	//当一个字符串过长的时候：注意：'+' 需要保留在上一行的最后
	var defS6 string = "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def" + "abc" +
		"def" + "abc" + "def"
	fmt.Println("defS6:", defS6)

	//基本类型转string类型
	//方式1（推荐）:fmt.Sprintf("%参数",表达式)
	//	Sprint采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。
	//	如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1对应的类型是：%T ，s1 = %q \n", s1, s1)
	//'f':有小数部分但无指数部分，如123.456
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("s2对应的类型是：%T ，s2 = %q \n", s2, s2)
	//'t':单词true或false
	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3对应的类型是：%T ，s3 = %q \n", s3, s3)
	//'c':该值对应的unicode码值
	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s4对应的类型是：%T ，s4 = %q \n", s4, s4)
	//方式2:使用strconv包的函数
	var strconvN1 int = 18
	var strconvS1 string = strconv.FormatInt(int64(strconvN1), 10) //参数：第一个参数必须转为int64类型 ，第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s1对应的类型是：%T ，s1 = %q \n", strconvS1, strconvS1)
	var strconvN2 float64 = 4.29
	//第二个参数：'f':有小数部分但无指数部分，如123.456
	//第三个参数：9 保留小数点后面9位
	//第四个参数：表示这个小数是float64类型
	var strconvS2 string = strconv.FormatFloat(strconvN2, 'f', 9, 64)
	fmt.Printf("s2对应的类型是：%T ，s2 = %q \n", strconvS2, strconvS2)
	var strconvN3 bool = true
	var strconvS3 string = strconv.FormatBool(strconvN3)
	fmt.Printf("s3对应的类型是：%T ，s3 = %q \n", strconvS3, strconvS3)

	//string类型转基本类型

	//func ParseFloat(s string, bitSize int) (f float64, err error)
	//解析一个表示浮点数的字符串并返回其值。
	//如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
	//bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
	//返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
	var convS1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(convS1)
	fmt.Printf("b的类型是：%T,b=%v \n", b, b)

	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号。
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	var convS2 string = "19"
	var convN1 int64
	convN1, _ = strconv.ParseInt(convS2, 10, 64)
	fmt.Printf("convN1的类型是：%T,num1=%v \n", convN1, convN1)
}
