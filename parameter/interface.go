package parameter

import "strings"

type Container map[TypeName][]Base

func (c Container) String() string {
	var v []string
	for _, bases := range c {
		for _, b := range bases {
			v = append(v, b.String())
		}
	}
	return strings.Join(v, ";")
}

type Base interface {
	implementParameter()
	String() string
}

func (c Container) GetTimezone() string {
	l, ok := c[TypeNameReferenceTimezone]
	if !ok {
		return ""
	}
	if len(l) != 1 {
		return ""
	}
	v, ok := l[0].(*ReferenceTimezone)
	if !ok {
		return ""
	}
	return v.Value
}
