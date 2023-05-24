package main

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

// State описывает тип символа: обычный символ, цифра, экранирующий символ
type State struct {
	symbol bool
	digit  bool
	escape bool
}

// repeat возвращает последовательность из n символов с
func repeat(c rune, n int) []rune {
	res := make([]rune, n)

	for i := 0; i < n; i++ {
		res[i] = c
	}

	return res
}

// isDigit проверяет, является ли символ цифрой
func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

// getCurrentState определяет принадлежность тип символа
func getCurrentState(c rune, previous State) State {
	current := State{}

	if previous.escape {
		if isDigit(c) || c == '\\' {
			current.symbol = true
		}
	} else {
		if isDigit(c) {
			current.digit = true
		} else if c == '\\' {
			current.escape = true
		} else {
			current.symbol = true
		}
	}

	return current
}

// Unpack распаковывает строку
func Unpack(s string) (string, error) {
	var result []rune

	str := []rune(s)
	previous := State{}
	current := State{}

	if len(str) == 0 {
		return "", nil
	}

	// Определяем тип первого символа
	current = getCurrentState(str[0], previous)

	// Если первый символ является цифрой, то строка невалидна
	if current.digit {
		return "", ErrInvalidString
	}

	// Если первый символ не является экранирующим символом, добавляем его в результат
	if !current.escape {
		result = append(result, str[0])
	}

	// проходим по всем элементам строки и в зависимости
	// от типа символа определяем дальнейшие действия
	for i := 1; i < len(str); i++ {
		previous = current
		current = getCurrentState(str[i], previous)

		if current.symbol {
			result = append(result, str[i])
		} else if current.digit {
			count := int(str[i] - '0')
			rep := repeat(result[len(result)-1], count)
			result = append(result[:len(result)-1], rep...)
		} else if current.escape {

		} else {
			return "", ErrInvalidString
		}
	}

	// Если последний символ является экранирующим, то строка невалидна
	if current.escape {
		return "", ErrInvalidString
	}

	return string(result), nil
}

func main() {
	result, err := Unpack(`a4bc2d5e`)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
