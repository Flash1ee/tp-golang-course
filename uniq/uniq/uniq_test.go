package uniq

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"uniq/read_write"
)

func TestGetUniqStringsPositive(t *testing.T) {
	var tests = []struct {
		data     []string
		flags    read_write.Flags
		expected []read_write.UniqRes
	}{
		{data: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
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
			data: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",

				"I love MuSIC of Kartik.",
				"I love music of kartik.",
				"Thanks.",
				"I love music of kartik.",
				"I love MuSIC of Kartik.",
			},
			flags: read_write.Flags{RegisterSkipF: true},
			expected: []read_write.UniqRes{{Str: "I LOVE MUSIC.", Cnt: 3}, {Str: "I love MuSIC of Kartik.", Cnt: 2},
				{Str: "Thanks.", Cnt: 1}, {Str: "I love music of kartik.", Cnt: 2}},
		},
		{
			data: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			flags: read_write.Flags{RepeatF: true},
			expected: []read_write.UniqRes{{Str: "I love music.", Cnt: 3}, {Str: "", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2}, {Str: "Thanks.", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2}},
		},
		{
			data: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			flags: read_write.Flags{CntSkipWordsF: 1},
			expected: []read_write.UniqRes{{Str: "We love music.", Cnt: 3}, {Str: "", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2}, {Str: "Thanks.", Cnt: 1}},
		},
		{
			data: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			flags: read_write.Flags{CntSkipCharsF: 1},
			expected: []read_write.UniqRes{{Str: "We love music.", Cnt: 1}, {Str: "I love music.", Cnt: 1},
				{Str: "They love music.", Cnt: 1}, {Str: "", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 1}, {Str: "We love music of Kartik.", Cnt: 1}, {Str: "Thanks.", Cnt: 1}},
		},
		{
			data: []string{
				"We love music.",
				"I love music.",
				"I love music.",
				"They love music.",
			},
			flags: read_write.Flags{CntSkipCharsF: 0},
			expected: []read_write.UniqRes{{Str: "We love music.", Cnt: 1}, {Str: "I love music.", Cnt: 2},
				{Str: "They love music.", Cnt: 1}},
		},
		{
			data: []string{
				"We love music.",
				"I love music.",
				"I love music.",
				"They love music.",
			},
			flags: read_write.Flags{CntSkipWordsF: 0},
			expected: []read_write.UniqRes{{Str: "We love music.", Cnt: 1}, {Str: "I love music.", Cnt: 2},
				{Str: "They love music.", Cnt: 1}},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.data, "\n"), func(t *testing.T) {
			res, err := GetUniqStrings(pair.data, pair.flags)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}

func TestGetUniqStringsNegative(t *testing.T) {
	var tests = []struct {
		data  []string
		flags read_write.Flags
		err   error
	}{
		{
			data:  []string{"I love music", "T love music"},
			flags: read_write.Flags{CntSkipCharsF: -1},
			err:   IncorrectArgs,
		},
		{
			data:  []string{"I love music", "We love music"},
			flags: read_write.Flags{CntSkipWordsF: -1},
			err:   IncorrectArgs,
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.data, "\n"), func(t *testing.T) {
			res, err := GetUniqStrings(pair.data, pair.flags)
			assert.Equal(t, res, []read_write.UniqRes{})
			assert.Equal(t, pair.err, err)
			assert.NotNil(t, err)
		})
	}
}
func TestSkipWordsNegative(t *testing.T) {
	var tests = []struct {
		prev    string
		cur     string
		cnt     int
		newPrev string
		newCur  string
		err     error
	}{
		{
			prev:    "I love music",
			cur:     "You love music",
			cnt:     0,
			newPrev: "I love music",
			newCur:  "You love music",
			err:     nil,
		},
		{
			prev:    "I love music",
			cur:     "You love music",
			cnt:     -1,
			newPrev: "I love music",
			newCur:  "You love music",
			err:     IncorrectArgs,
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join([]string{pair.prev, "\n", pair.cur}, "\n"), func(t *testing.T) {
			var prev, cur string
			prev, cur, err := SkipWords(pair.prev, pair.cur, pair.cnt)
			assert.Equal(t, prev, pair.newPrev)
			assert.Equal(t, cur, pair.newCur)
			if err != nil {
				assert.Equal(t, pair.err.Error(), err.Error())
			}
		})
	}
}
func TestSkipCharsNegative(t *testing.T) {
	var tests = []struct {
		prev    string
		cur     string
		cnt     int
		newPrev string
		newCur  string
		err     error
	}{
		{
			prev:    "I love music",
			cur:     "You love music",
			cnt:     0,
			newPrev: "I love music",
			newCur:  "You love music",
			err:     nil,
		},
		{
			prev:    "I love music",
			cur:     "You love music",
			cnt:     -1,
			newPrev: "I love music",
			newCur:  "You love music",
			err:     IncorrectArgs,
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join([]string{pair.prev, "\n", pair.cur}, "\n"), func(t *testing.T) {
			var prev, cur string
			prev, cur, err := SkipChars(pair.prev, pair.cur, pair.cnt)
			assert.Equal(t, prev, pair.newPrev)
			assert.Equal(t, cur, pair.newCur)
			if err != nil {
				assert.Equal(t, pair.err.Error(), err.Error())
			}
		})
	}
}
func TestUniq(t *testing.T) {
	var tests = []struct {
		data     []string
		flags    read_write.Flags
		expected []read_write.UniqRes
	}{
		{
			data: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				" ",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
			flags: read_write.Flags{},
			expected: []read_write.UniqRes{
				{Str: "I love music.", Cnt: 3}, {Str: " ", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2}, {Str: "Thanks.", Cnt: 1},
				{Str: "I love music of Kartik.", Cnt: 2},
			},
		},
		{
			expected: []read_write.UniqRes{},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.data, " "), func(t *testing.T) {
			res, err := Uniq(pair.data, pair.flags)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}
