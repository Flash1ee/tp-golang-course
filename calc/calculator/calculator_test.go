package calculator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetTokensPositive(t *testing.T) {
	var tests = []struct {
		data     string
		expected []string
	}{
		{
			data:     "(1+2)-3",
			expected: []string{"(", "1", "+", "2", ")", "-", "3"},
		},
		{
			data:     "(1+2)*3",
			expected: []string{"(", "1", "+", "2", ")", "*", "3"},
		},
		{
			data:     "",
			expected: []string{},
		},
		{
			data:     "  (1     +2 )",
			expected: []string{"(", "1", "+", "2", ")"},
		},
		{
			data:     "(10+20)*2/3",
			expected: []string{"(", "10", "+", "20", ")", "*", "2", "/", "3"},
		},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			res, err := GetTokens(pair.data)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}
func TestGetTokensNegative(t *testing.T) {
	var tests = []struct {
		data     string
		expected []string
	}{
		{
			data:     "(1+2)-3s",
			expected: nil,
		},
		{
			data:     "golang",
			expected: nil,
		},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			res, err := GetTokens(pair.data)
			assert.Equal(t, pair.expected, res)
			assert.NotNil(t, err)
		})
	}
}
func TestInfixToPostfix(t *testing.T) {
	var tests = []struct {
		tokens   []string
		expected []string
	}{
		{
			tokens:   []string{},
			expected: []string{},
		},
		{
			tokens:   []string{"110", "+", "50"},
			expected: []string{"110", "50", "+"},
		},
		{
			tokens:   []string{"110", "+", "50", "+", "(", "4", "-", "2", "*", "5", ")", "-", "10", "+", "40"},
			expected: []string{"110", "50", "+", "4", "2", "5", "*", "-", "+", "10", "-", "40", "+"},
		},
		{
			tokens:   []string{"(", "1", "+", "2", ")", "-", "3"},
			expected: []string{"1", "2", "+", "3", "-"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.tokens, ""), func(t *testing.T) {
			res := InfixToPostfix(pair.tokens)
			assert.Equal(t, pair.expected, res)
		})
	}
}
