package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	inputString := "inputFile.txt"

	out, err := exec.Command("kakasi", "-JK", "-HK", "-s", "-o utf8", "-i utf8", inputString).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input: ", inputString)
	fmt.Printf("Katakana: %s\n", out)
}
