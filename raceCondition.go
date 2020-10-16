package main

import "fmt"

var balance = 0
var sigovers = make(chan struct{}, 2)

func Deposit(amount int) {
    balance += amount
    fmt.Printf("balance = %d\n", balance)
    sigovers <- struct{}{}
}

func main() {
    for i := 0; i < 30; i++ {
        go Deposit(100)
    }
    // for i := 0; i < 2; i++ {
    //     <-sigovers
    // }
    fmt.Printf("balance = %d\n", balance)
}

// 在main goroutine里创建了30次goroutine。它们都修改balance的值。这就是发生race condition的地方。

// Create goroutine 30 times in main goroutine. They both modify the value of balance. This is where the race condition occurs.