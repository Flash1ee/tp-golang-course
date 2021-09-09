package calculator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokensPositive(t *testing.T) {
	var tests = []struct {
		description string
		data        string
		expected    []string
	}{
		{
			description: "test which split expression with () and operatons +/-",
			data:        "(1+2)-3",
			expected:    []string{"(", "1", "+", "2", ")", "-", "3"},
		},
		{
			description: "test which split expression with () and combination of +/*",
			data:        "(1+2)*3",
			expected:    []string{"(", "1", "+", "2", ")", "*", "3"},
		},
		{
			description: "test with empty expression",
			data:        "",
			expected:    []string{},
		},
		{
			description: "test with whitespaces",
			data:        "  (1     +2 )",
			expected:    []string{"(", "1", "+", "2", ")"},
		},
		{
			description: "test with combinations of (), [+, *, /]",
			data:        "(10+20)*2/3",
			expected:    []string{"(", "10", "+", "20", ")", "*", "2", "/", "3"},
		},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			res, err := GetTokens(pair.data)
			assert.Equal(t, pair.expected, res, pair.description+"\ngot: %v\nexpected: %v\n", res, pair.expected)
			assert.Nil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}
func TestGetTokensNegative(t *testing.T) {
	var tests = []struct {
		description string
		data        string
		expected    []string
	}{
		{
			description: "test which have incorrect symbol s in expression",
			data:        "(1+2)-3s",
			expected:    nil,
		},
		{
			description: "not valid expression - string",
			data:        "golang",
			expected:    nil,
		},
	}
	for _, pair := range tests {
		t.Run(pair.data, func(t *testing.T) {
			res, err := GetTokens(pair.data)
			assert.Equal(t, pair.expected, res, pair.description+"\ngot: %v\nexpected: %v\n", res, pair.expected)
			assert.NotNil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}
func TestInfixToPostfixPositive(t *testing.T) {
	var tests = []struct {
		description string
		tokens      []string
		expected    []string
	}{
		{
			description: "test with empty input",
			tokens:      []string{},
			expected:    []string{},
		},
		{
			description: "test infix to postfix notation - only + operation",
			tokens:      []string{"110", "+", "50"},
			expected:    []string{"110", "50", "+"},
		},
		{
			description: "test infix to postfix with combinations of operators [+, -, *, /, (, )]",
			tokens:      []string{"110", "+", "50", "+", "(", "4", "-", "2", "*", "5", ")", "-", "10", "/", "2"},
			expected:    []string{"110", "50", "+", "4", "2", "5", "*", "-", "+", "10", "2", "/", "-"},
		},
		{
			description: "test with ( in start of expression",
			tokens:      []string{"(", "1", "+", "2", ")", "-", "3"},
			expected:    []string{"1", "2", "+", "3", "-"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.tokens, ""), func(t *testing.T) {
			res, err := InfixToPostfix(pair.tokens)
			assert.Equal(t, pair.expected, res, pair.description+"\ngot: %v\nexpected: %v\n", res, pair.expected)
			assert.Nil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}
func TestInfixToPostfixNegative(t *testing.T) {
	var tests = []struct {
		description string
		tokens      []string
	}{
		{
			description: "test with incorrect sequence of tokens - one incorrect",
			tokens:      []string{"110", "ops", "50"},
		},
		{
			description: "all input tokens are incorrect",
			tokens:      []string{"bad string"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.tokens, ""), func(t *testing.T) {
			_, err := InfixToPostfix(pair.tokens)
			assert.NotNil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}
func TestCalculatePositive(t *testing.T) {
	var tests = []struct {
		description string
		postfixExpr []string
		expected    float64
	}{
		{
			description: "test with empty input",
			postfixExpr: []string{},
			expected:    0,
		},
		{
			description: "test with correct postfix notation and one operation +",
			postfixExpr: []string{"110", "50", "+"},
			expected:    160,
		},
		{
			description: "test with correct postfix notation and combination of operations",
			postfixExpr: []string{"110", "50", "+", "4", "2", "5", "*", "-", "+", "10", "2", "/", "-"},
			expected:    149,
		},
		{
			description: "test which evaluate expression and got 0 in result",
			postfixExpr: []string{"1", "2", "+", "3", "-"},
			expected:    0,
		},
		{
			description: "test which evaluate expression and got negative number",
			postfixExpr: []string{"110", "50", "+", "2", "4", "-", "*"},
			expected:    -320,
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.postfixExpr, ""), func(t *testing.T) {
			res, err := Calculate(pair.postfixExpr)
			assert.Equal(t, pair.expected, res, pair.description+"\ngot: %v\nexpected: %v\n", res, pair.expected)
			assert.Nil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}

func TestCalculateNegative(t *testing.T) {
	var tests = []struct {
		description string
		postfixExpr []string
	}{
		{
			description: "token is not expression",
			postfixExpr: []string{"bad tokens"},
		},
		{
			description: "sequence have not numbers, only operators",
			postfixExpr: []string{"-", "-", "-"},
		},
	}
	for _, pair := range tests {
		t.Run(strings.Join(pair.postfixExpr, ""), func(t *testing.T) {
			res, err := Calculate(pair.postfixExpr)
			assert.Equal(t, -1., res, pair.description+"\ngot: %v\nexpected: %v\n", res, -1.)
			assert.NotNil(t, err, pair.description+"\ngot: %v\nexpected: %v\n", err, nil)
		})
	}
}
