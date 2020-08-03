package ical

type StatusType string

const (
	StatusTypeTentative   StatusType = "TENTATIVE"
	StatusTypeConfirmed   StatusType = "CONFIRMED"
	StatusTypeCancelled   StatusType = "CANCELLED"
	StatusTypeNeedsAction StatusType = "NEEDS-ACTION"
	StatusTypeCompleted   StatusType = "COMPLETED"
	StatusTypeInProcess   StatusType = "IN-PROCESS"
	StatusTypeDraft       StatusType = "DRAFT"
	StatusTypeFinal       StatusType = "FINAL"
)
