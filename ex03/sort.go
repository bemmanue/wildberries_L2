package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	k *int
	n *bool
	r *bool
	u *bool
)

type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	if *r {
		return !less(h[i], h[j], *k, *n)
	}
	return less(h[i], h[j], *k, *n)
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

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func getNum(s []rune) ([]rune, int) {
	var num []rune
	var i int

	for i = 0; i < len(s) && isDigit(s[i]); i++ {
		num = append(num, s[i])
	}

	n, _ := strconv.Atoi(string(num))

	return s[i:], n
}

func less(str1, str2 string, n int, num bool) bool {
	s1 := skipWords(str1, n)
	s2 := skipWords(str2, n)

	var n1, n2 int
	if num {
		s1, n1 = getNum(s1)
		s2, n2 = getNum(s2)

		if n1 < n2 {
			return true
		} else if n1 > n2 {
			return false
		}
	}

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}

	if len(s1) < len(s2) {
		return true
	} else if len(s1) > len(s2) {
		return false
	}

	if n > 1 {
		return less(str1, str2, 1, false)
	}

	return true
}

func skipWords(s string, n int) []rune {
	res := []rune(s)
	i := 0

	for n > 1 {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		for i < len(s) && s[i] != ' ' {
			i++
		}
		if i < len(s) {
			i++
		}
		n--
	}

	return res[i:]
}

func main() {
	k = flag.Int("k", 1, "define a restricted sort key")
	n = flag.Bool("n", false, "sort fields numerically by arithmetic value")
	r = flag.Bool("r", false, "sort in reverse order")
	u = flag.Bool("u", false, "save only unique keys")
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
}
