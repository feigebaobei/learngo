package main

import "fmt"

func main() {
	var y *int
	y = foo()
	fmt.Println("%d", *y)
	fmt.Println("%d", y)
}
func foo() *int {
	x := 1
	return &x
}
