package main

import (
	"fmt"
)

// SelectionSort sorts an array using the selection sort algorithm
func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		// Find the minimum element in the unsorted part of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element of the unsorted part
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	array := []int{64, 25, 12, 22, 11}
	fmt.Println("Original array:", array)
	sortedArray := SelectionSort(array)
	fmt.Println("Sorted array:", sortedArray)
}
