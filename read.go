// 这个文件是做什么的。
package main

import (
	"fmt"
	"io/ioutil"
	"unsafe"
)

func main() {
	fmt.Println("please input path of *.txt")
	var path string
	fmt.Scanln(&path)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("err")
	}
	// fmt.Println(dat) // []byte
	fmt.Println(*(*string)(unsafe.Pointer(&dat)))
}
