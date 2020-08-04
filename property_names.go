package ical

type PropertyName string

const (
	// component block
	PropertyNameBegin PropertyName = "BEGIN"
	PropertyNameEnd   PropertyName = "END"

	// Alarm

	PropertyNameAction      PropertyName = "ACTION"
	PropertyNameRepeatCount PropertyName = "REPEAT"
	PropertyNameTrigger     PropertyName = "TRIGGER"

	// descriptive
	PropertyNameAttachment      PropertyName = "ATTACH"
	PropertyNameCategories      PropertyName = "CATEGORIES"
	PropertyNameClass           PropertyName = "CLASS"
	PropertyNameComment         PropertyName = "COMMENT"
	PropertyNameDescription     PropertyName = "DESCRIPTION"
	PropertyNameGeo             PropertyName = "GEO"
	PropertyNameLocaiton        PropertyName = "LOCATION"
	PropertyNamePercentComplete PropertyName = "PERCENT-COMPLETE"
	PropertyNamePriority        PropertyName = "PRIORITY"
	PropertyNameResources       PropertyName = "RESOURCES"
	PropertyNameStatus          PropertyName = "STATUS"
	PropertyNameSummary         PropertyName = "SUMMARY"

	// time
	PropertyNameDateTimeCompleted PropertyName = "COMPLETED"
	PropertyNameDateTimeEnd       PropertyName = "DTEND"
	PropertyNameDateTimeStart     PropertyName = "DTSTART"
	PropertyNameDuration          PropertyName = "DURATION"
	PropertyNameFreeBusyTime      PropertyName = "FREEBUSY"
	PropertyNameTimeTransparency  PropertyName = "TRANSP"

	// timezone
	PropertyNameTimezoneIdentifier PropertyName = "TZID"
	PropertyNameTimezoneName       PropertyName = "TZNAME"
	PropertyNameTimezoneOffsetFrom PropertyName = "TZOFFSETFROM"
	PropertyNameTimezoneOffsetTo   PropertyName = "TZOFFSETTO"
	PropertyNameTimezoneURL        PropertyName = "TZURL"

	// relationship
	PropertyNameAttendee     PropertyName = "ATTENDEE"
	PropertyNameOrganizer    PropertyName = "ORGANIZER"
	PropertyNameRecurrenceID PropertyName = "RECURRENCE-ID"
	PropertyNameRelatedTo    PropertyName = "RELATED-TO"
	PropertyNameURL          PropertyName = "URL"
	PropertyNameUID          PropertyName = "UID"

	// change management
	PropertyNameDateTimeStamp   PropertyName = "DTSTAMP"
	PropertyNameDateTimeCreated PropertyName = "CREATED"
	PropertyNameLastModified    PropertyName = "LAST-MODIFIED"
	PropertyNameSequenceNumber  PropertyName = "SEQUENCE"

	// recurrence
	PropertyNameExceptionDateTimes  PropertyName = "EXDATE"
	PropertyNameRecurrenceDateTimes PropertyName = "RDATE"
	PropertyNameRecurrenceRule      PropertyName = "RRULE"
)
