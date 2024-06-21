package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WeightedQuickUnionUF struct {
	id   []int
	size []int
}

func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	uf := &WeightedQuickUnionUF{
		id:   make([]int, n),
		size: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *WeightedQuickUnionUF) root(i int) int {
	for i != uf.id[i] {
		i = uf.id[i]
	}
	return i
}

func (uf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf *WeightedQuickUnionUF) Union(p, q int) {
	rootP := uf.root(p)
	rootQ := uf.root(q)

	if rootP == rootQ {
		return
	}

	// Union by size: attach smaller tree under larger tree
	if uf.size[rootP] < uf.size[rootQ] {
		uf.id[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	} else {
		uf.id[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	}
}

func (uf *WeightedQuickUnionUF) CountComponents() int {
	roots := make(map[int]bool)
	for i := 0; i < len(uf.id); i++ {
		root := uf.root(i)
		roots[root] = true
	}
	return len(roots)
}

func (uf *WeightedQuickUnionUF) GetComponents() map[int][]int {
	components := make(map[int][]int)
	for i := 0; i < len(uf.id); i++ {
		root := uf.root(i)
		components[root] = append(components[root], i)
	}
	return components
}

func main() {
	count, data := readFile()
	uf := NewWeightedQuickUnionUF(count)

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
