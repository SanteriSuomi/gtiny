package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SanteriSuomi/gtiny/src/error"
	"github.com/SanteriSuomi/gtiny/src/scanner"
)

func main() {
	hadError := false

	args := os.Args
	argsLen := len(args)
	if argsLen > 2 {
		fmt.Println("Usage: gtiny <source file> or gtiny for REPL mode")
		os.Exit(69)
	} else if argsLen == 2 {
		fmt.Println("Source mode")
		source, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(69)
		}
		scanner.RunSource(string(source), &error.PrintErrorReporter{})
	} else {
		fmt.Println("REPL mode")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(69)
			}
			scanner.RunPrompt(input, &error.PrintErrorReporter{})
			hadError = false
		}
	}

	if hadError {
		os.Exit(65)
	}
}
