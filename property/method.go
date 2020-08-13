package property

// Method is content-type
// list of methods is defined in https://www.iana.org/assignments/icalendar/icalendar.xhtml#methods
// used by https://tools.ietf.org/html/rfc5545#section-3.6
type MethodType string

const (
	MethodTypePublish        MethodType = "PUBLISH"
	MethodTypeRequest        MethodType = "REQUEST"
	MethodTypeReply          MethodType = "REPLY"
	MethodTypeAdd            MethodType = "ADD"
	MethodTypeCancel         MethodType = "CANCEL"
	MethodTypeRefresh        MethodType = "REFRESH"
	MethodTypeCounter        MethodType = "COUNTER"
	MethodTypeDeclinecounter MethodType = "DECLINECOUNTER"
)

func isMethod(m string) bool {
	switch MethodType(m) {
	case MethodTypePublish,
		MethodTypeRequest,
		MethodTypeReply,
		MethodTypeAdd,
		MethodTypeCancel,
		MethodTypeRefresh,
		MethodTypeCounter,
		MethodTypeDeclinecounter:
		return true
	}
	return false
}
