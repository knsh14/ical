package parameter

// CalendarUserTypeKind is types of calendar
// https://www.iana.org/assignments/icalendar/icalendar.xhtml#calendar-user-types
type CalendarUserTypeKind string

const (
	CalendarUserTypeKindIndividual CalendarUserTypeKind = "INDIVIDUAL"
	CalendarUserTypeKindGroup      CalendarUserTypeKind = "GROUP"
	CalendarUserTypeKindResource   CalendarUserTypeKind = "RESOURCE"
	CalendarUserTypeKindRoom       CalendarUserTypeKind = "ROOM"
	CalendarUserTypeKindUnknown    CalendarUserTypeKind = "UNKNOWN"
	CalendarUserTypeKindXToken     CalendarUserTypeKind = "XToken"
)
