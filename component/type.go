package component

// Type is
// https://tools.ietf.org/html/rfc5545#section-8.3.1
type Type string

const (
	TypeCalendar Type = "VCALENDAR"
	TypeEvent    Type = "VEVENT"
	TypeTODO     Type = "VTODO"
	TypeJournal  Type = "VJOURNAL"
	TypeFreeBusy Type = "VFREEBUSY"
	TypeTimezone Type = "VTIMEZONE"
	TypeAlarm    Type = "VALARM"
	TypeStandard Type = "STANDARD"
	TypeDaylight Type = "DAYLIGHT"
)
