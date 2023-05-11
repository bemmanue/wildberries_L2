package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
)

type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	return less(h[i], h[j])
}

func (h StringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StringHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func less(str1, str2 string) bool {
	s1 := []rune(str1)
	s2 := []rune(str2)

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}

	if len(s1) <= len(s2) {
		return true
	}

	return false
}

func main() {
	k := flag.Int("k", 1, "define a restricted sort key")
	n := flag.Bool("n", false, "sort fields numerically by arithmetic value")
	r := flag.Bool("r", false, "sort in reverse order")
	u := flag.Bool("u", false, "unique keys")
	flag.Parse()

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	h := &StringHeap{}
	heap.Init(h)

	newScanner := bufio.NewScanner(file)
	for newScanner.Scan() {
		heap.Push(h, newScanner.Text())
	}

	if err := newScanner.Err(); err != nil {
		log.Fatal(err)
	}

	for h.Len() > 0 {
		fmt.Printf("%s\n", heap.Pop(h))
	}

	fmt.Println(*k, *n, *r, *u)
}
