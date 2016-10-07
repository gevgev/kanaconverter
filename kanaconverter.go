package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	inputText := "私の名前はキコです"
	out := runKanaConverter(populateStdin(inputText))
	fmt.Printf("Input: %s\nOutput: %s\n", inputText, out)
}

func runKanaConverter(populate_stdin_func func(io.WriteCloser)) string {
	args := []string{"-JK", "-HK", "-s", "-o", "utf8", "-i", "utf8"}

	cmd := exec.Command("kakasi", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	populate_stdin_func(stdin)
	b, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func populateStdin(str string) func(io.WriteCloser) {
	return func(stdin io.WriteCloser) {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewBufferString(str))
	}
}
