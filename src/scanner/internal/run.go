package internal

import (
	"fmt"

	"github.com/SanteriSuomi/gtiny/src/error"
)

type Token struct {
}

func Run(source string, reporter error.ErrorReporter) {
	tokens := scanTokens(source)
	for token := range tokens {
		fmt.Println(token)
	}
}

func scanTokens(source string) []Token {
	return []Token{}
}
