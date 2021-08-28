package read_write

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
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

	fNameIn, fNameOut := flag.Arg(0), flag.Arg(1)

	return flags, fNameIn, fNameOut, nil
}

func ReadFile(fname string) ([]string, error) {
	var fileData []string
	f, err := os.Open(fname)

	if err != nil {
		if fname == "" {
			f = os.Stdin
		} else {
			return fileData, err
		}
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("error open file %s", f.Name())
		}
	}(f)

	buf := bufio.NewScanner(f)
	var lines []string

	for buf.Scan() {
		lines = append(lines, buf.Text())
	}
	return lines, buf.Err()

}
func WriteFile(data []string, fname string) error {
	var out io.Writer

	if fname != "" {
		f, err := os.Create(fname)
		if err != nil {
			return err
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Printf("error write file %s\n", f.Name())
			}
		}(f)

		out = f
	} else {
		out = os.Stdout
	}

	writer := bufio.NewWriter(out)

	for _, str := range data {
		_, err := writer.WriteString(str + "\n")
		if err != nil {
			return errors.New("write to file error")
		}
	}
	return writer.Flush()
}
