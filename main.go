package main

import (
	"fmt"
	"github.com/abmercy035/unix_clone/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command> [args...]")
		fmt.Println("Available commands: echo, cat")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "echo":
		cmd.RunEcho(args)
	case "cat":
		cmd.RunCat(args)
	case "grep":
		cmd.RunGrep(args)
	case "wc":
		cmd.RunWc(args)
	case "ls":
		cmd.RunLs(args)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
}
