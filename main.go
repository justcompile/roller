package main

import (
	"fmt"
	"github.com/justcompile/roller/pkg/roller"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments found")
		os.Exit(1)
	}

	args := os.Args[1:]

	dice, err := roller.ParseArgs(args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range dice {

		fmt.Println(d.Roll())
	}
}
