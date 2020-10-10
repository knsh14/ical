package property

// Name is list of contents line name
type Name string

const (
	// component block

	NameBegin Name = "BEGIN"
	NameEnd   Name = "END"

	// Alarm

	NameAction      Name = "ACTION"
	NameRepeatCount Name = "REPEAT"
	NameTrigger     Name = "TRIGGER"

	// descriptive

	NameAttachment      Name = "ATTACH"
	NameCategories      Name = "CATEGORIES"
	NameClass           Name = "CLASS"
	NameComment         Name = "COMMENT"
	NameDescription     Name = "DESCRIPTION"
	NameGeo             Name = "GEO"
	NameLocation        Name = "LOCATION"
	NamePercentComplete Name = "PERCENT-COMPLETE"
	NamePriority        Name = "PRIORITY"
	NameResources       Name = "RESOURCES"
	NameStatus          Name = "STATUS"
	NameSummary         Name = "SUMMARY"

	// time

	NameDateTimeCompleted Name = "COMPLETED"
	NameDateTimeEnd       Name = "DTEND"
	NameDateTimeDue       Name = "DUE"
	NameDateTimeStart     Name = "DTSTART"
	NameDuration          Name = "DURATION"
	NameFreeBusyTime      Name = "FREEBUSY"
	NameTimeTransparency  Name = "TRANSP"

	// timezone

	NameTimezoneIdentifier Name = "TZID"
	NameTimezoneName       Name = "TZNAME"
	NameTimezoneOffsetFrom Name = "TZOFFSETFROM"
	NameTimezoneOffsetTo   Name = "TZOFFSETTO"
	NameTimezoneURL        Name = "TZURL"

	// relationship

	NameAttendee     Name = "ATTENDEE"
	NameContact      Name = "CONTACT"
	NameOrganizer    Name = "ORGANIZER"
	NameRecurrenceID Name = "RECURRENCE-ID"
	NameRelatedTo    Name = "RELATED-TO"
	NameURL          Name = "URL"
	NameUID          Name = "UID"

	// change management

	NameDateTimeStamp   Name = "DTSTAMP"
	NameDateTimeCreated Name = "CREATED"
	NameLastModified    Name = "LAST-MODIFIED"
	NameSequenceNumber  Name = "SEQUENCE"

	// recurrence

	NameExceptionDateTimes  Name = "EXDATE"
	NameRecurrenceDateTimes Name = "RDATE"
	NameRecurrenceRule      Name = "RRULE"

	// Miscellaneous

	NameRequestStatus Name = "REQUEST-STATUS"

	// Calender component

	NameCalScale Name = "CALSCALE"
	NameMethod   Name = "METHOD"
	NameProdID   Name = "PRODID"
	NameVersion  Name = "VERSION"
)
