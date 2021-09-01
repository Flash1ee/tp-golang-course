package uniq

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"uniq/read_write"
)

func TestGetCountStringsPositive(t *testing.T) {
	var tests = []struct {
		data     []string
		flags    read_write.Flags
		expected []read_write.UniqRes
	}{
		{data: []string{"I love music.", "I love music.",
			"I love music.", " ", "I love music of Kartik.",
			"I love music of Kartik.", "Thanks.", "I love music of Kartik.",
			"I love music of Kartik."},
			flags: read_write.Flags{},
			expected: []read_write.UniqRes{
				{Str: "I love music.", Cnt: 3}, {Str: " ", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2}, {Str: "Thanks.", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2},
			},
		},
		{
			data:     []string{},
			flags:    read_write.Flags{},
			expected: []read_write.UniqRes{},
		},
		{
			data: []string{"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",

				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik."},
			flags: read_write.Flags{RegisterSkipF: true},
			expected: []read_write.UniqRes{{Str: "I LOVE MUSIC.", Cnt: 3}, {Str: "I love MuSIC of Kartik.", Cnt: 2},
				{Str: "Thanks.", Cnt: 1}, {Str: "I love music of kartik.", Cnt: 2}},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.data, "\n"), func(t *testing.T) {
			res, err := GetCountStrings(pair.data, pair.flags)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}
