// 这个文件是做什么的。
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("input a string, please")
	var s string
	fmt.Scanln(&s)
	if strings.Contains(s, "i") {
		if strings.Contains(s, "a") {
			if strings.Contains(s, "n") {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}
}