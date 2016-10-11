package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

var allOptions bool

func main() {
	inputText := "私の名前はキコです"

	flagAllOptions := flag.Bool("p", false, "`All Spelling Options`")
	flagInputString := flag.String("i", inputText, "`Input string to convert`")
	flag.Parse()
	if flag.Parsed() {
		allOptions = *flagAllOptions
		inputText = *flagInputString
	} else {
		flag.Usage()
	}

	out := runKanaConverter(populateStdin(inputText))
	fmt.Printf("Input: %s\nOutput: %s\n", inputText, out)
}

func runKanaConverter(populate_stdin_func func(io.WriteCloser)) string {
	//args := []string{"-JK", "-HK", "-s", "-p", "-o", "utf8", "-i", "utf8"}
	args := []string{"-JK", "-HK", "-s", "-o", "utf8", "-i", "utf8"}

	if allOptions {
		args = append(args, "-p")
	}
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
