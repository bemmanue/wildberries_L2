package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	data := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "valid",
			input:    `a4bc2d5e`,
			expected: `aaaabccddddde`,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `abcd`,
			expected: `abcd`,
			err:      nil,
		},
		{
			name:     "invalid",
			input:    `45`,
			expected: ``,
			err:      ErrInvalidString,
		},
		{
			name:     "valid",
			input:    ``,
			expected: ``,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `qwe\4\5`,
			expected: `qwe45`,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `qwe\45`,
			expected: `qwe44444`,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `\\\\\5`,
			expected: `\\5`,
			err:      nil,
		},
		{
			name:     "valid",
			input:    `\`,
			expected: ``,
			err:      ErrInvalidString,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := Unpack(d.input)

			if result != d.expected {
				t.Errorf("Expected %s, got %s", d.expected, result)
			}

			if err != d.err {
				t.Errorf("Expected %s, got %s", d.err.Error(), err.Error())
			}
		})
	}
}
