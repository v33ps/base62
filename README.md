# base62

This package converts to and from base62. This was built for a url shortener.

## Example:
```
package main

import (
	"fmt"

	base62 "github.com/v33ps/base62"
)

func main() {
	fmt.Println("4444 in base62 is: ", base62.Encode(4444))
	fmt.Println("19G in base10 is: ", base62.Decode("19G"))
}


// -------------------
$ go run example/main.go
4444 in base62 is:  19G
19G in base10 is:  4444
```
# base62
