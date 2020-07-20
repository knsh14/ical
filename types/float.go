package types

type Float struct {
	Parameters map[string][]string
	Value      string
}

func NewFloat(param map[string][]string, value string) *Float {
	return &Float{
		Parameters: param,
		Value:      value,
	}
}
