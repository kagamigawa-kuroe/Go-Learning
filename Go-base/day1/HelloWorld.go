// go语言既可以作为解释性语言 直接运行
// go run HelloWorld.go
// 也可以像C语言一样先编译 再运行
// go build HelloWorld.go
// ./HelloWorld

// 包的定义类似于java和python
// 每个go文件必须要属于一个包
// main包是一个特殊的包
// main包中应包括一个main
// 和其他语言类似，作为程序的入口

package main

/// fmt = format
import (
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")
}
