package main

import (
	"fmt"
	"time"
)

func main(){
	c := time.Tick(1 * time.Second)
	tic := true
	for now := range c {
	if tic {
			fmt.Println("tic:", now)
		} else {
			fmt.Println("tac:", now)
		}
	tic = !tic
}
}