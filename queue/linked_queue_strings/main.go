package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	item string
	next *Node
}

type LinkedQueueOfStrings struct {
	first *Node
	last  *Node
}

func newLinkedQueueOfStrings() *LinkedQueueOfStrings {
	return &LinkedQueueOfStrings{}
}

func (queue *LinkedQueueOfStrings) isEmpty() bool {
	return queue.first == nil
}

func (queue *LinkedQueueOfStrings) Push(item string) {
	oldLast := queue.last
	queue.last = &Node{item: item}
	if queue.isEmpty() {
		queue.first = queue.last
	} else {
		oldLast.next = queue.last
	}
}

func (queue *LinkedQueueOfStrings) Pop() (string, error) {
	if queue.isEmpty() {
		return "", fmt.Errorf("dequeue from empty queue")
	}
	item := queue.first.item
	queue.first = queue.first.next
	if queue.isEmpty() {
		queue.last = nil
	}
	return item, nil
}

func main() {
	stack := newLinkedQueueOfStrings()
	fmt.Println("Do you want to start your stack? (1, 0)")
	fmt.Print("If you insert 0, the program will stop: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	startPoint, _ := strconv.Atoi(scanner.Text())

	for startPoint == 1 {
		fmt.Print("Do you want to push your item? (1, 0): ")
		scanner.Scan()
		insert, _ := strconv.Atoi(scanner.Text())

		if insert == 1 {
			fmt.Print("Push: ")
			scanner.Scan()
			push := scanner.Text()
			stack.Push(push)
		}

		fmt.Print("Do you want to pull your item? (1, 0): ")
		scanner.Scan()
		remove, _ := strconv.Atoi(scanner.Text())

		if remove == 1 {
			item, err := stack.Pop()
			if err != nil {
				fmt.Println("There is no item. Please continue and add item to the stack!")
			} else {
				fmt.Printf("You removed: %s\n", item)
			}
		}

		fmt.Print("Continue (1, 0)?: ")
		scanner.Scan()
		startPoint, _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Thanks a lot!")
}
