package io

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFilePositive(t *testing.T) {
	tests := []struct {
		data     []UniqRes
		flags    Flags
		expected []string
	}{
		{
			data:     []UniqRes{{Str: "Hello", Cnt: 2}, {Str: "World", Cnt: 1}},
			expected: []string{"Hello", "World"},
		},
		{
			data:     []UniqRes{{Str: "Hello", Cnt: 2}, {Str: "World", Cnt: 1}},
			flags:    Flags{FNameOut: "test"},
			expected: []string{"Hello", "World"},
		},
		{
			data:     []UniqRes{{Str: "Hello", Cnt: 2}, {Str: "World", Cnt: 1}, {Str: "Guys", Cnt: 4}},
			flags:    Flags{CntF: true},
			expected: []string{"2\tHello", "1\tWorld", "4\tGuys"},
		},
		{
			data:     []UniqRes{{Str: "Hello", Cnt: 1}, {Str: "World", Cnt: 1}, {Str: "Guys", Cnt: 4}},
			flags:    Flags{RepeatF: true},
			expected: []string{"Guys"},
		},
		{
			data:     []UniqRes{{Str: "Hello", Cnt: 2}, {Str: "World", Cnt: 1}, {Str: "Guys", Cnt: 1}},
			flags:    Flags{NotRepeatF: true},
			expected: []string{"World", "Guys"},
		},
	}
	for _, pair := range tests {
		t.Run("Test write to file", func(t *testing.T) {
			fName := pair.flags.FNameOut
			var err error
			if fName == "" {
				tempFile, err := os.Create("test")
				if err != nil {
					log.Fatal(err)
				}
				os.Stdout = tempFile
				fName = "test"
			}
			err = WriteFile(pair.data, pair.flags)
			assert.Nil(t, err)

			f, err := os.Open(fName)
			if err != nil {
				log.Fatal(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					log.Fatal(err)
				}
				err = os.Remove(f.Name())
				if err != nil {
					log.Fatal(err)
				}
			}(f)
			var lines []string
			reader := bufio.NewScanner(f)
			for reader.Scan() {
				lines = append(lines, reader.Text())
			}
			assert.Equal(t, pair.expected, lines,
				"incorrect write to file\ngot: %v\nexpected: %v", lines, pair.expected)
		})
	}
}
