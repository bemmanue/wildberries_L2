package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type StringHeap []string

// Flag описывает флаги
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

// Sort описывает структуру сортировки
type Sort struct {
	keys map[string]bool
	flag Flag
	heap *StringHeap
}

// less возвращает меньшее значение
func less(str1, str2 string, flag Flag) bool {
	s1 := getKey(str1, flag.key)
	s2 := getKey(str2, flag.key)

	// пропуск бланковых символов
	if flag.ignoreBlanks {
		s1 = skipBlanks(s1)
		s2 = skipBlanks(s2)

		flag.ignoreBlanks = false
	}

	// сравнивание по числам
	if flag.numeric {
		var n1, n2 int
		s1, n1 = getNum(s1)
		s2, n2 = getNum(s2)

		if n1 < n2 {
			return true
		} else if n1 > n2 {
			return false
		}
	}

	// сравнивание по числам с учетом суффиксов
	if flag.humanNumeric {
		var n1, n2 *big.Int
		s1, n1 = getHumanNum(s1)
		s2, n2 = getHumanNum(s2)

		if n1.Cmp(n2) < 0 {
			return true
		} else if n1.Cmp(n2) > 0 {
			return false
		}
	}

	// сравнивание по месяцам
	if flag.month {
		var m1, m2 int
		s1, m1 = getMonth(s1)
		s2, m2 = getMonth(s2)

		if m1 < m2 {
			return true
		} else if m1 > m2 {
			return false
		}
	}

	// сравнивание по символам
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}

	// проверка случая, когда одна строка является префиксом другой
	if len(s1) < len(s2) {
		return true
	} else if len(s1) > len(s2) {
		return false
	}

	// сравнивание начала строк, если они равны с заданного столбца
	if flag.key > 1 {
		flag.key = 1
		return less(str1, str2, flag)
	}

	return true
}

// getKey возвращает ключевое поле,
// по которому должна производиться сортировка
func getKey(s string, k int) []rune {
	res := []rune(s)
	i := 0

	for k > 1 {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		for i < len(s) && s[i] != ' ' {
			i++
		}
		if i < len(s) {
			i++
		}
		k--
	}

	return res[i:]
}

// getNum возвращает число с учетом суффикса,
// по которому будет производиться сортировка
func getNum(s []rune) ([]rune, int) {
	var num []rune
	i := 0

	for i < len(s) && isSpace(s[i]) {
		i++
	}

	if i == len(s) || !isDigit(s[i]) {
		return s, 0
	}

	for ; i < len(s) && isDigit(s[i]); i++ {
		num = append(num, s[i])
	}

	n, _ := strconv.Atoi(string(num))

	return s[i:], n
}

// getHumanNum возвращает число с ,
// по которому будет производиться сортировка
func getHumanNum(s []rune) ([]rune, *big.Int) {
	var num []rune

	n := big.NewInt(0)
	ratio := big.NewInt(1)

	i := 0

	for i < len(s) && isSpace(s[i]) {
		i++
	}

	if i == len(s) || !isDigit(s[i]) {
		return s, n
	}

	for ; i < len(s) && isDigit(s[i]); i++ {
		num = append(num, s[i])
	}

	if i < len(s) {
		ratio = getMetric(s[i])
		i++
	}

	n.SetString(string(num), 10)
	n.Mul(n, ratio)

	return s[i:], n
}

// getMetric возвращает числовое значение суффикса
func getMetric(c rune) *big.Int {
	n := big.NewInt(1)

	if c == 'k' || c == 'K' {
		n, _ = n.SetString("1000", 10)
		return n
	} else if c == 'm' || c == 'M' {
		n, _ = n.SetString("1000000", 10)
		return n
	} else if c == 'g' || c == 'G' {
		n, _ = n.SetString("1000000000", 10)
		return n
	} else if c == 't' || c == 'T' {
		n, _ = n.SetString("1000000000000", 10)
		return n
	} else if c == 'p' || c == 'P' {
		n, _ = n.SetString("1000000000000000", 10)
		return n
	} else if c == 'e' || c == 'E' {
		n, _ = n.SetString("1000000000000000000", 10)
		return n
	} else if c == 'z' || c == 'Z' {
		n, _ = n.SetString("1000000000000000000000", 10)
		return n
	} else if c == 'y' || c == 'Y' {
		n, _ = n.SetString("1000000000000000000000000", 10)
		return n
	}

	return n
}

// skipBlanks пропускает пробелы
func skipBlanks(s []rune) []rune {
	i := 0
	for i < len(s) && isSpace(s[i]) {
		i++
	}
	return s[i:]
}

// getMonth возвращает числовое значение месяца
func getMonth(s []rune) ([]rune, int) {
	i := 0

	for i < len(s) && isSpace(s[i]) {
		i++
	}

	month := strings.ToUpper(string(s[i : i+3]))

	switch month {
	case "JAN":
		return s[i+3:], 1
	case "FEB":
		return s[i+3:], 2
	case "MAR":
		return s[i+3:], 3
	case "APR":
		return s[i+3:], 4
	case "MAY":
		return s[i+3:], 5
	case "JUN":
		return s[i+3:], 6
	case "JUL":
		return s[i+3:], 7
	case "AUG":
		return s[i+3:], 8
	case "SEP":
		return s[i+3:], 9
	case "OCT":
		return s[i+3:], 10
	case "NOV":
		return s[i+3:], 11
	case "DEC":
		return s[i+3:], 12
	}

	return s, 0
}

// isSpace проверяет, является ли символ пробельным символом
func isSpace(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r' {
		return true
	}
	return false
}

// isDigit проверяет, является ли символ цифрой
func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

// NewSort инициализирует структуру сортировки
func NewSort(flag Flag) *Sort {
	sort := &Sort{
		keys: make(map[string]bool),
		flag: flag,
		heap: &StringHeap{},
	}

	heap.Init(sort)

	return sort
}

// Check проверяет, является ли файл отсортированным по строкам
func (s *Sort) Check(in *os.File) error {
	var current, temp string
	newScanner := bufio.NewScanner(in)

	newScanner.Scan()
	temp = newScanner.Text()

	for newScanner.Scan() {
		current = newScanner.Text()
		if less(temp, current, s.flag) != s.flag.reverse {
			fmt.Fprintf(os.Stderr, "sort: disorder: %s\n", current)
			os.Exit(1)
		}
		temp = current
	}

	return newScanner.Err()
}

// Read считывает строки из файла и сохраняет их в сортированном виде
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

// Write выводит строки в сортированном виде
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

	// проверяем совместимость флагов
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

	// открываем файл на чтение
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// инициализируем структуру сортировки
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

	// считываем и сохраняем отсортированные строки
	if err = sort.Read(file); err != nil {
		log.Fatalln(err)
	}

	// выводим результат сортировки
	sort.Write(os.Stdout)
}
