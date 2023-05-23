package main

import (
	"math/big"
	"strconv"
	"strings"
)

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

func skipBlanks(s []rune) []rune {
	i := 0
	for i < len(s) && isSpace(s[i]) {
		i++
	}
	return s[i:]
}

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

func isSpace(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r' {
		return true
	}
	return false
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
