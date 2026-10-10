package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// echo prints all arguments separated by a space
func RunEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

// cat reads files or stdin, with optional line numbering (-n)
func RunCat(args []string) {
	numberLines := false
	var files []string

	// Parse flags and file paths
	for _, arg := range args {
		if arg == "-n" {
			numberLines = true
		} else {
			files = append(files, arg)
		}
	}

	// If no files passed, read from terminal / stdin
	if len(files) == 0 {
		printContent(os.Stdin, numberLines)
		return
	}

	// Read each file in order
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %v\n", err)
			continue
		}
		printContent(file, numberLines)
		file.Close()
	}
}

func printContent(r io.Reader, numberLines bool) {
	if !numberLines {
		io.Copy(os.Stdout, r)
		return
	}

	scanner := bufio.NewScanner(r)
	line := 1
	for scanner.Scan() {
		fmt.Printf("%6d  %s\n", line, scanner.Text())
		line++
	}
}