package read_write

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

type Flags struct {
	CntF bool
	RepeatF bool
	NotRepeatF bool
	CntSkipF int
	RegisterSkipF bool
}

func (f *Flags) getCntInit() int {
	cnt := B2i(f.CntF) + B2i(f.RepeatF) + B2i(f.NotRepeatF) + f.CntSkipF + B2i(f.RegisterSkipF)
	return cnt
}

func GetFlags() (Flags, string, string, error) {

	flags := Flags{}
	cntParams := len(os.Args)

	flag.BoolVar(&flags.CntF,"c", false, "prefix lines by the number of occurrences")
	flag.BoolVar(&flags.RepeatF,"d", false, "only print duplicate lines, one for each group")
	flag.BoolVar(&flags.NotRepeatF, "u", false, "only print unique lines")
	flag.IntVar(&flags.CntSkipF, "f",0, "avoid comparing the first N fields")
	flag.IntVar(&flags.CntSkipF, "s",0, "avoid comparing the first N characters")
	flag.BoolVar(&flags.RegisterSkipF, "i", false, "ignore differences in case when comparing")

	flag.Parse()

	cntFlags := flags.getCntInit()

	if cntParams < 2 ||  cntFlags == cntParams - 1 {
		err := errors.New("Error: The input file must be specified")
		return Flags{}, "", "", err
	}
	fmt.Println("Cnt params is", cntFlags)
	ioFiles := os.Args[cntFlags + 1:]

	fNameIn := ioFiles[0]
	fNameOut := ""
	if len(ioFiles) == 2 {
		fNameOut = ioFiles[1]
	}

	return flags, fNameIn, fNameOut, nil
}