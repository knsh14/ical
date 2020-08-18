ical
---

[![PkgGoDev](https://pkg.go.dev/badge/github.com/knsh14/ical)](https://pkg.go.dev/github.com/knsh14/ical)

parse iCal file based on [RFC 5545]( https://tools.ietf.org/html/rfc5545 )

# Example
https://play.golang.org/p/SDQl0Cwc67J
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/parser"
)

func main() {
	res, err := http.Get("https://www.google.com/calendar/ical/japanese__ja%40holiday.calendar.google.com/public/basic.ics")
	if err != nil {
		log.Fatal(err)
	}
	cal, err := parser.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cal.Components {
		if e, ok := c.(*ical.Event); ok {
			fmt.Println(e.Summary.Value)
		}
	}
}
```


