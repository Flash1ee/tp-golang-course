package read_write

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testCase struct {
	in       string
	flags    Flags
	fnameIn  string
	fnameOut string
	err      error
}

//@todo тесты, в которых позиционные аргументы не в конце
//var getNegativeFlagCases = []testCase{
//	{"-k", Flags{false, false, false, 0, 0, false}, "", "", nil},
//	{"-c -f -1", Flags{false, false, false, 0, 0, false}, "", "", nil},
//	{"-c -d", Flags{false, false, false, 0, 0, false}, "", "", nil},
//}

func TestGetFlagsCorrect(t *testing.T) {
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
		flags, err := GetFlags(os.Args[0], pair.args)
		assert.Equal(t, pair.conf, flags, "flags must be equal")
		assert.Nil(t, err)
	}
}
