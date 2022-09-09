package main

import (
	"fmt"
)

func main() {
	//类似c的遍历方式
	a := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	// 类似while遍历方式

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//类似py和js的遍历方式
	for k, v := range a {
		fmt.Println(k, v)
		//k为下表
		//v为内容
	}

	//单独遍历内容
	for _, m := range a {
		fmt.Println(m)
	}

	//遍历切片 map同理

	if i > 5 {
		fmt.Println("i > 5")
	}

	/// switch 不需要break
	switch i {
	case 1:
		println(i)
	case 2:
		println(i + 1)
	default:
		println(i + 2)
	}

	//defer关键字
	defer fmt.Println("this is defer")
	fmt.Println("done")

	// done会在defer前输出 因为只有在main return时 defer才会被执行
}
