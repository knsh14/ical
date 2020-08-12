ical
---

[![PkgGoDev](https://pkg.go.dev/badge/github.com/knsh14/ical)](https://pkg.go.dev/github.com/knsh14/ical)

parse iCal file based on [RFC 5545]( https://tools.ietf.org/html/rfc5545 )

# Usage

```go
package main

import (
	"fmt"

	"github.com/knsh14/ical/parser"
)

func main() {
	calendar, err := parser.Parse("~/holiday.ics")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", calendar)
}
```


