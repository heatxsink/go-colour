go-colour
=========

Make your terminal output colour-ized. :-)


Usage
-----

Please take a look at the tests! Below is a quick example of how you could use the colour package.

```go

import (
	"fmt"
	"github.com/heatxsink/colour"
)

func main() {
	fmt.Println(colour.Colourize("Hello World!", colour.Underscore + colour.Blink + colour.FgBlack + colour.BgYellow))
}

```