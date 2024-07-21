package error

import "fmt"

type ErrorReporter interface {
	error(row int, columnStart int, columnEnd int, message string)
}

type PrintErrorReporter struct{}

func (*PrintErrorReporter) error(line int, column int, length int, message string) {
	fmt.Printf("[line %d] [position %d-%d] Error: %s\n", line, column, column+length, message)
}
