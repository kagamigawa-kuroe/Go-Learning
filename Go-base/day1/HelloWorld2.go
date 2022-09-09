package main

import (
	"fmt"
	"math/rand"
	"time"
)

/// 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的
/// 已导出意味着我们可以在import这个文件后去访问他
/// 也就是可以访问文件中的大写开头的变量
/// 小写的则不行

func test1(a int, b int) int {
	return a+b
}

/// 当需要返回多个值时 需要用括号将其扩起来
func test2(a int, b int) (int,int) {
	return a,b
}

/// 多个相同类型的参数可以升落
func test3(a,b int) (int,int) {
	return a,b
}

/// 返回值可以被命名 也需要在括号内
func test4(a,b int)(c int){
	c = a+b
	return
}

func main() {
	fmt.Println("test")
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(10))
}
