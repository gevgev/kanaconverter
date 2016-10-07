package main

import (
	"bytes"
	//"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	inputText := "私の名前はキコです\n"

	runKanaConverter(populateStdin(inputText))
	//fmt.Printf("Katakana: %s\n", out)
}

func runKanaConverter(populate_stdin_func func(io.WriteCloser)) {
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
	io.Copy(os.Stdout, stdout)
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

func populateStdin(str string) func(io.WriteCloser) {
	return func(stdin io.WriteCloser) {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewBufferString(str))
	}
}
