package types

type PeriodType string

const (
	PeriodTypeInvalid  PeriodType = ""
	PeriodTypeExplicit PeriodType = "Explicit"
	PeriodTypeStart    PeriodType = "Start"
)
