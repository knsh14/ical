package types

import "fmt"

type TimeValue interface {
	fmt.Stringer
	timeValue()
}

type AttachmentValue interface {
	fmt.Stringer
	attachmentValue()
}

type RecurrenceDateTimeValue interface {
	fmt.Stringer
	recurrenceDateTimeValue()
}

type TriggerValue interface {
	fmt.Stringer
	triggerValue()
}
