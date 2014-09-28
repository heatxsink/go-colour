package colour

import (
	"fmt"
	"testing"
)

var (
	test_message = "Hello World"
)

func TestColour(t *testing.T) {
	fmt.Println("colour.Colourize")
	fmt.Println("\t", Colourize(test_message, FgWhite + BgRed))
	fmt.Println("\t\tTesting Reset.")
	fmt.Println("\t", Colourize(test_message, Blink + FgMagenta + BgGreen))
	fmt.Println("\t\tTesting Reset.")
	fmt.Println()
}

func TestDisable(t *testing.T) {
	fmt.Println("colour.Disable")
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			Disable(true)
		} else {
			Disable(false)
		}
		fmt.Println("\t", Colourize(test_message, Underscore + Blink + FgBlack + BgYellow))
	}
}