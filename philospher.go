package main

import (
	"fmt"
	"time"
	"sync"
)

// var queuef chan string // unbuffer channel 发送后接收前一起阻塞
// var goAheadf chan bool
var wg sync.WaitGroup
var queue = make(chan string, 2) // buffer channel buffer未满时不阻塞，满时阻塞。
var goAhead = make(chan bool, 2)
var cSticks [5]*CStick
var philos [5]*Philo

func host () {
	for _ = range queue {
		goAhead <- true
	}
}

type CStick struct {
	free chan bool
}

type Philo struct {
	leftCS, rightCS *CStick
	index int
}
func (p *Philo) eat(wg *sync.WaitGroup) {
	for i:=0; i<3; i++ {
		p.pickCS()
		p.pickCS()
		fmt.Printf("startint to eat %+v\n", p.index)
		fmt.Printf("finishing eating %+v\n", p.index)
		p.leftCS.free <- false
		p.rightCS.free <- false
		queue <- "string"
		<-goAhead
	}
	defer wg.Done()
}

func (p *Philo) pickCS () {
	select {
	case <- p.leftCS.free:
		// fmt.Printf("philos %d\n", p.index)
		time.Sleep(2*time.Second)
		return
	case <- p.rightCS.free:
		return
	}
}
func main () {
	for i:=0; i<5; i++ {
		cSticks[i] = &CStick{make(chan bool, 1)}
		cSticks[i].free <- false
	}
	for i:=0; i<5; i++ {
		philos[i] = &Philo{cSticks[i], cSticks[(i+1)%5], i}
	}

	wg.Add(5)
	go host()
	for i:=0; i<5; i++ {
		go philos[i].eat(&wg)
	}
	wg.Wait()
}


