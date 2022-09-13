package main

import "fmt"

type A struct{
	a int;
	b int;
}

type B struct{
	c int
}

type C struct{
	A
	B
}

func main(){
	var c1 C = C{A{1,2},B{3}}
	fmt.Println(c1.a)
	fmt.Println(c1.b)
}