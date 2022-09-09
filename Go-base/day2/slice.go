package main

import "fmt"

func main() {
	//切片类型
	//可以自由添加和删除
	//和数组很像 中括号内没有任何内容的数组 就是切片
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
	// 和c++的vector有点类似
	slice4 := make([]int, 2, 4)

	fmt.Println(slice3, slice4) //[] [0 0]
	println("---------------")
	println(len(slice4))
	println(cap(slice4))
	slice4 = append(slice4, 100)
	fmt.Println(slice4)
	slice4 = append(slice4, 1000)
	fmt.Println(slice4)
	slice4 = append(slice4, 100)
	fmt.Println(slice4)
	println(len(slice4))
	println(cap(slice4))
	/// 5 8 空间不足会扩容

	slice4 = append(slice4, 100)
	fmt.Println(slice4)
	slice4 = append(slice4, 1000)
	fmt.Println(slice4)
	slice4 = append(slice4, 10000)
	fmt.Println(slice4)
	slice4 = append(slice4, 10000)
	fmt.Println(slice4)
	println(len(slice4))
	println(cap(slice4))
	/// 9 16 可以看出扩容大小是按照2次方进行的
	println("---------------")

	fmt.Println(slice4)
	slice4 = slice4[2:]
	fmt.Println(slice4)
	println(len(slice4))
	println(cap(slice4))
	// 7 14 
	// 这说明当对切片内容进行修改时 如果是删除了部分左边的内容
	// 长度和容量都会被修改

	// 而如果只是删去了右边的内容 则只会修改长度 容量不变

	//切片的操作
	//似乎和R语言中的数组操作很像
	fmt.Println(slice2[1])
	fmt.Println(slice2[:])
	fmt.Println(slice2[1:3])   //左开右闭
	fmt.Println(slice2[0:2:3]) //取0-1 再预留一个位置

	//切片追加
	slice5 := append(slice2, 10)
	slice6 := append(slice5, 22)
	fmt.Println(slice5, slice6)
	fmt.Println(slice5)
	/// 直接拷贝会是一个浅拷贝
	/// 深拷贝用copy函数
	/// copy(dest,source)
	sss := slice5
	sss[1] = 20
	fmt.Println(slice5, sss)
	fmt.Println(slice5[1:2], sss)
}
