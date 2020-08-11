package parameter

type Container map[TypeName][]Base

type Base interface {
	implementParameter()
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
