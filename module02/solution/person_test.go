package module02

import (
	"sort"
	"testing"
)

// testPeople is a helper for testing functions that sort people by age then
// last name then first name.
func testPeople(t *testing.T, sortFn func([]Person)) {
	for name, list := range map[string][]Person{
		"sorted": []Person{
			{12, "Billy", "Tables"},
			{12, "Bobby", "Tables"},
			{12, "Jordan", "Tables"},
			{12, "Alex", "Zero"},
			{21, "Frank", "Smith"},
			{33, "Bob", "Smilesalot"},
			{45, "Johnny", "Testuser"},
			{65, "Harry", "Hippo"},
			{71, "Thomas", "Train"},
			{53, "Percy", "Engine"},
		},
		"reverse": []Person{
			{71, "Thomas", "Train"},
			{65, "Harry", "Hippo"},
			{53, "Percy", "Engine"},
			{45, "Johnny", "Testuser"},
			{33, "Bob", "Smilesalot"},
			{21, "Frank", "Smith"},
			{12, "Alex", "Zero"},
			{12, "Jordan", "Tables"},
			{12, "Bobby", "Tables"},
			{12, "Billy", "Tables"},
		},
		"duplicates": []Person{
			{53, "Percy", "Engine"},
			{45, "Johnny", "Testuser"},
			{12, "Billy", "Tables"},
			{65, "Harry", "Hippo"},
			{33, "Bob", "Smilesalot"},
			{12, "Bobby", "Tables"},
			{12, "Alex", "Zero"},
			{45, "Johnny", "Testuser"},
			{12, "Billy", "Tables"},
			{21, "Frank", "Smith"},
			{71, "Thomas", "Train"},
			{21, "Frank", "Smith"},
			{12, "Jordan", "Tables"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			want := make([]Person, len(list))
			for i, val := range list {
				want[i] = val
			}
			sort.Sort(People(want))
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
					t.Errorf("list[%d] = %v; want %v", i, list[i], want[i])
				}
			}
		})
	}
}
