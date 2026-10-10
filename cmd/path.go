package cmd

import (
	"fmt"
	"io/fs"
	"os"
	// "path/filepath"
	"strings"
	"time"
)

// RunLs lists directory contents with optional -a and -l flags
func RunLs(args []string) {
	showAll := false
	longFormat := false
	var targetDirs []string


	// Parse flags and target directories
	for _, arg := range args {

		// checks if our first argument is a flag, and if it's not the only argument
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			if strings.Contains(arg, "a") {
				showAll = true
			}
			if strings.Contains(arg, "l") {
				longFormat = true
			}
		} else {
			targetDirs = append(targetDirs, arg)
		}
	}

	// Default to current directory if none specified
	if len(targetDirs) == 0 {
		targetDirs = append(targetDirs, ".")
	}

	for _, dir := range targetDirs {
		if len(targetDirs) > 1 {
			fmt.Printf("%s:\n", dir)
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ls: %v\n", err)
			continue
		}

		for _, entry := range entries {
			name := entry.Name()

			// Skip hidden files unless -a is enabled
			if !showAll && strings.HasPrefix(name, ".") {
				continue
			}

			if longFormat {
				printLongEntry(dir, entry)
			} else {
				fmt.Printf("%s  ", name)
			}
		}	
	}
}

func printLongEntry(dir string, entry fs.DirEntry) {
	info, err := entry.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: %v\n", err)
		return
	}

	mode := info.Mode().String()
	size := info.Size()
	modTime := info.ModTime().Format(time.Stamp)
	name := info.Name()

	// differentiating directories
	if entry.IsDir() {
		name += "/"
	}

	// mode-size-modTime-name
	fmt.Printf("%-11s %8d %s %s\n", mode, size, modTime, name)
}