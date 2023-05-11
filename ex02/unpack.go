package main

import (
	"errors"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

type State struct {
	symbol bool
	digit  bool
	escape bool
}

func repeat(c rune, count int) []rune {
	res := make([]rune, count)

	for i := 0; i < count; i++ {
		res[i] = c
	}

	return res
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

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

func Unpack(s string) (string, error) {
	var result []rune

	str := []rune(s)
	previous := State{}
	current := State{}

	if len(str) == 0 {
		return "", nil
	}

	current = getCurrentState(str[0], previous)

	if current.digit {
		return "", ErrInvalidString
	}

	if !current.escape {
		result = append(result, str[0])
	}

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

	if current.escape {
		return "", ErrInvalidString
	}

	return string(result), nil
}
