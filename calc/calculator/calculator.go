package calculator

import (
	"errors"
	"fmt"
	"strconv"
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
				_, err := strconv.Atoi(cur)
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
func InfixToPostfix(tokens []string) []string {
	if len(tokens) == 0 {
		return []string{}
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
		} else if token == "+" || token == "-" || token == "/" || token == "*" {
			for !stack.isEmpty() && priority[stack.Peek().(string)] >= priority[token] {
				res = append(res, stack.Pop().(string))
			}
			stack.Push(token)
		} else {
			res = append(res, token)
		}
	}

	return res
}

//func Calculate(tokens []string) (int, error) {
//	if len(tokens) == 0 {
//		return 0, nil
//	}
//	actions := map[string]func(a int, b int) int{
//		"+": func(a int, b int) int {
//			return a + b
//		},
//		"-": func(a int, b int) int {
//			return a - b
//		},
//		"/": func(a int, b int) int {
//			return a / b
//		},
//		"*": func(a int, b int) int {
//			return a * b
//		},
//	}
//	stack := NewStack()
//	res := 0
//
//	for _, cur := range tokens {
//
//		if val, err := strconv.Atoi(cur); err == nil {
//			operands.Push(val)
//
//		} else if _, exist := priority[cur]; exist {
//			last := fmt.Sprintf("%s", operations.Peek())
//
//			if operations.isEmpty() {
//				operations.Push(cur)
//			} else if last != "(" && last != ")" {
//				if priority[last] < priority[cur] {
//					operations.Push(cur)
//				} else {
//					for priority[last] >= priority[cur] && last != "(" && last != ")" {
//
//						a := operands.Pop()
//						b := operands.Pop()
//						if a == nil || b == nil {
//							return -1, errors.New(fmt.Sprintf("Incorrect calculate string %s", tokens))
//						}
//						first, second := a.(int), b.(int)
//						operands.Push(actions[cur](first, second))
//						operations.Pop()
//						if !operations.isEmpty() {
//							last = operations.Peek().(string)
//						}
//					}
//
//				}
//			}
//		} else if cur == "(" {
//			operations.Push(cur)
//
//		} else if cur == ")" {
//
//		} else {
//			return -1, errors.New(fmt.Sprintf("Incorrect token %s", cur))
//		}
//	}
//
//	return res, nil
//
//}
