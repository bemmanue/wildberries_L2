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

type Flag struct {
	key          int
	numeric      bool
	reverse      bool
	unique       bool
	month        bool
	ignoreBlanks bool
	check        bool
	humanNumeric bool
}

type Sort struct {
	keys map[string]bool
	flag Flag
	heap *StringHeap
}

func NewSort(flag Flag) *Sort {
	sort := &Sort{
		keys: make(map[string]bool),
		flag: flag,
		heap: &StringHeap{},
	}

	heap.Init(sort)

	return sort
}

func (s *Sort) Check(in *os.File) error {
	var current, temp string
	newScanner := bufio.NewScanner(in)

	newScanner.Scan()
	temp = newScanner.Text()

	if s.flag.reverse {
		for newScanner.Scan() {
			current = newScanner.Text()
			if less(temp, current, s.flag) {
				fmt.Fprintf(os.Stderr, "sort: disorder: %s\n", current)
				os.Exit(1)
			}
			temp = current
		}
	} else {
		for newScanner.Scan() {
			current = newScanner.Text()
			if less(current, temp, s.flag) {
				fmt.Fprintf(os.Stderr, "sort: disorder: %s\n", current)
				os.Exit(1)
			}
			temp = current
		}
	}

	return newScanner.Err()
}

func (s *Sort) Read(in *os.File) error {
	newScanner := bufio.NewScanner(in)

	if s.flag.check {
		return s.Check(in)
	}

	for newScanner.Scan() {
		heap.Push(s, newScanner.Text())
	}

	return newScanner.Err()
}

func (s *Sort) Write(out *os.File) {
	for s.Len() > 0 {
		fmt.Fprintln(out, heap.Pop(s).(string))
	}
}

func (s Sort) Len() int {
	return len(*s.heap)
}

func (s Sort) Less(i, j int) bool {
	if s.flag.reverse {
		return !less((*s.heap)[i], (*s.heap)[j], s.flag)
	}
	return less((*s.heap)[i], (*s.heap)[j], s.flag)
}

func (s Sort) Swap(i, j int) {
	(*s.heap)[i], (*s.heap)[j] = (*s.heap)[j], (*s.heap)[i]
}

func (s *Sort) Push(x any) {
	if s.flag.unique {
		key := string(getKey(x.(string), s.flag.key))
		if _, ok := s.keys[key]; ok {
			return
		}
		s.keys[key] = true
	}

	*s.heap = append(*s.heap, x.(string))
}

func (s *Sort) Pop() any {
	old := *s.heap
	n := len(old)
	x := old[n-1]
	*s.heap = old[0 : n-1]
	return x
}

func main() {
	k := flag.Int("k", 1, "define a restricted sort key")
	n := flag.Bool("n", false, "sort fields numerically by arithmetic value")
	r := flag.Bool("r", false, "Sort in reverse order")
	u := flag.Bool("u", false, "save only unique keys")
	m := flag.Bool("M", false, "sort by month abbreviations")
	b := flag.Bool("b", false, "ignore tail spaces")
	c := flag.Bool("c", false, "check if data is already sorted")
	h := flag.Bool("h", false, "sort fields numerically with suffixes")
	flag.Parse()

	if *n == true && *m == true {
		fmt.Fprintln(os.Stderr, "sort: n:M: mutually exclusive flags")
		os.Exit(1)
	} else if *n == true && *h == true {
		fmt.Fprintln(os.Stderr, "sort: h:n: mutually exclusive flags")
		os.Exit(1)
	} else if *h == true && *m == true {
		fmt.Fprintln(os.Stderr, "sort: h:M: mutually exclusive flags")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//defer file.Close()

	sort := NewSort(Flag{
		key:          *k,
		numeric:      *n,
		reverse:      *r,
		unique:       *u,
		month:        *m,
		ignoreBlanks: *b,
		check:        *c,
		humanNumeric: *h,
	})

	if err = sort.Read(file); err != nil {
		log.Fatalln(err)
	}

	sort.Write(os.Stdout)
}
