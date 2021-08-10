package main

import (
	"io"
	"strings"
	"testing"

	"github.com/felipebool/prodcon/internal/cache"
	_ "github.com/jackc/pgx/stdlib"
)

func Test_warmUpCache(t *testing.T) {
	//strings.NewReader("This is a test string")
	cases := map[string]struct {
		input    io.Reader
		expected int
	}{
		"empty file": {
			input:    strings.NewReader(""),
			expected: 0,
		},
		"single entry file": {
			input:    strings.NewReader("aaabbbc\n"),
			expected: 1,
		},
		"multiple entries file": {
			input:    strings.NewReader("aaabbbc\naaabbbc\naaabbbc\n"),
			expected: 3,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			c := cache.New()
			got, _ := warmUpCache(c, tc.input)
			if got != tc.expected {
				t.Errorf("got: %d; expected: %d", got, tc.expected)
			}
		})
	}
}
