package cache_test

import (
	"sync"
	"testing"

	"github.com/felipebool/prodcon/internal/cache"
)

func Test_cache(t *testing.T) {
	cases := map[string]struct {
		elements []string
		expected map[string]int
	}{
		"empty list": {
			elements: []string{},
			expected: map[string]int{},
		},
		"single element": {
			elements: []string{"aaa"},
			expected: map[string]int{
				"aaa": 1,
			},
		},
		"multiple element": {
			elements: []string{"aaa", "aaa", "aaa", "bbb", "ccc"},
			expected: map[string]int{
				"aaa": 3,
				"bbb": 1,
				"ccc": 1,
			},
		},
	}

	for label := range cases {
		tc := cases[label]

		// covers set and fetch
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			c := cache.New()
			wg := &sync.WaitGroup{}

			// concurrent access to cache
			for _, key := range tc.elements {
				wg.Add(1)
				go func(key string) {
					c.Save(key)
					wg.Done()
				}(key)
			}
			wg.Wait()

			for key, value := range tc.expected {
				got := c.Fetch(key)
				if got != value {
					t.Errorf("[%s] got: %d; expected: %d", key, got, value)
				}
			}
		})
	}
}
