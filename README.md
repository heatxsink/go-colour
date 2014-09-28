go-colour
=========

Make your terminal output colour-ized. :-)


Usage
-----

Take a look at the tests! In short this is how you use colour.

```go

import (
	"fmt"
	"github.com/heatxsink/colour"
)

func main() {
	fmt.Println(colour.Colourize("Hello World!", colour.Underscore + colour.Blink + colour.FgBlack + colour.BgYellow))
}

```