package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (r *Person) to_string(){
	fmt.Println(r.name)
	fmt.Println(r.age)
}

func main() {
	var p1 Person
	p2 := Person{name: "hongzhe", age: 18}
	var p3 Person = Person{name: "hongzhe", age: 22}

	//这里我们可以总结出go语言变量赋值新的规律
	//var 变量名 变量类型 = 变量类型 变量值
	//对于默认的类型 字符串 整形等 右边变量类型不需要写
	// var a int = 8
	// 但对于数组 或者更高级的类型 以及自定义类型
	// 需要加上
	// var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(p1, p2, p3)

	//匿名结构体
	p4 := struct {
		name string
		age  int
	}{name: "123", age: 123}

	fmt.Println(p4)
	p2.to_string();
}
