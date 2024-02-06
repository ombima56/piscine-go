package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("File name missing")
	case 1:
		if args[0] == "quest8.txt" {
			if content, err := ioutil.ReadFile(args[0]); err == nil {
				fmt.Print(string(content))
			} else {
				fmt.Println(err)
			}
		}
	default:
		fmt.Println("Too many arguments")
	}
}
