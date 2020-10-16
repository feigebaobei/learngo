// 这个文件是做什么的。
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("please input integer, one at a time.\n")
	var a int
	var sli = make([]int, 0)
	for {
		fmt.Scanln(&a)
		sli = append(sli, a)
		sort.Slice(sli, func(i, j int) bool {
			return sli[i] < sli[j]
		})
		fmt.Printf("%v\n", sli)
	}
}