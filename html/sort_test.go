package html

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var a []string
	for i, x := range []string{
		"foo",  // init
		"bar",  // first
		"baz",  // middle
		"quux", // last
	} {
		var ok bool
		a, ok = InsertSorted(a, x)
		if len(a) != i+1 || !sort.StringsAreSorted(a) {
			t.Fatal(a)
		}
		if !ok {
			t.Fatal(ok)
		}
	}
	if _, ok := InsertSorted(a, "foo"); ok {
		t.Fatal(ok)
	}
	if a[0] != "bar" || a[1] != "baz" || a[2] != "foo" || a[3] != "quux" {
		t.Fatal(a)
	}
}
