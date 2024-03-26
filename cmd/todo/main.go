package main

import (
	"fmt"
	"os"
	"strings"

	"rggo/interacting/todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
		// os.Args is like Bash args, the first argument is the file of execution
		case len(os.Args) == 1: // meaning no arguments passed
			for _, item := range *l {
				fmt.Println(item.Task)
			}
		default:
			item := strings.Join(os.Args[1:], " ")
			l.Add(item)

			if err := l.Save(todoFileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
	}
}