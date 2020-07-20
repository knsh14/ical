package ical

// Method is content-type
// list of methods is defined in https://www.iana.org/assignments/icalendar/icalendar.xhtml#methods
// used by https://tools.ietf.org/html/rfc5545#section-3.6
type Method string

const (
	PUBLISH        Method = "PUBLISH "
	REQUEST               = "REQUEST"
	REPLY                 = "REPLY"
	ADD                   = "ADD"
	CANCEL                = "CANCEL"
	REFRESH               = "REFRESH"
	COUNTER               = "COUNTER"
	DECLINECOUNTER        = "DECLINECOUNTER"
)

func isMethod(m string) bool {
	switch Method(m) {
	case PUBLISH,
		REQUEST,
		REPLY,
		ADD,
		CANCEL,
		REFRESH,
		COUNTER,
		DECLINECOUNTER:
		return true
	}
	return false
}
