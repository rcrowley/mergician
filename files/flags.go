package files

import "strings"

type StringSliceFlag []string

func (f *StringSliceFlag) String() string {
	return strings.Join(*f, ", ")
}

func (f *StringSliceFlag) Set(s string) bool {
	*f = append(*f, s)
	return true
}
