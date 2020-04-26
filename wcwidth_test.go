package wcwidth_test

import (
	"fmt"
	"testing"
	"wcwidth"
)

func TestStringWidth(t *testing.T) {
	input := "✨hh"
	fmt.Println(wcwidth.StringWidth(input))

	input = "你好"
	fmt.Println(wcwidth.StringWidth(input))

	input = "1.32↓"
	fmt.Println(wcwidth.StringWidth(input))

	input = "1.32↓⚡"
	fmt.Println(wcwidth.StringWidth(input))
}
