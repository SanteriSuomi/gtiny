package gtiny

import (
	"fmt"
	"os"

	"github.com/SanteriSuomi/gtiny/scanner"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("Usage: gtiny <source file>")
		os.Exit(1)
	}
	source, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner.Main(string(source))
}
