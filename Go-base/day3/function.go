package main

import "fmt"

// func function_name(input1 type1, input2 type2) (type1, type2) {
// 	// 函数体
// 	// 返回多个值
// 	return value1, value2
//  }

func test_mutiple_return() (int, int) {
	return 1, 2
}

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

func sum(args ...int) {
	total := 0
	for _, v := range args {
		total += v
	}
}

func plusN(n int) func(int) int {
	f := func(x int) int {
		return x + n
	}
	return f
}

func main() {
	a, b := add2(2, 1)
	fmt.Println(a, b)
	fmt.Println(add2(2, 1))
	fmt.Println(add3(2, 1))
	x, y := test_mutiple_return()
	fmt.Println(x, y)
	x2, _ := test_mutiple_return()
	fmt.Println(x2)
	// add(1,2,3,4);
	// go默认的传参方式同c语言
	// 为值传递
	// 也就是说函数内形参的变化并不会影响上一级主函数内的值
	// 解决方法同c
	// 采用传递地址 或者说引用的方式
	plus1 := plusN(1)
	fmt.Println(plus1(12))
	fmt.Println(plus1(12))
}
