package read_write

import (
	"errors"
	"flag"
)

type Flags struct {
	CntF          bool
	RepeatF       bool
	NotRepeatF    bool
	CntSkipF      int
	RegisterSkipF bool
}

func GetFlags() (Flags, string, string, error) {

	flags := Flags{}

	flag.BoolVar(&flags.CntF, "c", false, "prefix lines by the number of occurrences")
	flag.BoolVar(&flags.RepeatF, "d", false, "only print duplicate lines, one for each group")
	flag.BoolVar(&flags.NotRepeatF, "u", false, "only print unique lines")
	flag.IntVar(&flags.CntSkipF, "f", 0, "avoid comparing the first N fields")
	flag.IntVar(&flags.CntSkipF, "s", 0, "avoid comparing the first N characters")
	flag.BoolVar(&flags.RegisterSkipF, "i", false, "ignore differences in case when comparing")

	flag.Parse()

	fNameIn := ""

	if fNameIn = flag.Arg(0); fNameIn == "" {
		err := errors.New("Error: The input file must be specified")
		return Flags{}, "", "", err
	}

	fNameOut := flag.Arg(1)

	return flags, fNameIn, fNameOut, nil
}
