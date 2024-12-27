package files

import (
	"flag"
	"testing"
)

func TestStringSliceFlag(t *testing.T) {
	flagSet := &flag.FlagSet{}
	f := NewStringSliceFlag(flagSet, "t", "test")
	if s := f.String(); s != "" {
		t.Fatal(s)
	}
	f.Set("foo")
	if s := f.String(); s != "foo" {
		t.Fatal(s)
	}
	f.Set("bar")
	if s := f.String(); s != "foo, bar" {
		t.Fatal(s)
	}
}
