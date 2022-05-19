package utils

import (
	"fmt"
	"os"
)

func SyntaxError(str string, line int) {
	fmt.Printf("Ru SyntaxError: %s at Line %d", str, line)

	os.Exit(1)
}

func TypeError(str string, line int) {
	fmt.Printf("Ru TypeError: %s at Line %d", str, line)
	os.Exit(1)
}
