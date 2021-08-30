package read_write

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGetFlagsPositive(t *testing.T) {
	var tests = []struct {
		args []string
		conf Flags
	}{
		{[]string{}, Flags{false, false, false, 0, 0, false, "", ""}},
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
		{[]string{"-k"}, Flags{false, false, false, 0, 0, false, "", ""}, UnknownFlagError.Error()},
		{[]string{"-d", "-u", "-c"}, Flags{false, false, false, 0, 0, false, "", ""}, TogetherArgs.Error()},
		{[]string{"in.txt", "-c", "-s", "10"}, Flags{false, false, false, 0, 0, false, "", ""}, IncorrectPosition.Error()},
		{[]string{"-c", "in.txt", "-s", "10"}, Flags{false, false, false, 0, 0, false, "", ""}, IncorrectPosition.Error()},
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
