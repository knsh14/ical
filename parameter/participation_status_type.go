package parameter

type ParticipationStatusType string

const (
	ParticipationStatusTypeNeedsAction ParticipationStatusType = "NEEDS-ACTION"
	ParticipationStatusTypeAccepted    ParticipationStatusType = "ACCEPTED "
	ParticipationStatusTypeDeclined    ParticipationStatusType = "DECLINED"
	ParticipationStatusTypeTentative   ParticipationStatusType = "TENTATIVE"
	ParticipationStatusTypeDelegated   ParticipationStatusType = "DELEGATED"
	ParticipationStatusTypeCompleted   ParticipationStatusType = "COMPLETED"
	ParticipationStatusTypeInProcess   ParticipationStatusType = "IN-PROCESS"
	ParticipationStatusTypeXToken      ParticipationStatusType = "XToken"
)
