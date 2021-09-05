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

//priority := {
//	"+" : 1,
//	"-" : 1,
//	"/" : 2,
//	"*" : 2,
//}

type Node interface{}

type Stack struct {
	data []Node
}

func (s *Stack) NewStack() *Stack {
	s.data = []Node{}
	return s
}
func (s *Stack) Push(elem Node) {
	if s.data == nil {
		s.data = []Node{}
	}
	s.data = append(s.data, elem)
}
func (s *Stack) Pop() Node {
	elem := s.data[len(s.data)-1:]
	s.data = s.data[0 : len(s.data)-1]

	return elem
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

//func toPolishNotation(str string) (Stack, error) {
//	if len(str) == 0 {
//		return Stack{}, nil
//	}
//
//}
