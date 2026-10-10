package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type wcCounts struct {
	lines int
	words int
	bytes int
}

// RunWc counts lines, words, and bytes
func RunWc(args []string) {
	showLines, showWords, showBytes := false, false, false
	var files []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			if strings.Contains(arg, "l") {
				showLines = true
			}
			if strings.Contains(arg, "w") {
				showWords = true
			}
			if strings.Contains(arg, "c") {
				showBytes = true
			}
		} else {
			files = append(files, arg)
		}
	}

// by behaviour of wc all, it shows both lines, words and bytes by defautl.
	if !showLines && !showWords && !showBytes {
		showLines, showWords, showBytes = true, true, true
	}

	// Read from stdin if no files specified
	if len(files) == 0 {
		counts, err := countReader(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", err)
			return
		}
		printWcResult(counts, "", showLines, showWords, showBytes)
		return
	}

	total := wcCounts{}
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", err)
			continue
		}

		counts, err := countReader(file)
		file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", err)
			continue
		}

		printWcResult(counts, filename, showLines, showWords, showBytes)
		total.lines += counts.lines
		total.words += counts.words
		total.bytes += counts.bytes
	}

	if len(files) > 1 {
		printWcResult(total, "total", showLines, showWords, showBytes)
	}
}

func countReader(r io.Reader) (wcCounts, error) {
	scanner := bufio.NewScanner(r)
	counts := wcCounts{}

	for scanner.Scan() {
		counts.lines++
		line := scanner.Text()
		counts.bytes += len(scanner.Bytes()) + 1 // +1 for the newline byte
		counts.words += len(strings.Fields(line))
	}

	return counts, scanner.Err()
}

func printWcResult(c wcCounts, name string, l, w, b bool) {
	var parts []string
	if l {
		parts = append(parts, fmt.Sprintf("%7d", c.lines))
	}
	if w {
		parts = append(parts, fmt.Sprintf("%7d", c.words))
	}
	if b {
		parts = append(parts, fmt.Sprintf("%7d", c.bytes))
	}
	if name != "" {
		parts = append(parts, " "+name)
	}
	fmt.Println(strings.Join(parts, " "))
}