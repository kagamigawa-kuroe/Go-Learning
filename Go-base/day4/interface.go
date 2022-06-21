package main

import (
	"fmt"
	"strconv"
)

type person struct {
	age  int
	name string
}

type dog struct {
	age  int
	name string
}

//接口和struct的写法类似

type person_action interface {
	present()
}

func (d *dog) present() {
	println("wangwangwangwangwang" + d.name)
}

func (p *person) present() {
	s1 := strconv.Itoa(p.age)
	fmt.Println("my name is " + p.name + " my age is " + s1)
}

func main() {
	// 总结一下interface的用法
	// 先实现一个接口类
	// 在写出你想实现这个接口的具体类
	// 然后采用为类添加方法的形式
	// 将其添加到类中

	//接口类可以实现多态
	//方法为用接口对象去接受不同的类的引用
	//前提是不同的类已经用不同的方法实现了接口中的函数
	//这样当用接口对象调用方法时
	//会根据具体类的不同产生不同的结果
	var pa1 person_action = &person{18, "123"}
	var pa2 person_action = &dog{10, "ttt"}
	pa1.present()
	pa2.present()
}
