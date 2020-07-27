package component

// ComponentType is
// https://tools.ietf.org/html/rfc5545#section-8.3.1
type ComponentType string

const (
	ComponentTypeCalender ComponentType = "VCALENDER"
	ComponentTypeEvent    ComponentType = "VEVENT"
	ComponentTypeTODO     ComponentType = "VTODO"
	ComponentTypeJournal  ComponentType = "VJOURNAL"
	ComponentTypeFreeBusy ComponentType = "VFREEBUSY"
	ComponentTypeTimezone ComponentType = "VTIMEZONE"
	ComponentTypeAlarm    ComponentType = "VALARM"
	ComponentTypeStandard ComponentType = "STANDARD"
	ComponentTypeDaylight ComponentType = "DAYLIGHT"
)
