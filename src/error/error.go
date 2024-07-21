package error

import "fmt"

type ErrorReporter interface {
	Error(line int, column int, length int, message string)
}

type PrintErrorReporter struct{}

func (*PrintErrorReporter) Error(line int, column int, length int, message string) {
	fmt.Printf("[line %d] [position %d-%d] Error: %s\n", line, column, column+length, message)
}
