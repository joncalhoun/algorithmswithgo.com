package sorttest

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

// TestString is a helper for testing functions that sort stringeger slices.
func TestString(t *testing.T, sortFn func([]string)) {
	seed := time.Now().UnixNano()
	t.Log("Seed for random cases:", seed)
	rand.Seed(seed)

	for name, list := range map[string][]string{
		"sorted":     []string{"apple", "banana", "cat", "dog"},
		"reverse":    []string{"dog", "cat", "ball", "alphabet"},
		"duplicates": []string{"alphabet", "ball", "cat", "alphabet", "ball", "cat"},
		"similar":    []string{"apple", "app", "alligator", "all", "all-in", "alphabet", "apoplexy", "apology", "apologize"},
		"random-1000-20": func() []string {
			alphabet := "abcdefghijklmnopqrstuvwxyz"
			var list []string
			for i := 0; i < 1000; i++ {
				var sb strings.Builder
				for j := 0; j < 20; j++ {
					sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
					if rand.Intn(3) == 0 {
						break
					}
				}
				list = append(list, sb.String())
			}
			return list
		}(),
	} {
		t.Run(name, func(t *testing.T) {
			want := make([]string, len(list))
			for i, val := range list {
				want[i] = val
			}
			sort.Strings(want)
			sortFn(list)
			errorCount := 0
			if len(list) != len(want) {
				t.Fatalf("got len %d; want %d", len(list), len(want))
			}
			for i := 0; i < len(want); i++ {
				if errorCount >= 5 {
					t.Fatalf("additional errors omitted for brevity")
				}
				if list[i] != want[i] {
					errorCount++
					t.Errorf("list[%d] = %d; want %d", i, list[i], want[i])
				}
			}
		})
	}
}
