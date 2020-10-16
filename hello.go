package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, world!")
	// three()
	// fourth()
  exerSlice()
}

func three() {
  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)
}
func fourth() {
  x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default:
      fmt.Printf("4")
  }
}

func exerSlice() {
  a1 := [4]string{"a", "b", "c", "d"}
  sli1 := a1[1:3]
  // fmt.Printf(len(sli1), cap(sli1))
  // fmt.Println(len(sli1))
  fmt.Println(len(sli1), cap(sli1))
  fmt.Println(sli1)
  sli1[1] = "str"
  fmt.Println(sli1)
  // sli1.append("e")
  // sli1 = append(sli1, "e")
  append(sli1, "e")
  fmt.Println(a1)
}