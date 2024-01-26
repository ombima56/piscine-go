package main

import "fmt"

func BasicAtoi(s string) int {
	var answer int
	for _, char := range s {
		answer = answer*10 + int(char-'0')
	}
	return answer
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
