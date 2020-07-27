package parameter

type FreeBusyTimeTypeKind string

const (
	FreeBusyTimeTypeKindFree            FreeBusyTimeTypeKind = "FREE"
	FreeBusyTimeTypeKindBusy            FreeBusyTimeTypeKind = "BUSY"
	FreeBusyTimeTypeKindBusyUnavailable FreeBusyTimeTypeKind = "BUSY-UNAVAILABLE"
	FreeBusyTimeTypeKindBusyTentative   FreeBusyTimeTypeKind = "BUSY-TENTATIVE"
)
