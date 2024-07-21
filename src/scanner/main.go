package scanner

import (
	"github.com/SanteriSuomi/gtiny/src/error"
	"github.com/SanteriSuomi/gtiny/src/scanner/internal"
)

func RunSource(input string, reporter error.ErrorReporter) {
	scanner := internal.Scanner{}
	scanner.Run(input, reporter)
}

func RunPrompt(input string, reporter error.ErrorReporter) {
	scanner := internal.Scanner{}
	scanner.Run(input, reporter)
}
