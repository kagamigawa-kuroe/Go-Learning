package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

// 调用t1时候 返回值为1
// 我们可以看到 返回的是a
// return时 先给返回值赋值
// 返回值为1
// 然后再a++ 变成2
// 注意返回值和a是两个东西
// 第一步只是把a的值赋给了返回值
// 然后第二部为把a变成了2
// 但这并不代表返回值改变了
// 因为在第一步中返回值已经赋值了
// 最后才是将其返回
// 也就是返回回了1

func t1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

// 再来看这个f2函数
// 这个函数运行的返回结果为2
// 为什么呢
// 注意 t2的返回值
// 返回值为a这个变量
// 第一步 我们现将1赋给a
// 第二步 refer执行 a++
// 变为2
// 第三步 返回a
// 所以也就返回了2

// go语言 很奇妙吧.jpg

func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

func main() {
	fmt.Println(t1())
	fmt.Println(t2())

	var a1 = 1
	var b1 = 2
	defer fmt.Println(a1 + b1)
	a1 = 2
	//输出3

	var a2 = 1
	var b2 = 2
	defer func() {
		fmt.Println(a2 + b2)
	}()
	a2 = 2
	//输出4

	var a = 1
	var b = 2
	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
	a = 2
	//输出 3

	// 关于闭包
	// 我们用一个函数返回另一个函数
	x := test()
	fmt.Println(x()) //1
	fmt.Println(x()) //2

}

func test() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
