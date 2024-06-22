package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FixedCapacityStackOfStrings struct {
	s []string
	N int
}

func newFixedCapacityStackOfStrings(capacity int) *FixedCapacityStackOfStrings {
	return &FixedCapacityStackOfStrings{
		s: make([]string, capacity),
		N: 0,
	}
}

func (stack *FixedCapacityStackOfStrings) isEmpty() bool {
	return stack.N == 0
}

func (stack *FixedCapacityStackOfStrings) Push(item string) {
	stack.s[stack.N] = item
	stack.N++
}

func (stack *FixedCapacityStackOfStrings) Pop() (string, error) {
	if stack.isEmpty() {
		return "", fmt.Errorf("pop from empty stack")
	}
	stack.N--
	item := stack.s[stack.N]
	stack.s[stack.N] = ""
	return item, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Give me Capacity ;D")
	scanner.Scan()
	capacity, _ := strconv.Atoi(scanner.Text())
	stack := newFixedCapacityStackOfStrings(capacity)
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
