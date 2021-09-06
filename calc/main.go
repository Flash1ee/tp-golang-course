package main

import (
	"calc/calculator"
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	args := os.Args[1:]
	if len(args) != 1 {
		strErr := fmt.Sprintf("many arguments were passed\nExpected: 1\nReceived: %d\n", len(args))
		fmt.Println(errors.New(strErr))
		os.Exit(1)
	}
	calcString := args[0]
	tokens, err := calculator.GetTokens(calcString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	postfixExpression, err := calculator.InfixToPostfix(tokens)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	res, err := calculator.Calculate(postfixExpression)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(res)
}
