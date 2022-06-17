package main

import "fmt"

func main() {
	//切片类型
	//可以自由添加和删除
	var slice1 []int = []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println(slice1, slice2) //[1 2 3] [1 2 3]

	//空切片
	slice3 := []int{}

	//make方法 构造切片
	//make方法也可以用于构造别的数据类型
	//第一个参数为构造的数据类型
	//第二个参数为初始大小
	//第三个参数为预留的空间大小

	slice4 := make([]int, 2, 4)

	fmt.Println(slice3, slice4) //[] [0 0]
	println(len(slice4))
	println(cap(slice4))

	//切片的操作
	//似乎和R语言中的数组操作很像
	//这么一看go还真是把各个语言都给缝完了
	//有点逆天
	fmt.Println(slice2[1])
	fmt.Println(slice2[:])
	fmt.Println(slice2[1:3])   //左开右闭
	fmt.Println(slice2[0:2:3]) //取0-1 再预留一个位置

	//切片追加
	slice5 := append(slice2, 10)
	slice6 := append(slice5, 22)
	fmt.Println(slice5, slice6)
}
