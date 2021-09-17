package io

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFlagsPositive(t *testing.T) {
	tests := []struct {
		description string
		args        []string
		conf        Flags
	}{
		{"empty flags", []string{}, Flags{}},
		{"one flag -c", []string{"-c"}, Flags{true, false, false, 0, 0, false, "", ""}},
		{"flag -f with argument 10", []string{"-f", "10"}, Flags{false, false, false, 10, 0, false, "", ""}},
		{"one flag -u", []string{"-u"}, Flags{false, false, true, 0, 0, false, "", ""}},
		{"flag -f with argument 10 usage =", []string{"-f=10"}, Flags{false, false, false, 10, 0, false, "", ""}},
		{"with input file", []string{"in.txt"}, Flags{false, false, false, 0, 0, false, "in.txt", ""}},
		{"with input and out file", []string{"in.txt", "out.txt"}, Flags{false, false, false, 0, 0, false, "in.txt", "out.txt"}},
		{"flags and io files", []string{"-c", "-s", "10", "in.txt", "out.txt"}, Flags{true, false, false, 0, 10, false, "in.txt", "out.txt"}},
		{"flags ordinary and with args + io files", []string{"-d", "-i", "-f", "10", "-s=20", "in.txt", "out.txt"}, Flags{false, true, false, 10, 20, true, "in.txt", "out.txt"}},
	}
	for _, pair := range tests {
		t.Run(pair.description, func(t *testing.T) {
			res, _, err := GetFlags(os.Args[0], pair.args)
			assert.Equal(t, pair.conf, res, pair.description+"\ngot %v\nexpected %v", pair.args, pair.conf)
			assert.Nil(t, err)
		})
	}
}

func TestGetFlagsNegative(t *testing.T) {
	tests := []struct {
		description string
		args        []string
		conf        Flags
		errstr      string
	}{
		{"incorrect flag", []string{"-k"}, Flags{}, UnknownFlagError.Error()},
		{"-d -u -c cannot be used together", []string{"-d", "-u", "-c"}, Flags{}, TogetherArgs.Error()},
		{"count skip of words must be not negative num", []string{"-c", "-f", "-10"}, Flags{}, SkipNegative.Error()},
		{"positional arguments should be placed at the end", []string{"in.txt", "-c", "-s", "10"}, Flags{}, IncorrectPosition.Error()},
		{"positional arguments should be placed at the end", []string{"-c", "in.txt", "-s", "10"}, Flags{}, IncorrectPosition.Error()},
	}

	for _, pair := range tests {
		t.Run(pair.description, func(t *testing.T) {
			res, _, err := GetFlags(os.Args[0], pair.args)
			assert.Equal(t, pair.conf, res, pair.description+"\nconf got %+v, want %+v", pair.conf, res)
			assert.Equal(t, pair.errstr, err.Error())
			assert.NotNil(t, err, "err got %v, want nil", err)
		})
	}
}

func TestReadFilePositive(t *testing.T) {
	tests := []struct {
		fName    string
		data     string
		expected []string
	}{
		{data: "Hello\nWorld\n", expected: []string{"Hello", "World"}},
		{data: ""},
		{fName: "test.txt", data: "Lorem\nIpsum\n", expected: []string{"Lorem", "Ipsum"}},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			fname := pair.fName
			tempFile, err := ioutil.TempFile("", fname)
			if err != nil {
				log.Fatal(err)
			}
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					log.Fatal(err)
				}
			}(tempFile.Name())
			if _, err = tempFile.Write([]byte(pair.data)); err != nil {
				log.Fatal(err)
			}

			if _, err := tempFile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			oldStdin := os.Stdin

			defer func() {
				os.Stdin = oldStdin
			}()

			os.Stdin = tempFile
			if fname != "" {
				fname = tempFile.Name()
			}

			res, err := ReadFile(fname)
			assert.Equal(t, pair.expected, res, "incorrect write to file\n got: %v\nexpected: %v", res, pair.expected)
			assert.Nil(t, err)
		})
	}
}

func TestReadFileNegative(t *testing.T) {
	tests := []struct {
		fName    string
		data     string
		expected []string
	}{
		{fName: "test.txt", data: "Hello\nWorld\n", expected: nil},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			tempFile, err := ioutil.TempFile("", "")
			if err != nil {
				log.Fatal(err)
			}
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					log.Fatal(err)
				}
			}(tempFile.Name())

			err = os.Chmod(tempFile.Name(), 0)
			if err != nil {
				log.Fatal(err)
			}

			res, err := ReadFile(tempFile.Name())
			assert.Equal(t, pair.expected, res,
				"incorrect write to file\ngot: %v\nexpected: %v", res, pair.expected)
			assert.NotNil(t, err)
		})
	}
}
