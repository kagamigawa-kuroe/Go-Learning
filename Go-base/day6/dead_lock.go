package main

import "fmt"
func main() {
    c := make(chan int)
    c <- 12 // write to a channel
    val := <-c // read from a channel
    fmt.Println(val)
}