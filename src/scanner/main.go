package scanner

import (
	"github.com/SanteriSuomi/gtiny/src/error"
	"github.com/SanteriSuomi/gtiny/src/scanner/internal"
)

func RunSource(source string, reporter error.ErrorReporter) {
	internal.Run(source, reporter)
}

func RunPrompt(prompt string, reporter error.ErrorReporter) {
	internal.Run(prompt, reporter)
}
