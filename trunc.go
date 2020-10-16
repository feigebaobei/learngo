// 这个文件是做什么的。
package main

import (
	"fmt"
)

func main() {
	fmt.Println("input a float number, please")
	// var f float64
	// fmt.Scan(&f)
	// fmt.Println(int(f))
	var (a, b, c, d int)
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}