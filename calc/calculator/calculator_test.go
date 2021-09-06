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
func TestInfixToPostfixPositive(t *testing.T) {
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
			tokens:   []string{"110", "+", "50", "+", "(", "4", "-", "2", "*", "5", ")", "-", "10", "/", "2"},
			expected: []string{"110", "50", "+", "4", "2", "5", "*", "-", "+", "10", "2", "/", "-"},
		},
		{
			tokens:   []string{"(", "1", "+", "2", ")", "-", "3"},
			expected: []string{"1", "2", "+", "3", "-"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.tokens, ""), func(t *testing.T) {
			res, err := InfixToPostfix(pair.tokens)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}
func TestInfixToPostfixNegative(t *testing.T) {
	var tests = []struct {
		tokens []string
	}{
		{
			tokens: []string{"110", "ops", "50"},
		},
		{
			tokens: []string{"bad string"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.tokens, ""), func(t *testing.T) {
			_, err := InfixToPostfix(pair.tokens)
			assert.NotNil(t, err)
		})
	}
}
func TestCalculatePositive(t *testing.T) {
	var tests = []struct {
		postfixExpr []string
		expected    float64
	}{
		{
			postfixExpr: []string{},
			expected:    0,
		},
		{
			postfixExpr: []string{"110", "50", "+"},
			expected:    160,
		},
		{
			postfixExpr: []string{"110", "50", "+", "4", "2", "5", "*", "-", "+", "10", "2", "/", "-"},
			expected:    149,
		},
		{
			postfixExpr: []string{"1", "2", "+", "3", "-"},
			expected:    0,
		},
		{
			postfixExpr: []string{"110", "50", "+", "2", "4", "-", "*"},
			expected:    -320,
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.postfixExpr, ""), func(t *testing.T) {
			res, err := Calculate(pair.postfixExpr)
			assert.Equal(t, pair.expected, res)
			assert.Nil(t, err)
		})
	}
}

func TestCalculateNegative(t *testing.T) {
	var tests = []struct {
		postfixExpr []string
	}{
		{
			postfixExpr: []string{"bad tokens"},
		},
		{
			postfixExpr: []string{"-", "-", "-"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.postfixExpr, ""), func(t *testing.T) {
			res, err := Calculate(pair.postfixExpr)
			assert.Equal(t, -1., res)
			assert.NotNil(t, err)
		})
	}
}
