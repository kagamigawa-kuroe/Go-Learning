package main

import (
	"fmt"
	"time"
)

//用chan+select是一种很常用的监控协程的方法
//select和switch很类似
//只不过他监控的是chan的行为
// <-chan-name 代表从某个chan中读取的行为
// chan-name<- 代表从chan中写入的行为

//空的chan无法读取和写入 都会阻塞
//而在一个chan被关闭后 下一次读取他会读到0

//所以我们可以采用这种方式来监控一个携程
//当某个携程关闭了一个空chan时
//另一个携程再读取这个空chan就会读到0
//就意味着他监控的那个携程结束了

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop<- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}

//此外 for-range也是一种办法
//for-rang从channel上接收值，直到channel关闭，
//该结构在Go并发编程中很常用
//这对于从单一通道上获取数据去执行某些任务是十分方便的

// package main

// import (
//   "fmt"
//   "sync"
// )

// var wg sync.WaitGroup

// func worker(ch chan int) {
//   defer wg.Done()
//   for value := range ch {
//     fmt.Println(value) // do something
//   }
// }

// func main() {
//   ch := make(chan int)

//   wg.Add(1)
//   go worker(ch)

//   for i := 0; i < 3; i++ {
//     ch <- i
//   }

//   close(ch)
//   wg.Wait()