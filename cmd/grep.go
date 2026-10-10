package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// RunGrep searches for patterns in files or stdin
func RunGrep(args []string) {
	ignoreCase := false
	showLineNum := false
	invertMatch := false
	var pattern string
	var files []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			if strings.Contains(arg, "i") {
				ignoreCase = true
			}
			if strings.Contains(arg, "n") {
				showLineNum = true
			}
			if strings.Contains(arg, "v") {
				invertMatch = true
			}
		} else if pattern == "" {
			pattern = arg
		} else {
			files = append(files, arg)
		}
	}

	if pattern == "" {
		fmt.Fprintln(os.Stderr, "grep: missing pattern")
		return
	}

	if ignoreCase {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: invalid regex: %v\n", err)
		return
	}

	if len(files) == 0 {
		grepStream(os.Stdin, "", re, showLineNum, invertMatch, false)
		return
	}

	multipleFiles := len(files) > 1
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "grep: %v\n", err)
			continue
		}
		grepStream(file, filename, re, showLineNum, invertMatch, multipleFiles)
		file.Close()
	}
}

func grepStream(r io.Reader, filename string, re *regexp.Regexp, showLine, invert, showFile bool) {
	scanner := bufio.NewScanner(r)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.MatchString(line)

		if (matches && !invert) || (!matches && invert) {
			prefix := ""
			if showFile {
				prefix += filename + ":"
			}
			if showLine {
				prefix += fmt.Sprintf("%d:", lineNum)
			}
			fmt.Printf("%s%s\n", prefix, line)
		}
		lineNum++
	}
}