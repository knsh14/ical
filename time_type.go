package ical

import (
	"fmt"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

func NewTimeType(params parameter.Container, s string) (types.TimeValue, error) {
	var tz string
	tzs := params[parameter.TypeNameReferenceTimezone]
	if len(tzs) > 0 {
		tz = tzs[0].(*parameter.ReferenceTimezone).Value
	}
	dt, err := types.NewDateTime(s, tz)
	if err == nil {
		return dt, nil
	}
	t, err := types.NewDate(s)
	if err == nil {
		return t, nil
	}
	return nil, fmt.Errorf("%s cant convert DATE or DATE-TIME", s)
}
