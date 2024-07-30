package scanner

import (
	"github.com/SanteriSuomi/gtiny/src/report"
	"github.com/SanteriSuomi/gtiny/src/scanner/internal"
	"github.com/SanteriSuomi/gtiny/src/token"
)

func RunSource(input string, reporter report.ErrorReporter) []token.Token {
	scanner := internal.Scanner{}
	return scanner.Run(input, reporter)
}

func RunPrompt(input string, reporter report.ErrorReporter) {
	scanner := internal.Scanner{}
	scanner.Run(input, reporter)
}
