package main

import "fmt"

//go 中map数据类型为
//map[type of key]type of value
//比如 map[int]string
//不得不说真的很离谱

func main() {
	m1 := map[int]string{1: "123", 2: "456"}
	//或者使用make
	var m2 map[int]string = make(map[int]string)
	fmt.Println(m1, m2)
	fmt.Println(m1[2])
	fmt.Println(m1[1])
	m1[7] = "hello"

	//遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//单遍历key或者value
	// for _, v := range m1 {

	delete(m1, 1)
	fmt.Println(m1)
	for k, v := range m1 {
		fmt.Println(k, v)
	}

}
