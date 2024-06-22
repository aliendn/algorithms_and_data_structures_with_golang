package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type LinkedStackOfStrings struct {
	first *Node
}

type Node struct {
	item string
	next *Node
}

func NewNode(item string, next *Node) *Node {
	return &Node{item: item, next: next}
}

func NewLinkedStackOfStrings() *LinkedStackOfStrings {
	return &LinkedStackOfStrings{first: nil}
}

func (stack *LinkedStackOfStrings) IsEmpty() bool {
	return stack.first == nil
}

func (stack *LinkedStackOfStrings) Push(item string) {
	oldFirst := stack.first
	stack.first = NewNode(item, oldFirst)
}

func (stack *LinkedStackOfStrings) Pop() (string, error) {
	if stack.IsEmpty() {
		return "", fmt.Errorf("pop from empty stack")
	}
	item := stack.first.item
	stack.first = stack.first.next
	return item, nil
}

func main() {
	stack := NewLinkedStackOfStrings()
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
