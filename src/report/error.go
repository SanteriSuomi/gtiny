package report

import "fmt"

type ErrorReporter interface {
	Error(line int, column int, length int, format string, arg ...any)
}

type PrintErrorReporter struct {
	OnError func(line int, column int, length int, message string)
}

func (p *PrintErrorReporter) Error(line int, column int, length int, format string, arg ...any) {
	message := fmt.Sprintf(format, arg...)
	fmt.Printf("[Line %d] [Position %d-%d] Error: %s\n", line, column, column+length, message)
	p.OnError(line, column, length, message)
}
