# powershell
golang powershell

``` go
package main

import (
	"fmt"

	ps "github.com/Tobotobo/powershell"
)

func main() {
	out, _ := ps.Execute("1+2")
	fmt.Println(out)
}
```