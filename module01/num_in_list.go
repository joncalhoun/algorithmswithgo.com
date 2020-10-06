package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}

	return false
}
