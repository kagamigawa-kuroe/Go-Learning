package main

import "fmt"

func main() {

	// var/const + 变量名 + 类型 + 等于号 + 变量值
	// 分别表示变量和常量
	var name1 string = "Lucas"
	const name2 string = "Tom"

	// 可以省略变量的类型
	name3 := "Chloe"
	// 这样写等价于 var name3 = "Chloe"

	const name4 = "Jerry"

	// 可以一次性给多个变量赋值
	const name5, name6 string = "kitty", "Jay"
	const name7, number1 = "kimi", 3

	//其他数据类型
	const number2 int = 10
	const number3 int8 = 20
	const number4 int16 = 20
	//int/unit 8 16 32 64

	//浮点数
	const float1 float32 = 1.887
	const float2 float64 = 2.00923

	//布尔
	const bool1 bool = true

	fmt.Print(name1 + " " + name2 + "\n")
	fmt.Print(name3 + " " + name4 + "\n")
	fmt.Print(name5 + " " + name6 + "\n")
	fmt.Print("hello world\n")

	//println 输出并换行
	fmt.Println(bool1)

	//printf 与c语言的printf类似
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}
