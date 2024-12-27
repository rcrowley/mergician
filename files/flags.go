package files

import (
	"flag"
	"strings"
)

type StringSliceFlag []string

func NewStringSliceFlag(flagSet *flag.FlagSet, name, usage string) *StringSliceFlag {
	f := &StringSliceFlag{}
	flagSet.Var(f, name, usage)
	return f
}

func (f *StringSliceFlag) String() string {
	return strings.Join(*f, ", ")
}

func (f *StringSliceFlag) Set(s string) error {
	*f = append(*f, s)
	return nil
}
