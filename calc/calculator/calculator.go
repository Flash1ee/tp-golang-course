package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

type Node interface{}

type Stack struct {
	data []Node
}

func NewStack() *Stack {
	return &Stack{data: []Node{}}
}
func (s *Stack) Push(elem Node) {
	if s.data == nil {
		s.data = []Node{}
	}
	s.data = append(s.data, elem)
}
func (s *Stack) Pop() Node {
	if s.isEmpty() {
		return nil
	}
	elem := s.data[len(s.data)-1:]
	s.data = s.data[0 : len(s.data)-1]

	return elem[0]
}
func (s *Stack) Peek() Node {
	if s.isEmpty() {
		return nil
	}
	elem := s.data[len(s.data)-1:]
	return elem[0]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func GetTokens(data string) ([]string, error) {
	res := make([]string, 0, 0)
	var flagNum bool
	var curNum string

	validTokens := []string{"(", ")", "-", "+", "/", "*"}

	for _, val := range data {
		cur := string(val)
		if cur != " " {
			if Contains(validTokens, cur) {
				if flagNum {
					res = append(res, curNum)
					flagNum = false
					curNum = ""
				}
				res = append(res, cur)
			} else {
				_, err := strconv.ParseFloat(cur, 64)
				if err == nil {
					curNum += cur
					if !flagNum {
						flagNum = true
					}
				} else {
					return nil, errors.New(fmt.Sprintf("Error parse string\nUndefined token %s", cur))
				}
			}
		}

	}
	if curNum != "" {
		res = append(res, curNum)
	}
	return res, nil

}
func InfixToPostfix(tokens []string) ([]string, error) {
	if len(tokens) == 0 {
		return []string{}, nil
	}
	priority := map[string]int{
		"+": 2,
		"-": 2,
		"/": 3,
		"*": 4,
		"(": 1,
	}

	stack := NewStack()
	stack.Push("(")
	tokens = append(tokens, ")")

	var res []string

	for _, token := range tokens {
		if token == "(" {
			stack.Push(token)
		} else if token == ")" {
			for cur := stack.Pop().(string); cur != "("; cur = stack.Pop().(string) {
				res = append(res, cur)
			}
		} else if isOperation(token) {
			for !stack.isEmpty() && priority[stack.Peek().(string)] >= priority[token] {
				res = append(res, stack.Pop().(string))
			}
			stack.Push(token)
		} else if _, ok := strconv.ParseFloat(token, 64); ok == nil {
			res = append(res, token)
		} else {
			return nil, errors.New(fmt.Sprintf("Incorrect sequence token %s in expression %s",
				token, strings.Join(tokens, "")))
		}
	}

	return res, nil
}
func isOperation(token string) bool {
	validOperations := []string{
		"+",
		"-",
		"/",
		"*",
		"(",
	}
	return Contains(validOperations, token)
}
func Calculate(tokens []string) (float64, error) {
	if len(tokens) == 0 {
		return 0, nil
	}
	actions := map[string]func(a float64, b float64) float64{
		"+": func(a float64, b float64) float64 {
			return a + b
		},
		"-": func(a float64, b float64) float64 {
			return a - b
		},
		"/": func(a float64, b float64) float64 {
			return a / b
		},
		"*": func(a float64, b float64) float64 {
			return a * b
		},
	}
	stack := NewStack()

	for _, token := range tokens {
		if val, err := strconv.ParseFloat(token, 64); err == nil {
			stack.Push(val)

		} else if isOperation(token) {
			second, okSecond := stack.Pop().(float64)
			first, okFirst := stack.Pop().(float64)
			if !okSecond || !okFirst {
				return -1, errors.New(fmt.Sprintf("Error sequence of tokens %s", tokens))
			}
			stack.Push(actions[token](first, second))
		} else {
			return -1, errors.New(fmt.Sprintf("Incorrect token %s", token))
		}
	}

	return stack.Pop().(float64), nil

}
