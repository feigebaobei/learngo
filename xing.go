package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(sli []int) {
	for {
		passRequiredModification := false
		for index, _ := range sli {
			if index == len(sli)-1 {
				break
			}

			if sli[index] > sli[index+1] {
				Swap(sli, index)
				passRequiredModification = true
			}
		}

		if passRequiredModification == false {
			return
		}
	}
}

func Swap(sli []int, index int) {
	first := sli[index]
	second := sli[index+1]

	sli[index] = second
	sli[index+1] = first
}

func getEntries() []int {
	entries := make([]int, 0, 1)

	for {
		if len(entries) == 10 {
			return entries
		}

		fmt.Println("Give me an integer (or x to exit)")
		var currentInput string
		fmt.Scan(&currentInput)

		if currentInput == "x" {
			return entries
		}

		currentInputAsInt, err := strconv.Atoi(currentInput)

		if err != nil {
			fmt.Println("I said an integer! That's not an integer!")
			continue
		}

		entries = append(entries, currentInputAsInt)
	}
}

func main() {
	entries := getEntries()

	BubbleSort(entries)
	fmt.Print("Sorted: ")
	fmt.Print(entries)
}
