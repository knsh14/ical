package types

type TimeType interface {
	isTime()
}

type Attachmentable interface {
	attachmentable()
}

type RecurrenceDateTime interface {
	implementRecurrenceDateTime()
}
