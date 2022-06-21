package main

import "fmt"

//严格意义上来说 go并不是面向对象语言
//但它也可以通过一些方法来实现面向对象

type person struct {
	name string
	age  int
}

//如果一个struct嵌套了另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type student struct {
	person //匿名字段，struct
	mobile string
}

//如果一个struct嵌套了另一个【有名】的结构体，叫做组合
type teacher struct {
	p      person //有名字段，struct
	mobile string
}

// 接收器 用来给struct添加方法
// Golang 中指针接收器与非指针接收器的唯一区别就是指针接收器可以对成员变量进行修改，而非指针接收器不可以对指针变量进行修改
func (p *person) run() {
	fmt.Println(p.name, " run")
}

func (p *person) reading() {
	fmt.Println(p.name, " reading")
}

func (s *student) reading() {
	fmt.Println(s.name, " reading")
}

func main() {
	p := person{"zhangsan", 22}
	s := student{person{"lisi", 20}, "000"}
	t := teacher{person{"wangwu", 25}, "000"}

	fmt.Println(s.name)   //访问【匿名】结构体的字段
	fmt.Println(t.p.name) //访问【有名】结构体的字段。不是继承，需要指定结构体
	p.run()
	s.run()
}
