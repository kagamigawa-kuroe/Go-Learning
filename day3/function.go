package main

import "fmt"

// func function_name(input1 type1, input2 type2) (type1, type2) {
// 	// 函数体
// 	// 返回多个值
// 	return value1, value2
//  }

func add1(v1 int, v2 int) int {
	return v1 + v2
}

func add2(v1 int, v2 int) (int, int) {
	return v1 + v2, v2 - v1
}

func add3(v1 int, v2 int) (v3 int) {
	v4 := 8
	return v4
}

func main() {
	a, b := add2(2, 1)
	fmt.Println(a, b)
	fmt.Println(add2(2, 1))
	fmt.Println(add3(2, 1))

	// go默认的传参方式同c语言
	// 为值传递
	// 也就是说函数内形参的变化并不会影响上一级主函数内的值
	// 解决方法同c
	// 采用传递地址 或者说引用的方式

}
