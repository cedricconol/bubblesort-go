package main

import "fmt"

func main() {
	var numbers []int = []int{8, 22, 76, 33, 91}
	var ascending bool = true

	fmt.Println("Original list: ", numbers)
	bubblesort(numbers, ascending)
	fmt.Println("Sorted list: ", numbers)
}

func bubblesort(list []int, ascending bool) {
	var N int = len(list)
	var count int

	for i := 1; i < N; i++ {
		swapCount := pass(list, ascending)
		count++
		if swapCount == 1 {
			break
		}
	}

	fmt.Println("Number of iterations: ", count)

}

func swap(list []int, firstIndex int) {
	list[firstIndex], list[firstIndex+1] = list[firstIndex+1], list[firstIndex]
}

func pass(list []int, ascending bool) int {
	var N int = len(list)
	var swapCount int

	for i := 0; i < N-1; i++ {
		var firstElement int = list[i]
		var secondElement int = list[i+1]

		if ascending == true {
			if firstElement > secondElement {
				swap(list, i)
				swapCount++
			}
		} else {
			if firstElement < secondElement {
				swap(list, i)
				swapCount++
			}
		}

	}

	return swapCount
}
