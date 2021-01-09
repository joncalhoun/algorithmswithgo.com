package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
