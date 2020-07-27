package parameter

// CalenderUserType is defined in https://www.iana.org/assignments/icalendar/icalendar.xhtml#calendar-user-types
type CalenderUserTypeKind string

const (
	CalenderUserTypeKindIndividual CalenderUserTypeKind = "INDIVIDUAL"
	CalenderUserTypeKindGroup      CalenderUserTypeKind = "GROUP"
	CalenderUserTypeKindResource   CalenderUserTypeKind = "RESOURCE"
	CalenderUserTypeKindRoom       CalenderUserTypeKind = "ROOM"
	CalenderUserTypeKindUnknown    CalenderUserTypeKind = "UNKNOWN"
	CalenderUserTypeKindXToken     CalenderUserTypeKind = "XToken"
)
