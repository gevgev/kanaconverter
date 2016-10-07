package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	inputString := "inputFile.txt"

	cmd := exec.Command("kakasi", "-JK", "-HK", "-s", "-o utf8", "-i utf8", inputString)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Input: ", inputString)
	fmt.Printf("Katakana: %s\n", b)
}
