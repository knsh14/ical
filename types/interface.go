package types

type TimeValue interface {
	timeValue()
}

type AttachmentValue interface {
	attachmentValue()
}

type RecurrenceDateTimeValue interface {
	recurrenceDateTimeValue()
}

type TriggerValue interface {
	triggerValue()
}
