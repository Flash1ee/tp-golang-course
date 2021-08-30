package read_write

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Errors struct {
	UnknownFlag       error
	TogetherArgs      error
	SkipNegative      error
	IncorrectPosition error
}

func (e *Errors) init() {
	e.UnknownFlag = errors.New("You are using an unknown flag.\n" +
		"Run the program with the --help flag to output the supported commands")
	e.TogetherArgs = errors.New("Flags [-c -d -u] cannot be used together")
	e.SkipNegative = errors.New("Count of skipped symbols(words) must be a positive number")
}

type Flags struct {
	CntF          bool
	RepeatF       bool
	NotRepeatF    bool
	CntSkipWordsF int
	CntSkipCharsF int
	RegisterSkipF bool
	FNameIn       string
	FNameOut      string
}

func B2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func GetFlags(progName string, args []string) (Flags, string, error) {

	var output bytes.Buffer
	confFlags := Flags{}

	flags := flag.NewFlagSet(progName, flag.ContinueOnError)
	flags.SetOutput(&output)

	myErrors := Errors{}
	myErrors.init()

	flags.BoolVar(&confFlags.CntF, "c", false, "prefix lines by the number of occurrences")
	flags.BoolVar(&confFlags.RepeatF, "d", false, "only print duplicate lines, one for each group")
	flags.BoolVar(&confFlags.NotRepeatF, "u", false, "only print unique lines")
	flags.IntVar(&confFlags.CntSkipWordsF, "f", 0, "avoid comparing the first N fields")
	flags.IntVar(&confFlags.CntSkipCharsF, "s", 0, "avoid comparing the first N characters")
	flags.BoolVar(&confFlags.RegisterSkipF, "i", false, "ignore differences in case when comparing")

	err := flags.Parse(args)
	if err != nil {
		return Flags{}, output.String(), UnknownFlagError
	}
	if confFlags.CntSkipCharsF < 0 || confFlags.CntSkipWordsF < 0 {
		return Flags{}, output.String(), SkipNegative
	}

	tmp := B2i(confFlags.CntF) + B2i(confFlags.NotRepeatF) + B2i(confFlags.RepeatF)
	if tmp != 0 && tmp != 1 {
		return Flags{}, output.String(), TogetherArgs
	}
	if flags.NArg() > 2 {
		return Flags{}, output.String(), IncorrectPosition
	}

	confFlags.FNameIn, confFlags.FNameOut = flags.Arg(0), flags.Arg(1)

	return confFlags, output.String(), nil
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
