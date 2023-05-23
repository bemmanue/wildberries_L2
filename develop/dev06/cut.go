package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrIllegalListValue = errors.New("cut: [-f] list: illegal list value")
	ErrBadDelimiter     = errors.New("cut: [-d] delim: bad delimiter")
)

type Flag struct {
	f string
	d string
	s bool
}

func getKeys(m map[int]bool) []int {
	s := make([]int, 0, len(m))

	for i := range m {
		s = append(s, i)
	}

	sort.Ints(s)
	return s
}

func getList(fields string) ([]int, int, error) {
	data := make(map[int]bool)
	endlessFrom := 0

	if len(fields) == 0 {
		return nil, 0, ErrIllegalListValue
	}

	arr := strings.Split(fields, ",")

	for _, s := range arr {
		if len(s) == 0 {
			return nil, 0, ErrIllegalListValue
		}

		v := strings.Split(s, "-")

		switch len(v) {
		case 1:
			n := 0

			if n, _ = strconv.Atoi(v[0]); n == 0 {
				return nil, 0, ErrIllegalListValue
			}

			data[n] = true
		case 2:
			n, m := 0, 0

			if len(v[0]) == 0 && len(v[1]) == 0 {
				return nil, 0, ErrIllegalListValue
			} else if len(v[0]) == 0 {
				n = 1

				if m, _ = strconv.Atoi(v[1]); m == 0 {
					return nil, 0, ErrIllegalListValue
				}
			} else if len(v[1]) == 0 {
				if n, _ = strconv.Atoi(v[0]); n == 0 {
					return nil, 0, ErrIllegalListValue
				}

				if endlessFrom == 0 || n < endlessFrom {
					endlessFrom = n
				}
			} else {
				if n, _ = strconv.Atoi(v[0]); n == 0 {
					return nil, 0, ErrIllegalListValue
				}

				if m, _ = strconv.Atoi(v[0]); m == 0 {
					return nil, 0, ErrIllegalListValue
				}
			}

			for n <= m {
				data[n] = true
				n++
			}
		default:
			return nil, 0, ErrIllegalListValue
		}
	}

	list := getKeys(data)

	return list, endlessFrom, nil
}

func getDelim(delim string) (string, error) {
	d := []rune(delim)

	if len(d) != 1 {
		return "", ErrBadDelimiter
	}

	return delim, nil
}

func Cut(flags Flag, in, out *os.File) error {
	list, endlessFrom, err := getList(flags.f)
	if err != nil {
		return err
	}

	delim, err := getDelim(flags.d)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, delim)

		if flags.s && len(arr) == 1 {
			continue
		}

		switch endlessFrom {
		case 0:
			for i := 0; i < len(list) && list[i] < len(arr); i++ {
				fmt.Fprint(out, arr[list[i]-1], delim)
			}
		default:
			if len(arr) == 1 {
				fmt.Fprint(out, arr[0], delim)
			}

			for i := 0; i < len(list) && list[i] < len(arr) && list[i] < endlessFrom; i++ {
				fmt.Fprint(out, arr[list[i]-1], delim)
			}

			for i := endlessFrom - 1; i < len(arr); i++ {
				fmt.Fprint(out, arr[i], delim)
			}
		}

		fmt.Fprintln(out)
	}

	return nil
}

func isFlagPassed(name string) bool {
	found := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

func main() {
	f := flag.String("f", "", "Выбрать поля")
	d := flag.String("d", "\t", "Использовать другой разделитель")
	s := flag.Bool("s", false, "Показать строки с разделителем")
	flag.Parse()

	if !isFlagPassed("f") {
		fmt.Fprintln(os.Stderr, "usage: cut -f list [-s] [-d delim]")
	}

	flags := Flag{
		f: *f,
		d: *d,
		s: *s,
	}

	if err := Cut(flags, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
