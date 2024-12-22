package calculator

import (
	"errors"
	"unicode"
)

// Calc ...
func Calc(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	output, err := shuntingYard(tokens)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(output)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func tokenize(expression string) ([]string, error) {
	var tokens []string
	var number string

	for _, r := range expression {
		switch {
		case unicode.IsDigit(r) || r == '.':
			number += string(r)
		case r == '+' || r == '-' || r == '*' || r == '/' || r == '(' || r == ')':
			if len(number) > 0 {
				tokens = append(tokens, number)
				number = ""
			}
			tokens = append(tokens, string(r))
		case unicode.IsSpace(r):
			continue
		default:
			return nil, errors.New("invalid character in expression")
		}
	}

	if len(number) > 0 {
		tokens = append(tokens, number)
	}

	return tokens, nil
}

func shuntingYard(tokens []string) ([]string, error) {
	var output []string
	var operators []string

	precedence := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2,
	}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(operators) > 0 {
				top := operators[len(operators)-1]
				if top == "(" || precedence[top] < precedence[token] {
					break
				}
				output = append(output, top)
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		case "(":
			operators = append(operators, token)
		case ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
		default:
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}

		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func evaluateRPN(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res float64
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}

				res = a / b
			}

			stack = append(stack, res)
		default:
			value, err := stringToFloat(token)

			if err != nil {
				return 0, err
			}

			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func stringToFloat(s string) (float64, error) {
	var result float64
	var decimalPlace float64 = 1
	var isDecimal bool

	for _, r := range s {
		if r == '.' {
			if isDecimal {
				return 0, errors.New("invalid number format")
			}
			isDecimal = true
			continue
		}
		
		if !unicode.IsDigit(r) {
			return 0, errors.New("invalid number format")
		}

		digit := float64(r - '0')
		if isDecimal {
			decimalPlace *= 0.1
			result += digit * decimalPlace
		} else {
			result = result*10 + digit
		}
	}

	return result, nil
}
