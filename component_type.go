package ical

// ComponentType is
// https://tools.ietf.org/html/rfc5545#section-8.3.1
type ComponentType string

const (
	ComponentTypeCalender ComponentType = "VCALENDER"
	ComponentTypeEvent                  = "VEVENT"
	ComponentTypeTODO                   = "VTODO"
	ComponentTypeJournal                = "VJOURNAL"
	ComponentTypeFreeBusy               = "VFREEBUSY"
	ComponentTypeTimezone               = "VTIMEZONE"
	ComponentTypeAlarm                  = "VALARM"
	ComponentTypeStandard               = "STANDARD"
	ComponentTypeDaylight               = "DAYLIGHT"
)
