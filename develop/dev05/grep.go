package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
)

type Flag struct {
	after      uint
	before     uint
	context    uint
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// printLine печатает форматированную строку
func printLine(lineNum bool, num int, line string) {
	switch lineNum {
	case true:
		fmt.Printf("%d:%s\n", num, line)
	case false:
		fmt.Printf("%s\n", line)
	}
}

// sub возвращает разность двух чисел.
// Если разность меньше 0, возвращается 0
func sub(a, b int) int {
	if a <= b {
		return 0
	}
	return a - b
}

// sum возвращает сумму двух чисел.
// Если сумма больше максимального числа, возвращается максимальное число
func sum(a, b int) int {
	max := math.MaxInt

	if b >= max-a {
		return max
	}

	return a + b
}

// max возвращает большее число
func max(a, b uint) int {
	max := math.MaxInt

	if a >= uint(max) || b >= uint(max) {
		return max
	}

	if a >= b {
		return int(a)
	}

	return int(b)
}

// writeCount выводит количество строк, соответствующих паттерну
func writeCount(flags Flag, pattern string, lines []string) {
	reg := regexp.MustCompile(pattern)
	count := 0

	for _, s := range lines {
		if reg.MatchString(s) != flags.invert {
			count++
		}
	}

	fmt.Println(count)
}

// writeLines выводит строки, соответствующие паттерну
func writeLines(flags Flag, pattern string, lines []string) {
	reg := regexp.MustCompile(pattern)
	printedMap := make(map[int]bool)

	before := max(flags.before, flags.context)
	after := max(flags.after, flags.context)

	for i := range lines {
		if reg.MatchString(lines[i]) != flags.invert {
			a := sub(i, before)
			b := sum(i, after)

			for j := i; j >= a && !printedMap[j]; j-- {
				printedMap[j] = true
			}

			for j := i; j <= b && j < len(printedMap); j++ {
				printedMap[j] = true
			}
		}
	}

	printedList := getKeys(printedMap)

	group := flags.after > 0 || flags.before > 0 || flags.context > 0

	if len(printedList) != 0 {
		printLine(flags.lineNum, printedList[0], lines[printedList[0]])
	}

	for i := 1; i < len(printedList); i++ {
		if group && printedList[i]-printedList[i-1] > 1 {
			fmt.Println("--")
		}
		printLine(flags.lineNum, printedList[i]+1, lines[printedList[i]])
	}
}

func getKeys(m map[int]bool) []int {
	s := make([]int, 0, len(m))

	for i := range m {
		s = append(s, i)
	}

	sort.Ints(s)
	return s
}

// Grep ищет строки, соответствующие заданному паттерну
func Grep(flags Flag, pattern string, lines []string) {
	if flags.ignoreCase {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}

	if flags.fixed {
		pattern = fmt.Sprintf("^%s$", pattern)
	}

	if flags.count {
		writeCount(flags, pattern, lines)
	} else {
		writeLines(flags, pattern, lines)
	}
}

// readLines открывает файл и считывает из него строки
func readLines(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	//defer file.Close()

	newScanner := bufio.NewScanner(file)

	for newScanner.Scan() {
		lines = append(lines, newScanner.Text())
	}

	return lines, newScanner.Err()
}

func main() {
	// парсинг флагов
	A := flag.Uint("A", 0, `Печатать +N строк после совпадения`)
	B := flag.Uint("B", 0, `Печатать +N строк до совпадения`)
	C := flag.Uint("C", 0, `Печатать ±N строк вокруг совпадения`)
	c := flag.Bool("c", false, `Количество строк`)
	i := flag.Bool("i", false, `Игнорировать регистр`)
	v := flag.Bool("v", false, `Вместо совпадения, исключать`)
	F := flag.Bool("F", false, `Точное совпадение со строкой, не паттерн`)
	n := flag.Bool("n", false, `Напечатать номер строки`)
	flag.Parse()

	// заполнение структуры флагов
	flags := Flag{
		after:      *A,
		before:     *B,
		context:    *C,
		count:      *c,
		ignoreCase: *i,
		invert:     *v,
		fixed:      *F,
		lineNum:    *n,
	}

	// получение ключевого слова
	pattern := flag.Arg(0)

	// считывание строк из файла
	lines, err := readLines(flag.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Grep(flags, pattern, lines)
}
