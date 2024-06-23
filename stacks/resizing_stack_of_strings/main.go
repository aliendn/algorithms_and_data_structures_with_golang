package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ResizingArrayStackOfStrings struct {
	s []string
	N int
}

func newResizingArrayStackOfStrings() *ResizingArrayStackOfStrings {
	return &ResizingArrayStackOfStrings{
		s: make([]string, 1),
		N: 0,
	}
}

func (stack *ResizingArrayStackOfStrings) isEmpty() bool {
	return stack.N == 0
}

func (stack *ResizingArrayStackOfStrings) Push(item string) {
	if stack.N == len(stack.s) {
		stack.resize(2 * len(stack.s))
	}
	stack.s[stack.N] = item
	stack.N++
}

func (stack *ResizingArrayStackOfStrings) Pop() (string, error) {
	if stack.isEmpty() {
		return "", fmt.Errorf("pop from empty stack")
	}
	stack.N--
	item := stack.s[stack.N]
	stack.s[stack.N] = ""
	if stack.N > 0 && stack.N == len(stack.s)/4 {
		stack.resize(len(stack.s) / 2)
	}
	return item, nil
}

func (stack *ResizingArrayStackOfStrings) resize(capacity int) {
	copy := make([]string, capacity)
	for i := 0; i < stack.N; i++ {
		copy[i] = stack.s[i]
	}
	stack.s = copy
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stack := newResizingArrayStackOfStrings()

	fmt.Println("Do you want to start your stack? (1, 0)")
	fmt.Print("If you insert 0, the program will stop: ")
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
