package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QuickFindUF struct {
	id []int
}

func (quickfind *QuickFindUF) countComponents() int {
	unique := make(map[int]bool)

	// Insert elements into the map to ensure uniqueness
	for _, num := range quickfind.id {
		unique[num] = true
	}

	return len(unique)
}

func (quickfind *QuickFindUF) getComponents() map[int][]int {
	components := make(map[int][]int)
	for i := 0; i < len(quickfind.id); i++ {
		root := quickfind.id[i]
		if _, ok := components[root]; !ok {
			components[root] = make([]int, 0)
		}
		components[root] = append(components[root], i)
	}
	return components
}

func (quickfind *QuickFindUF) Connected(p int, q int) bool {
	return quickfind.id[p] == quickfind.id[q]
}

func (quickfind *QuickFindUF) Union(p int, q int) {
	pid := quickfind.id[p]
	qid := quickfind.id[q]

	for i := 0; i < len(quickfind.id); i++ {
		if quickfind.id[i] == pid {
			quickfind.id[i] = qid
		}
	}

}

func newQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &QuickFindUF{
		id: id,
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

func main() {
	count, data := readFile()
	uf := newQuickFindUF(count)
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
	fmt.Printf("Number of Connected components : %v\n", uf.countComponents())
	fmt.Println("Connected components: ")
	components := uf.getComponents()
	for _, comp := range components {
		fmt.Printf("%v\n", comp)
	}
}
