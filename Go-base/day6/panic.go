package main

import "fmt"

// panic and recover
func main(){
	defer func() {
		str := recover()
		fmt.Println(str)
    }()
	panic("panic!")
}