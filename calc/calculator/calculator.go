package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func GetTokens(data string) ([]string, error) {
	res := make([]string, 0)
	var flagNum bool
	var curNum string

	for _, val := range data {
		cur := string(val)
		if cur != " " {
			if ok, _ := validTokens[cur]; ok {
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
					return nil, fmt.Errorf("error parse string\nUndefined token %s", cur)
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

	stack := NewStack()
	stack.Push("(")
	tokens = append(tokens, ")")

	var res []string

	for _, token := range tokens {
		switch token {
		case "(":
			stack.Push(token)
		case ")":
			for cur := stack.Pop().(string); cur != "("; cur = stack.Pop().(string) {
				res = append(res, cur)
			}
		default:
			if _, ok := validOperations[token]; ok {
				for !stack.isEmpty() && priority[stack.Peek().(string)] >= priority[token] {
					res = append(res, stack.Pop().(string))
				}
				stack.Push(token)
			} else if _, ok := strconv.ParseFloat(token, 64); ok == nil {
				res = append(res, token)
			} else {
				return nil, fmt.Errorf("incorrect sequence token %s in expression %s",
					token, strings.Join(tokens, ""))
			}
		}
	}

	return res, nil
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
		} else if ok, _ := validOperations[token]; ok {
			second, okSecond := stack.Pop().(float64)
			first, okFirst := stack.Pop().(float64)
			if !okSecond || !okFirst {
				return -1, fmt.Errorf("error sequence of tokens %s", tokens)
			}
			stack.Push(actions[token](first, second))
		} else {
			return -1, fmt.Errorf("incorrect token %s", token)
		}
	}

	return stack.Pop().(float64), nil
}
