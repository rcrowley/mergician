package html

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// InsertSorted inserts x into a, which is presumed to be sorted, maintaining
// its sorted order, and returns a (which has potentially been copied in order
// to be grown) and boolean true if x was inserted and false if it wasn't
// (because it was already present).
func InsertSorted[T constraints.Ordered](a []T, x T) ([]T, bool) {
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) {
		if a[i] == x {
			return a, false
		} else {
			b := make([]T, len(a)+1)
			copy(b[:i], a[:i])
			b[i] = x
			copy(b[i+1:], a[i:])
			return b, true
		}
	} else {
		return append(a, x), true
	}
}
