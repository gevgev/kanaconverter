package main

import (
	"fmt"
	"os"
)

func main() {
	inputText := "私の名前はキコです"
	args := os.Args[1:]
	if len(args) > 0 {
		inputText = args[0]
	}

	fmt.Printf("Input: %s\nOutput: %x\n", inputText, inputText)
	fmt.Printf("%v\n", getHex(inputText))
}

func getHex(inString string) []string {
	var list []string
	for i := 0; i < len(inString); i++ {
		list = append(list, fmt.Sprintf("%x", inString[i]))
	}
	return list
}
