package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuickUnionUF struct {
	id []int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	qu := &QuickUnionUF{
		id: make([]int, n),
	}
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}
	return qu
}

func (qu *QuickUnionUF) root(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}
	return i
}

func (qu *QuickUnionUF) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *QuickUnionUF) Union(p, q int) {
	rootP := qu.root(p)
	rootQ := qu.root(q)
	qu.id[rootP] = rootQ
}

func (qu *QuickUnionUF) CountComponents() int {
	roots := make(map[int]bool)
	for i := 0; i < len(qu.id); i++ {
		root := qu.root(i)
		roots[root] = true
	}
	return len(roots)
}

func (qu *QuickUnionUF) GetComponents() map[int][]int {
	components := make(map[int][]int)
	for i := 0; i < len(qu.id); i++ {
		root := qu.root(i)
		components[root] = append(components[root], i)
	}
	return components
}

func main() {
	count, data := readFile()
	uf := NewQuickUnionUF(count)

	index := 1
	for index < len(data) {
		p := data[index]
		q := data[index+1]
		index += 2

		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Printf("%v %v\n", p, q)
		}
	}
	fmt.Printf("Number of connected components: %v\n", uf.CountComponents())
	fmt.Println("Connected components:")
	components := uf.GetComponents()
	for _, comp := range components {
		fmt.Println(comp)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() (int, []int) {
	file, err := os.Open("../tinyUF.txt")
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	var N int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.Fields(line)

		for _, str := range parts {
			num, err := strconv.Atoi(str)
			Check(err)
			data = append(data, num)
		}

		if N == 0 && len(parts) > 0 {
			N, err = strconv.Atoi(parts[0])
			Check(err)
		}
	}
	return N, data
}
