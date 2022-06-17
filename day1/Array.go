package main

import (
	"fmt"
)

func main() {
	//和c语言的数组一样
	//go语言的数组也是固定长度 无法改变的
	//一维数组
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	arr2 := [3]int{9, 8, 7}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)

	// arr4 := [...][...]int{{1,2,3,4},{5,6,7,8}}
	// 二维数组的列数必须是确定的值
	arr4 := [...][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(arr4)

	//同样长度的数组可以相互赋值
	arr5 := [...]int{1, 2, 3}
	arr5 = arr2
	fmt.Println(arr5)

	//c语言中很常见的形参问题
	var arr = [5]int{1, 2, 3, 4, 5}
	var ptr = &arr
	modifyArr1(arr)
	fmt.Println(arr)
	modifyArr2(&arr)
	fmt.Println(arr)
	fmt.Println(*ptr)
}

func modifyArr1(a [5]int) {
	a[1] = 20
}

func modifyArr2(a *[5]int) {
	a[1] = 20
}
