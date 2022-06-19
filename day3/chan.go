package main

import "fmt"

//chan 类似于队列
//先进先出

func main() {
	fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()
	//goroutine 并发 又或者说是协程 （不是并行 或者线程）
	//来回切换进行任务
	//main 函数是一个主线程，是因为主线程执行太快了，子线程还没来得及执行，所以看不到输出。
	//time.Sleep(1 * time.Second) 加入这个后 就可以看到输出

	fmt.Println("main end")

	// chan 通道
	// 创建一个缓冲区为10的通道
	ch1 := make(chan string, 10)
	// 进入通道
	ch1 <- "abc"
	ch1 <- "efg"
	val := <-ch1
	fmt.Println(val)
}
