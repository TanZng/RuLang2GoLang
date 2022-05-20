package translator

import (
	"fmt"
)

func SyntaxError(str string, line int) Value {
	fmt.Printf("Ru SyntaxError: %s at Line %d\n", str, line)
	return Value{Value: nil}
}

func TypeError(str string, line int) Value {
	fmt.Printf("Ru TypeError: %s at Line %d\n", str, line)
	return Value{Value: nil}
}
