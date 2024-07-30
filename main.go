package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SanteriSuomi/gtiny/src/report"
	"github.com/SanteriSuomi/gtiny/src/scanner"
)

func main() {
	args := os.Args
	argsLen := len(args)
	hadError := false
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
		tokens := scanner.RunSource(string(source), &report.PrintErrorReporter{})
		for _, token := range tokens {
			fmt.Println(token.Type)
		}
	} else {
		fmt.Println("REPL mode")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(69)
			}
			scanner.RunPrompt(input, &report.PrintErrorReporter{OnError: func(line int, column int, length int, message string) {
				hadError = true
			}})
			// If had error, don't execute code, but scan whole source.
		}
	}

	if hadError {
		os.Exit(65)
	}
}
