package io

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func GetFlags(programName string, args []string) (Flags, string, error) {
	var output bytes.Buffer
	confFlags := Flags{}

	flags := flag.NewFlagSet(programName, flag.ContinueOnError)
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

	tmp := boolToInt(confFlags.CntF) + boolToInt(confFlags.NotRepeatF) + boolToInt(confFlags.RepeatF)
	if tmp != 0 && tmp != 1 {
		return Flags{}, output.String(), TogetherArgs
	}
	if flags.NArg() > 2 {
		return Flags{}, output.String(), IncorrectPosition
	}

	confFlags.FNameIn, confFlags.FNameOut = flags.Arg(0), flags.Arg(1)

	return confFlags, output.String(), nil
}

func ReadFile(fileName string) ([]string, error) {
	var fileData []string
	f, err := os.Open(fileName)
	if err != nil {
		if fileName == "" {
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
