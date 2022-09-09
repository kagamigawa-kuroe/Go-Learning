package main

import (
	"fmt"
)

func main() {

	// var/const + 变量名 + 类型 + 等于号 + 变量值
	// 分别表示变量和常量
	var name1 string = "Lucas"
	const name2 string = "Tom"

	// 可以省略变量的类型
	name3 := "Chloe"
	// 这样写等价于 var name3 = "Chloe"
	// 但要注意:=不能在函数外使用

	// var n = "123" ok
	// n string = "123" false
	// 说明:=的结构本质是省略了var的申明

	const name4 = "Jerry"

	// 可以一次性给多个变量赋值
	const name5, name6 string = "kitty", "Jay"
	const name7, number1 = "kimi", 3

	var (
		a1 = 1
		b1 = 2
		c1 = 3
	)

	/// 不申明初始值时 int = 0 string = "" bool = false

	//其他数据类型
	const number2 int = 10
	const number3 int8 = 20
	const number4 int16 = 20
	//int/unit 8 16 32 64

	//浮点数
	const float1 float32 = 1.887
	const float2 float64 = 2.00923

	//布尔
	/// 只有true和false 不能用0/1
	const bool1 bool = true

	fmt.Print(name1 + " " + name2 + "\n")
	fmt.Print(name3 + " " + name4 + "\n")
	fmt.Print(name5 + " " + name6 + "\n")
	fmt.Println(a1, b1, c1)
	fmt.Print("hello world\n")

	//println 输出并换行
	fmt.Println(bool1)

	//printf 与c语言的printf类似
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}
