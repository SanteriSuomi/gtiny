package report

import "fmt"

type ErrorReporter interface {
	Error(line int, column int, length int, format string, a ...any)
}

type PrintErrorReporter struct {
	OnError func(line int, column int, length int, message string)
}

func (p *PrintErrorReporter) Error(line int, column int, length int, format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("[Line %d] [Position %d-%d] Error: %s\n", line, column, column+length, message)
	p.OnError(line, column, length, message)
}
