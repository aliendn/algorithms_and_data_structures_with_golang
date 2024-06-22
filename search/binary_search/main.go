package main

import (
	"fmt"
)

func binarySearch(arr []int, value int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == value {
		return mid
	} else if arr[mid] > value {
		return binarySearch(arr, value, start, mid-1)
	} else {
		return binarySearch(arr, value, mid+1, end)
	}
}

func main() {
	arr := []int{0, 4, 8, 10, 25, 26, 49, 54, 67, 87, 93}
	var value int
	fmt.Print("What value do you want to search? ")
	fmt.Scanf("%d", &value)
	bs := binarySearch(arr, value, 0, len(arr)-1)
	if bs == -1 {
		fmt.Println("Value not found.")
	} else {
		fmt.Printf("Value found at index %d.\n", bs)
	}
}
