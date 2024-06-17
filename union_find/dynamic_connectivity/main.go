package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p]) // Path compression
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP != rootQ {
		// Union by rank
		if uf.rank[rootP] > uf.rank[rootQ] {
			uf.parent[rootQ] = rootP
		} else if uf.rank[rootP] < uf.rank[rootQ] {
			uf.parent[rootP] = rootQ
		} else {
			uf.parent[rootQ] = rootP
			uf.rank[rootP]++
		}
		uf.count--
	}
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) CountComponents() int {
	return uf.count
}

func (uf *UnionFind) GetComponents() map[int][]int {
	components := make(map[int][]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		if _, ok := components[root]; !ok {
			components[root] = make([]int, 0)
		}
		components[root] = append(components[root], i)
	}
	return components
}

func main() {
	// Read the file
	file, err := os.Open("../tinyUF.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int
	var N int // Variable to hold the first integer read from the file

	// Process each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // Skip empty lines
		}

		// Split the line into parts based on whitespace
		parts := strings.Fields(line)

		// Convert each part to an integer
		for _, str := range parts {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			data = append(data, num)
		}

		// Store the first integer as N (if it's the first non-empty line)
		if N == 0 && len(parts) > 0 {
			N, err = strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Printf("this is data : %v\n", data)
	fmt.Printf("this is N : %v\n", N)
	uf := NewUnionFind(N)
	fmt.Printf("this is uf: %v\n", uf)
	// Simulating the while loop equivalent in Go
	index := 1
	for index < len(data) {
		p := data[index]
		q := data[index+1]
		index += 2

		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Printf("p: %v, q: %v\n", p, q)
		}
	}
	fmt.Printf("Number of connected components: %v\n", uf.CountComponents())
	fmt.Println("Connected components:")
	components := uf.GetComponents()
	for _, comp := range components {
		fmt.Printf("%v\n", comp)
	}
}
