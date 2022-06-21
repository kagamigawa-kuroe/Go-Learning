package main

func test1(v1 int, v2 int) int {
	println(v1 + v2)
	return v1 + v2
}

func test2(v1 int, v2 int) int {
	println(v2 - v1)
	return v2 - v1
}

func main() {
	defer test1(1, 3)
	defer test2(1, 3)
	//先输出2 后输出4
	//defer将函数先注册
	//在结束时统一执行
	//后注册的先执行
	//FILO
}
