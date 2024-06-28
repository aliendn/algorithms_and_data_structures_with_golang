package main

import (
	"fmt"
)

// Function to Heapify a subtree rooted with node i which is an index in arr[].
func Heapify(arr []int, n int, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left = 2*i + 1
	right := 2*i + 2 // right = 2*i + 2

	// See if left child of root exists and is greater than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// See if right child of root exists and is greater than root
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Change root, if needed
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // swap

		// Heapify the root.
		Heapify(arr, n, largest)
	}
}

// Main function to do heap sort
func HeapSort(arr []int) {
	n := len(arr)

	// Build a maxheap.
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(arr, n, i)
	}

	// One by one extract elements
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // swap
		Heapify(arr, i, 0)
	}
}

func Partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		pi := Partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		leftHalf := make([]int, mid)
		rightHalf := make([]int, len(arr)-mid)

		copy(leftHalf, arr[:mid])
		copy(rightHalf, arr[mid:])

		MergeSort(leftHalf)
		MergeSort(rightHalf)

		i, j, k := 0, 0, 0

		for i < len(leftHalf) && j < len(rightHalf) {
			if leftHalf[i] < rightHalf[j] {
				arr[k] = leftHalf[i]
				i++
			} else {
				arr[k] = rightHalf[j]
				j++
			}
			k++
		}

		for i < len(leftHalf) {
			arr[k] = leftHalf[i]
			i++
			k++
		}

		for j < len(rightHalf) {
			arr[k] = rightHalf[j]
			j++
			k++
		}
	}
}

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

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

const MIN_RUN = 32

// Function to perform insertion sort on a subarray
func InsertionSortForTimSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Function to Merge two subarrays
func Merge(arr []int, left, mid, right int) {
	len1, len2 := mid-left+1, right-mid
	leftPart := make([]int, len1)
	rightPart := make([]int, len2)

	for i := 0; i < len1; i++ {
		leftPart[i] = arr[left+i]
	}
	for i := 0; i < len2; i++ {
		rightPart[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, left

	for i < len1 && j < len2 {
		if leftPart[i] <= rightPart[j] {
			arr[k] = leftPart[i]
			i++
		} else {
			arr[k] = rightPart[j]
			j++
		}
		k++
	}

	for i < len1 {
		arr[k] = leftPart[i]
		k++
		i++
	}

	for j < len2 {
		arr[k] = rightPart[j]
		k++
		j++
	}
}

// Function to perform Tim Sort
func TimSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i += MIN_RUN {
		InsertionSortForTimSort(arr, i, MinForTimSort(i+MIN_RUN-1, n-1))
	}

	size := MIN_RUN
	for size < n {
		for start := 0; start < n; start += size * 2 {
			mid := MinForTimSort(n-1, start+size-1)
			right := MinForTimSort(start+size*2-1, n-1)
			if mid < right {
				Merge(arr, start, mid, right)
			}
		}
		size *= 2
	}
}

func MinForTimSort(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	array := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", array)
	// sortedArray := SelectionSort(array)
	// sortedArray := BubbleSort(array)
	// sortedArray := InsertionSort(array)
	//QuickSort(array, 0, len(array)-1)
	// MergeSort(array)
	fmt.Println("Sorted array:", array)
}
