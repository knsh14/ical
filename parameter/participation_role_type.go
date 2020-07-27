package parameter

type ParticipationRoleType string

const (
	ParticipationRoleTypeChair                ParticipationRoleType = "CHAIR"
	ParticipationRoleTypeRequestedParticipant ParticipationRoleType = "REQ-PARTICIPANT"
	ParticipationRoleTypeOptionalParticipant  ParticipationRoleType = "OPT-PARTICIPANT"
	ParticipationRoleTypeNonParticipant       ParticipationRoleType = "NON-PARTICIPANT"
	ParticipationRoleTypeXName                ParticipationRoleType = "X-NAME"
)
