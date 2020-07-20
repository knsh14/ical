package types

type Text struct {
	Parameters map[string][]string
	Value      string
}

func NewText(param map[string][]string, value string) *Text {
	return &Text{
		Parameters: param,
		Value:      value,
	}
}

type TextList struct {
	Parameters map[string][]string
	Values     []string
}

func NewTextList(param map[string][]string, values []string) *TextList {
	return &TextList{
		Parameters: param,
		Values:     values,
	}
}
