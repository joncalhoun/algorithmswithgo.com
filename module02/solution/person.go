package module02

// Person is used in a few sorting practice problems.
type Person struct {
	Age       int
	FirstName string
	LastName  string
}

type People []Person

func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool {
	a, b := p[i], p[j]
	if a.Age != b.Age {
		return a.Age < b.Age
	}
	if a.LastName != b.LastName {
		return a.LastName < b.LastName
	}
	return a.FirstName < b.FirstName
}
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
