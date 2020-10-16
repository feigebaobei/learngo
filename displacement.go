package main
import (
	"fmt"
	"math"
)
func main() {
	var a, v, s, t float64
	fmt.Println("please input acceleration")
	fmt.Scanln(&a)
	fmt.Println("please input initial velocity")
	fmt.Scanln(&v)
	fmt.Println("please input initial displacement")
	fmt.Scanln(&s)
	fmt.Println("please input time")
	fmt.Scanln(&t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a, vo, so float64) func (float64) float64 {
	return func (t float64) float64 {
		return a / 2 * math.Pow(t, 2) + vo * t + so
	}
}