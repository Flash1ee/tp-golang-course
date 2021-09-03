package read_write

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGetFlagsPositive(t *testing.T) {
	var tests = []struct {
		args []string
		conf Flags
	}{
		{[]string{}, Flags{}},
		{[]string{"-c"}, Flags{true, false, false, 0, 0, false, "", ""}},
		{[]string{"-f", "10"}, Flags{false, false, false, 10, 0, false, "", ""}},
		{[]string{"-u"}, Flags{false, false, true, 0, 0, false, "", ""}},
		{[]string{"-f=10"}, Flags{false, false, false, 10, 0, false, "", ""}},
		{[]string{"in.txt"}, Flags{false, false, false, 0, 0, false, "in.txt", ""}},
		{[]string{"in.txt", "out.txt"}, Flags{false, false, false, 0, 0, false, "in.txt", "out.txt"}},
		{[]string{"-c", "-s", "10", "in.txt", "out.txt"}, Flags{true, false, false, 0, 10, false, "in.txt", "out.txt"}},
		{[]string{"-d", "-i", "-f", "10", "-s=20", "in.txt", "out.txt"}, Flags{false, true, false, 10, 20, true, "in.txt", "out.txt"}},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.args, " "), func(t *testing.T) {
			res, _, err := GetFlags(os.Args[0], pair.args)
			assert.Equal(t, pair.conf, res, "flags must be equal")
			assert.Nil(t, err)
		})
	}
}
func TestGetFlagsNegative(t *testing.T) {
	var tests = []struct {
		args   []string
		conf   Flags
		errstr string
	}{
		{[]string{"-k"}, Flags{}, UnknownFlagError.Error()},
		{[]string{"-d", "-u", "-c"}, Flags{}, TogetherArgs.Error()},
		{[]string{"-c", "-f", "-10"}, Flags{}, SkipNegative.Error()},
		{[]string{"in.txt", "-c", "-s", "10"}, Flags{}, IncorrectPosition.Error()},
		{[]string{"-c", "in.txt", "-s", "10"}, Flags{}, IncorrectPosition.Error()},
	}

	for _, pair := range tests {
		t.Run(strings.Join(pair.args, " "), func(t *testing.T) {
			res, _, err := GetFlags(os.Args[0], pair.args)
			assert.Equal(t, pair.conf, res, "conf got %+v, want %+v", pair.conf, res)
			assert.Equal(t, pair.errstr, err.Error())
			assert.NotNil(t, err, "err got %v, want nil", err)
		})
	}
}

func TestReadFilePositive(t *testing.T) {
	var tests = []struct {
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
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}
func TestReadFileNegative(t *testing.T) {
	var tests = []struct {
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
			assert.Equal(t, pair.expected, res)
			assert.NotNil(t, err)
		})
	}
}
