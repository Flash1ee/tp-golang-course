package read_write

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

type UniqRes struct {
	Str string
	Cnt int
}

func (u UniqRes) WriteRepeatStr(writer *bufio.Writer) error {
	var err error

	if u.Cnt > 1 {
		_, err = writer.WriteString(u.Str + "\n")
	}

	if err != nil {
		return errors.New("write to file error")
	}
	return nil
}
func (u UniqRes) WriteNotRepeatStr(writer *bufio.Writer) error {
	var err error

	if u.Cnt == 1 {
		_, err = writer.WriteString(u.Str + "\n")
	}

	if err != nil {
		return errors.New("write to file error")
	}
	return nil
}
func (u UniqRes) WriteWithCntStr(writer *bufio.Writer) error {
	var err error

	_, err = writer.WriteString(strconv.Itoa(u.Cnt) + "\t" + u.Str + "\n")

	if err != nil {
		return errors.New("write to file error")
	}
	return nil
}

func (u UniqRes) WriteDefault(writer *bufio.Writer) error {
	var err error
	_, err = writer.WriteString(u.Str + "\n")
	if err != nil {
		return errors.New("write to file error")
	}
	return nil
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
func WriteFile(cnts []UniqRes, flags Flags) error {
	var out io.Writer
	var err error

	if flags.FNameOut != "" {
		f, err := os.Create(flags.FNameOut)
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

	for _, val := range cnts {
		if flags.CntF {
			err = val.WriteWithCntStr(writer)
		} else if flags.RepeatF {
			err = val.WriteRepeatStr(writer)
		} else if flags.NotRepeatF {
			err = val.WriteNotRepeatStr(writer)
		} else {
			err = val.WriteDefault(writer)
		}
		if err != nil {
			return errors.New("write to file error")
		}
	}
	return writer.Flush()
}
