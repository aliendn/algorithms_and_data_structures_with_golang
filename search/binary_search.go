package main

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
