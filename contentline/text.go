package contentline

import (
	"fmt"

	"github.com/knsh14/ical/types"
)

func NewText(cl *ContentLine) (*types.Text, error) {
	if cl == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if len(cl.Values) != 1 {
		return nil, fmt.Errorf("value must be 1, but %d", len(cl.Values))
	}
	var param map[string][]string
	for _, p := range cl.Parameters {
		param[p.Name] = p.Values
	}
	return types.NewText(param, cl.Values[0]), nil
}

func NewTextList(cl *ContentLine) (*types.TextList, error) {
	if cl == nil {
		return nil, fmt.Errorf("input is nil")
	}
	var param map[string][]string
	for _, p := range cl.Parameters {
		param[p.Name] = p.Values
	}
	return types.NewTextList(param, cl.Values), nil
}
