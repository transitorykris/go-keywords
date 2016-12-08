# Keywords

A little package for matching up keywords and IDs found in strings of text.

## Example usage

```go
package main

import (
	"fmt"

	"github.com/transitorykris/go-keywords"
)

func main() {
	// Create our new keywords
	kw := keywords.New()

	// Add some users
	kw.Add("hello", 1)
	kw.Add("Hello", 1)
	kw.Add("hello", 2)
	kw.Add("WORLD", 3)

	// See who matches this
	fmt.Println(kw.MatchedUsers("Hello, world!"))            // [1 2 3]
	fmt.Println(kw.MatchedUsers("This is an example World")) // [3]

	kw.Remove("world", 3)
	fmt.Println(kw.MatchedUsers("This is an example World")) // []
}
```
