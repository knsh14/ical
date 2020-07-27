package parameter

type InlineEncodingType string

const (
	InlineEncodingType8BIT   InlineEncodingType = "8BIT"
	InlineEncodingTypeBASE64 InlineEncodingType = "BASE64"
	InlineEncodingTypeXName  InlineEncodingType = "X-NAME"
)
