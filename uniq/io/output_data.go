package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (u UniqRes) WriteRepeatStr(writer *bufio.Writer) error {
	var err error

	if u.Cnt > 1 {
		_, err = writer.WriteString(u.Str + "\n")
	}

	if err != nil {
		return fmt.Errorf("write to file error")
	}
	return nil
}

func (u UniqRes) WriteNotRepeatStr(writer *bufio.Writer) error {
	var err error

	if u.Cnt == 1 {
		_, err = writer.WriteString(u.Str + "\n")
	}

	if err != nil {
		return fmt.Errorf("write to file error")
	}
	return nil
}

func (u UniqRes) WriteWithCntStr(writer *bufio.Writer) error {
	var err error

	_, err = writer.WriteString(strconv.Itoa(u.Cnt) + "\t" + u.Str + "\n")

	if err != nil {
		return fmt.Errorf("write to file error")
	}
	return nil
}

func (u UniqRes) WriteDefault(writer *bufio.Writer) error {
	var err error
	_, err = writer.WriteString(u.Str + "\n")
	if err != nil {
		return fmt.Errorf("write to file error")
	}
	return nil
}

func WriteFile(counts []UniqRes, flags Flags) error {
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

	for _, val := range counts {
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
			return fmt.Errorf("write to file error")
		}
	}
	return writer.Flush()
}
