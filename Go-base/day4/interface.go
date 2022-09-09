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

	/// 这里的借口和java中的interface不太一样 更像是c++中实现多态的子类们的父类
	/// 接口实现多态的本质是 每个接口都会储存一个信息元组(value,type) 对应着实现接口类的变量值 以及 类型
	/// 因此接口也会用于做一些别的事情 比如空接口
	/// 空接口可以被赋予任何类型 也就是有点像java或早js中的object对象

	/// go中的断言
	/// t := i.(T) i是一个变量 我们这里断言i的类型是T 并将其(值)赋值给t
	/// 如果我们想要获取某个变量的类型 可以如下 t := i.(type) type是一个关键词 就会直接返回类型
	/// 断言失败则会产生异常
	/// 或者用如下方式去接异常
	/// t, ok := i.(T)
}
