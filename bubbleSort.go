// 这个文件是做什么的。
package main

import (
	"fmt"
	"strconv"
)

func main() {
	sli := opSli()
	sli = BubbleSort(sli)
	fmt.Printf("%+v\n", sli)
}
func opSli() []int {
	var a string
	sli := make([]int, 0)
	for {
		fmt.Printf("please input a number.(or input x for exit.)")
		fmt.Scanln(&a)
		if a == "x" {
			return sli
		} else {
			n, err := strconv.Atoi(a)
			if err != nil {
				fmt.Printf("please input a number.(or input x for exit.)")
			} else {
				sli = append(sli, n)
			}
			if len(sli) == 10 {
				return sli
			}
		}
	}
}
func BubbleSort(sli []int) []int {
	for i:=0; i<len(sli); i++ {
		for j:=0; j<len(sli)-i-1; j++ {
			if sli[j]>sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	return sli
}
func Swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}