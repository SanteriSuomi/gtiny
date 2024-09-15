package scanner

import (
	"github.com/SanteriSuomi/gtiny/src/entities"
	"github.com/SanteriSuomi/gtiny/src/report"
	"github.com/SanteriSuomi/gtiny/src/scanner/internal"
)

func RunSource(input string, rep report.ErrorReporter) []entities.Token {
	return internal.RunSource(input, rep)
}

// Wrapper for prompt scanning
// Creates an instance of a scanner for the whole lifecycle of the REPL
type PromptRunner struct {
	scanner internal.Scanner
}

func ConstructRun(rep report.ErrorReporter) PromptRunner {
	return PromptRunner{scanner: internal.ConstructPromptScanner(rep)}
}

func (r *PromptRunner) RunPrompt(input string) {
	r.scanner.RunPrompt(input)
}

func (r *PromptRunner) EndPrompting() []entities.Token {
	return r.scanner.PromptingResult()
}
