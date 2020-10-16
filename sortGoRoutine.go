package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

func main() {
  fmt.Println("Please enter at least 4 digits.")
  inputReader := bufio.NewReader(os.Stdin)
  numsStr, errNumsStr := inputReader.ReadString('\n')
  if errNumsStr != nil {
    fmt.Println("illegal input")
  }
  numsStr = strings.Replace(numsStr, "\n", "", -1)
  sliInput := strings.Split(numsStr, " ")
  slires := make([]int, 0)
  for _, ele := range sliInput {
    if ele != "" {
      eleInt, _ := strconv.Atoi(ele)
      slires = append(slires, eleInt)
    }
  }

  step := len(slires) / 4
  slia := slires[0:step]
  slib := slires[step:step*2]
  slic := slires[step*2:step*3]
  slid := slires[step*3:]
  ca := make(chan []int)
  cb := make(chan []int)
  cc := make(chan []int)
  cd := make(chan []int)
  go orderPart(slia, ca)
  go orderPart(slib, cb)
  go orderPart(slic, cc)
  go orderPart(slid, cd)
  sliaRes := <- ca
  slibRes := <- cb
  slicRes := <- cc
  slidRes := <- cd

  res := mergeSlice(sliaRes, slibRes, slicRes, slidRes)
  fmt.Println(res)
}
func orderPart(sli []int, c chan []int) {
  sort.Ints(sli)
  c <- sli
}

func mergeSlicePart(slia, slib []int, c chan []int) {
  leni, lenj, i, j := len(slia), len(slib), 0, 0
  res := make([]int, 0)
  for ;i < leni && j < lenj; {
    if slia[i] < slib[j] {
      res = append(res, slia[i])
      i++
    } else {
      res = append(res, slib[j])
      j++
    }
  }
  if i < leni {
    res = append(res, slia[i:]...)
  }
  if j < lenj {
    res = append(res, slib[j:]...)
  }
  c <- res
}

func mergeSlice(s0, s1, s2, s3 []int) []int {
  c0 := make(chan []int)
  c1 := make(chan []int)
  c2 := make(chan []int)
  go mergeSlicePart(s0, s1, c0)
  go mergeSlicePart(s2, s3, c1)
  part0 := <- c0
  part1 := <- c1
  go mergeSlicePart(part0, part1, c2)
  part2 := <- c2
  return part2
}