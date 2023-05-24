package main

import "testing"

func Test_findAnagram(t *testing.T) {
	data := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input:    []string{"тяпка", "ПЯТАК", "Пятка", "бетон", "СЛИТОК", "столик", "листок"},
			expected: map[string][]string{"слиток": {"листок", "слиток", "столик"}, "тяпка": {"пятак", "пятка", "тяпка"}},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, d := range data {
		t.Run("find anagram", func(t *testing.T) {
			result := findAnagram(d.input)

			if len(result) != len(d.expected) {
				t.Fatalf("Expected %s, got %s", d.expected, result)
			}

			for key, a1 := range d.expected {
				a2, ok := result[key]
				if !ok {
					t.Fatalf("Expected %s, got %s", d.expected, result)
				}

				if len(a1) != len(a2) {
					t.Fatalf("Expected %s, got %s", d.expected, result)
				}

				for i := 0; i < len(a1) && i < len(a2); i++ {
					if a1[i] != a2[i] {
						t.Fatalf("Expected %s, got %s", d.expected, result)
					}
				}
			}
		})
	}
}
