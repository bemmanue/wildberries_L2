package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagram(arr []string) map[string][]string {
	data := make(map[string][]string)
	keys := make(map[string]string)

	for _, s := range arr {
		s = strings.ToLower(s)
		s1 := []rune(s)

		sort.Slice(s1, func(i, j int) bool {
			return s1[i] < s1[j]
		})

		if _, ok := keys[string(s1)]; !ok {
			data[s] = append(data[s], s)
			keys[string(s1)] = s
			continue
		}

		key := keys[string(s1)]
		data[key] = append(data[key], s)
	}

	for i := range data {
		if len(data[i]) < 2 {
			delete(data, i)
		} else {
			sort.Strings(data[i])
		}
	}

	return data
}

func main() {
	arr := []string{"тяпка", "ПЯТАК", "Пятка", "бетон", "СЛИТОК", "столик", "листок"}

	data := findAnagram(arr)

	for _, i := range data {
		fmt.Println(i)
	}
}
