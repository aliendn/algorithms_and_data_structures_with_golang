package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ResizingArrayQueueOfStrings struct {
	q    []string
	head int
	tail int
	n    int
}

func newResizingArrayQueueOfStrings() *ResizingArrayQueueOfStrings {
	return &ResizingArrayQueueOfStrings{
		q:    make([]string, 2), // initial capacity of 2
		head: 0,
		tail: 0,
		n:    0, // number of elements in the queue
	}
}

func (queue *ResizingArrayQueueOfStrings) isEmpty() bool {
	return queue.n == 0
}

func (queue *ResizingArrayQueueOfStrings) size() int {
	return queue.n
}

func (queue *ResizingArrayQueueOfStrings) resize(capacity int) {
	if capacity < queue.n {
		panic("capacity less than current size")
	}
	copy := make([]string, capacity)
	for i := 0; i < queue.n; i++ {
		copy[i] = queue.q[(queue.head+i)%len(queue.q)]
	}
	queue.q = copy
	queue.head = 0
	queue.tail = queue.n
	fmt.Println(queue.size())
}

func (queue *ResizingArrayQueueOfStrings) Enqueue(item string) {
	if queue.n == len(queue.q) {
		queue.resize(2 * len(queue.q)) // double the array size if necessary
	}
	queue.q[queue.tail] = item
	queue.tail = (queue.tail + 1) % len(queue.q)
	queue.n++
	fmt.Println(queue.size())
}

func (queue *ResizingArrayQueueOfStrings) Dequeue() (string, error) {
	if queue.isEmpty() {
		return "", fmt.Errorf("dequeue from empty queue")
	}
	item := queue.q[queue.head]
	queue.q[queue.head] = ""
	queue.head = (queue.head + 1) % len(queue.q)
	queue.n--
	if queue.n > 0 && queue.n <= len(queue.q)/4 {
		queue.resize(len(queue.q) / 2) // shrink the array if necessary
	}
	fmt.Println(queue.size())
	return item, nil
}

func main() {
	queue := newResizingArrayQueueOfStrings()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Queue operations - Enqueue and Dequeue")
	fmt.Print("Enter 1 to start or 0 to exit: ")
	scanner.Scan()
	startPoint, err := strconv.Atoi(scanner.Text())
	if err != nil || (startPoint != 1 && startPoint != 0) {
		fmt.Println("Invalid input. Exiting.")
		return
	}

	for startPoint == 1 {
		fmt.Print("Do you want to enqueue an item? (1 for Yes, 0 for No): ")
		scanner.Scan()
		insert, err := strconv.Atoi(scanner.Text())
		if err != nil || (insert != 1 && insert != 0) {
			fmt.Println("Invalid input. Exiting.")
			return
		}

		if insert == 1 {
			fmt.Print("Enter the item to enqueue: ")
			scanner.Scan()
			item := scanner.Text()
			queue.Enqueue(item)
		}

		fmt.Print("Do you want to dequeue an item? (1 for Yes, 0 for No): ")
		scanner.Scan()
		remove, err := strconv.Atoi(scanner.Text())
		if err != nil || (remove != 1 && remove != 0) {
			fmt.Println("Invalid input. Exiting.")
			return
		}

		if remove == 1 {
			item, err := queue.Dequeue()
			if err != nil {
				fmt.Println("There is no item to dequeue.")
			} else {
				fmt.Printf("You dequeued: %s\n", item)
			}
		}

		fmt.Print("Continue? (1 for Yes, 0 for No): ")
		scanner.Scan()
		startPoint, err = strconv.Atoi(scanner.Text())
		if err != nil || (startPoint != 1 && startPoint != 0) {
			fmt.Println("Invalid input. Exiting.")
			return
		}
	}

	fmt.Println("Thanks a lot!")
}
