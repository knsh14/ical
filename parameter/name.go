package parameter

type TypeName string

const (
	TypeNameAlternateTextRepresentation TypeName = "ALTREP"
	TypeNameCommonName                  TypeName = "CN"
	TypeNameCalenderUserType            TypeName = "CUTYPE"
	TypeNameDelegator                   TypeName = "DELEGATED-FROM"
	TypeNameDelegatee                   TypeName = "DELEGATED-TO"
	TypeNameDirectoryEntry              TypeName = "DIR"
	TypeNameInlineEncoding              TypeName = "ENCODING"
	TypeNameFormatType                  TypeName = "FMTTYPE"
	TypeNameFreeBusyTimeType            TypeName = "FBTYPE"
	TypeNameLanguage                    TypeName = "LANGUAGE"
	TypeNameMembership                  TypeName = "MEMBER"
	TypeNameParticipationStatus         TypeName = "PARTSTAT"
	TypeNameRecurrenceIDRange           TypeName = "RANGE"
	TypeNameAlarmTriggerRelationship    TypeName = "RELATED"
	TypeNameRelationshipType            TypeName = "RELTYPE"
	TypeNameParticipationRole           TypeName = "ROLE"
	TypeNameRSVP                        TypeName = "RSVP"
	TypeNameSentBy                      TypeName = "SENT-BY"
	TypeNameReferenceTimezone           TypeName = "TZID"
	TypeNameValueType                   TypeName = "VALUE"
)
