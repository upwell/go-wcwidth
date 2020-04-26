package wcwidth

import (
	"fmt"
	"testing"
)

func TestStringWidth(t *testing.T) {
	input := "✨hh"
	fmt.Println(StringWidth(input))

	input = "你好"
	fmt.Println(StringWidth(input))

	input = "1.32↓"
	fmt.Println(StringWidth(input))

	input = "1.32↓⚡"
	fmt.Println(StringWidth(input))
}
