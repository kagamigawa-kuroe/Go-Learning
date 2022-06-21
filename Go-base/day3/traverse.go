package main

import "fmt"

func main() {
	//类似c的遍历方式
	a := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
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
}
