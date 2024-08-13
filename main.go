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
	hadError := false
	onError := func(line int, column int, length int, message string) {
		hadError = true
	}
	if len(args) > 2 {
		fmt.Println("Usage: gtiny <source file> or gtiny for REPL mode")
		os.Exit(69)
	} else if len(args) == 2 {
		// Process the whole source file in one go
		fmt.Println("Source mode")
		source, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(69)
		}
		tokens := scanner.RunSource(string(source), &report.PrintErrorReporter{OnError: onError})
		for _, token := range tokens {
			fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
		}
	} else {
		// Process lines given one by one
		fmt.Println("REPL mode")
		reader := bufio.NewReader(os.Stdin)
		run := scanner.ConstructRun(&report.PrintErrorReporter{OnError: onError})
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(69)
			}
			if input == "exit\n" || input == "exit\r\n" {
				break
			}
			run.RunPrompt(input)
		}
		for _, token := range run.EndPrompting() {
			fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
		}
	}

	if hadError {
		os.Exit(65)
	}
}
