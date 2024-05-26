package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1, err1 := Atoi(os.Args[1])
	operator := os.Args[2]
	value2, err2 := Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		return
	}
	if (value1 >= 9223372036854775807 || value1 <= -9223372036854775808) ||
		(value2 >= 9223372036854775807 || value2 <= -9223372036854775808) ||
		(string(value1) >= "9223372036854775807" || string(value1) <= "-9223372036854775808") ||
		(string(value2) >= "9223372036854775807" || string(value2) <= "-9223372036854775808"){
		return
	}

	var result int
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			noDivisionByZero()
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			noModuloByZero()
			return
		}
		result = value1 % value2
	default:
		return
	}
	Itoa(result)
}

func noDivisionByZero() {
	printError("No division by 0")
}

func noModuloByZero() {
	printError("No modulo by 0")
}

func printError(msg string) {
	for _, char := range msg {
		os.Stdout.WriteString(string(char))
	}
	os.Stdout.WriteString("\n")
}

func Atoi(s string) (int, error) {
	sign := 1
	var number int

	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
		} else if i == 0 && ch == '+' {
			sign = 1
		} else if ch < '0' || ch > '9' {
			return 0, &InvalidDigitError{ch}
		} else {
			number = number*10 + int(ch-'0')
		}
	}
	return sign * number, nil
}

type InvalidDigitError struct {
	Ch rune
}

func (e *InvalidDigitError) Error() string {
	return "invalid digit: " + string(e.Ch)
}

func Itoa(n int) {
	if n == 0 {
		os.Stdout.WriteString("0")
	}
	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	for _, num := range digits {
		os.Stdout.WriteString(string(num))
	}
	os.Stdout.WriteString("\n")
}
