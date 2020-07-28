package parser

import (
	"fmt"

	"github.com/knsh14/ical/component"
	"github.com/morikuni/failure"
)

const (
	Invalid failure.StringCode = "Invalid"
)

type NoEndError component.ComponentType

func (t NoEndError) Error() string {
	return fmt.Sprintf("finished without END:%s", string(t))
}
