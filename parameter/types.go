package parameter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/mime"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
	"golang.org/x/text/language"
)

func NewAlternateTextRepresentation(value string) (*AlternateTextRepresentation, error) {
	uri, err := types.NewURI(value)
	if err != nil {
		return nil, fmt.Errorf("parse input to uri: %w", err)
	}
	return &AlternateTextRepresentation{
		URI: uri,
	}, nil
}

// AlternateTextRepresentation is defined in https://tools.ietf.org/html/rfc5545#section-3.2.1
type AlternateTextRepresentation struct {
	URI types.URI
}

func (a *AlternateTextRepresentation) implementParameter() {}

func NewCommonName(value string) *CommonName {
	return &CommonName{
		Value: types.NewText(value),
	}
}

// CommonName is defined in https://tools.ietf.org/html/rfc5545#section-3.2.2
type CommonName struct {
	Value types.Text
}

func (cn *CommonName) implementParameter() {}

func NewCalenderUserType(value string) (*CalenderUserType, error) {
	switch v := CalendarUserTypeKind(value); v {
	case CalendarUserTypeKindIndividual,
		CalendarUserTypeKindGroup,
		CalendarUserTypeKindResource,
		CalendarUserTypeKindRoom,
		CalendarUserTypeKindUnknown:
		return &CalenderUserType{Type: v}, nil
	default:
		if token.IsXName(value) {
			return &CalenderUserType{Type: CalendarUserTypeKindXToken, Value: value}, nil
		}
		return nil, fmt.Errorf("undefined CalenderUserType %s", value)
	}
}

// CalenderUserType is defined in https://tools.ietf.org/html/rfc5545#section-3.2.3
type CalenderUserType struct {
	Type  CalendarUserTypeKind
	Value string
}

func (cut *CalenderUserType) implementParameter() {}

func NewDelegator(values []string) (*Delegator, error) {
	var addresses []types.CalenderUserAddress
	for _, value := range values {
		//v, err := strconv.Unquote(value)
		//if err != nil {
		//	return nil, fmt.Errorf("unquote input[%s]: %w", value, err)
		//}
		a, err := types.NewCalenderUserAddress(value)
		if err != nil {
			return nil, fmt.Errorf("convert value[%s] to CALENDER-USER-ADDRESS: %w", value, err)
		}
		addresses = append(addresses, a)
	}
	return &Delegator{Addresses: addresses}, nil
}

type Delegator struct {
	Addresses []types.CalenderUserAddress
}

func (d *Delegator) implementParameter() {}

func NewDelegatee(values []string) (*Delegatee, error) {
	var addresses []types.CalenderUserAddress
	for _, value := range values {
		a, err := types.NewCalenderUserAddress(value)
		if err != nil {
			return nil, fmt.Errorf("convert value[%s] to CALENDER-USER-ADDRESS: %w", value, err)
		}
		addresses = append(addresses, a)
	}
	return &Delegatee{Addresses: addresses}, nil
}

type Delegatee struct {
	Addresses []types.CalenderUserAddress
}

func (d *Delegatee) implementParameter() {}

func NewDirectoryEntry(value string) (*DirectoryEntry, error) {
	// v, err := strconv.Unquote(value)
	// if err != nil {
	// 	return nil, fmt.Errorf("unquote input: %w", err)
	// }
	uri, err := types.NewURI(value)
	if err != nil {
		return nil, fmt.Errorf("parse input to uri: %w", err)
	}
	return &DirectoryEntry{URI: uri}, nil
}

type DirectoryEntry struct {
	URI types.URI
}

func (de *DirectoryEntry) implementParameter() {}

func NewInlineEncoding(value string) (*InlineEncoding, error) {
	switch v := InlineEncodingType(value); v {
	case InlineEncodingType8BIT, InlineEncodingTypeBASE64:
		return &InlineEncoding{Type: v, Value: value}, nil
	default:
		if token.IsXName(value) {
			return &InlineEncoding{Type: InlineEncodingTypeXName, Value: value}, nil
		}
		return nil, fmt.Errorf("undefined InlineEncodingType %s", value)
	}
}

type InlineEncoding struct {
	Type  InlineEncodingType
	Value string // for x-name
}

func (ie *InlineEncoding) implementParameter() {}

func NewFormatType(value string) (*FormatType, error) {
	if mime.IsMIMEType(value) {
		return &FormatType{Value: types.NewText(value)}, nil
	}
	return nil, fmt.Errorf("invalid format type %s", value)
}

type FormatType struct {
	Value types.Text
}

func (ft *FormatType) implementParameter() {}

func NewFreeBusyTimeType(value string) (*FreeBusyTimeType, error) {
	switch v := FreeBusyTimeTypeKind(value); v {
	case FreeBusyTimeTypeKindFree,
		FreeBusyTimeTypeKindBusy,
		FreeBusyTimeTypeKindBusyUnavailable,
		FreeBusyTimeTypeKindBusyTentative:
		return &FreeBusyTimeType{Type: v}, nil
	default:
		if token.IsXName(value) {
			return &FreeBusyTimeType{Type: FreeBusyTimeTypeKindXName, Value: value}, nil
		}
		return nil, fmt.Errorf("invalid FreeBusyTimeType %s", v)
	}
}

// FreeBusyTimeType is definded in https://tools.ietf.org/html/rfc5545#section-3.2.9
type FreeBusyTimeType struct {
	Type  FreeBusyTimeTypeKind
	Value string // for X-NAME
}

func (fbtt *FreeBusyTimeType) implementParameter() {}

func NewLanguage(value string) (*Language, error) {
	tag, err := language.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("parse language tag: %w", err)
	}
	return &Language{Tag: tag}, nil
}

type Language struct {
	Tag language.Tag
}

func (l *Language) implementParameter() {}

func NewMembership(values []string) (*Membership, error) {
	var l []types.CalenderUserAddress
	for _, value := range values {
		uri, err := types.NewCalenderUserAddress(value)
		if err != nil {
			return nil, fmt.Errorf("parse %s to uri: %w", value, err)
		}
		l = append(l, uri)
	}
	return &Membership{URIs: l}, nil
}

type Membership struct {
	URIs []types.CalenderUserAddress
}

func (m *Membership) implementParameter() {}

func NewParticipationStatus(value string, kind component.Type) (*ParticipationStatus, error) {
	var list map[ParticipationStatusType]struct{}
	switch kind {
	case component.TypeEvent:
		list = map[ParticipationStatusType]struct{}{
			ParticipationStatusTypeNeedsAction: {},
			ParticipationStatusTypeAccepted:    {},
			ParticipationStatusTypeDeclined:    {},
			ParticipationStatusTypeDelegated:   {},
			ParticipationStatusTypeXToken:      {},
		}
	case component.TypeTODO:
		list = map[ParticipationStatusType]struct{}{
			ParticipationStatusTypeNeedsAction: {},
			ParticipationStatusTypeAccepted:    {},
			ParticipationStatusTypeDeclined:    {},
			ParticipationStatusTypeTentative:   {},
			ParticipationStatusTypeDelegated:   {},
			ParticipationStatusTypeCompleted:   {},
			ParticipationStatusTypeInProcess:   {},
			ParticipationStatusTypeXToken:      {},
		}
	case component.TypeJournal:
		list = map[ParticipationStatusType]struct{}{
			ParticipationStatusTypeNeedsAction: {},
			ParticipationStatusTypeAccepted:    {},
			ParticipationStatusTypeDeclined:    {},
			ParticipationStatusTypeXToken:      {},
		}
	default:
		return nil, fmt.Errorf("invalid kind type %s, must be VEVENT, VTODO or VJOURNAL", kind)
	}
	t := ParticipationStatusType(value)
	if token.IsXName(value) {
		t = ParticipationStatusTypeXToken
	}
	if _, ok := list[t]; ok {
		return &ParticipationStatus{
			Kind:  kind,
			Type:  t,
			Value: value,
		}, nil
	}
	return nil, fmt.Errorf("%s is not for %s", value, kind)
}

type ParticipationStatus struct {
	Kind  component.Type
	Type  ParticipationStatusType
	Value string
}

func (ps *ParticipationStatus) implementParameter() {}

func NewRecurrenceIDRange(value string) (*RecurrenceIDRange, error) {
	if strings.ToUpper(value) != "THISANDFUTURE" {
		return nil, fmt.Errorf("value must be THISANDFUTURE, but %s", value)
	}
	return &RecurrenceIDRange{}, nil
}

type RecurrenceIDRange struct {
}

func (ridr *RecurrenceIDRange) implementParameter() {}

func NewAlarmTriggerRelationship(value string) (*AlarmTriggerRelationship, error) {
	if value == "START" {
		return &AlarmTriggerRelationship{IsStart: true}, nil
	}
	if value == "END" {
		return &AlarmTriggerRelationship{IsStart: false}, nil
	}
	return nil, fmt.Errorf("value must be START or END, but %s", value)
}

type AlarmTriggerRelationship struct {
	IsStart bool // value must be START or END, so true means Start
}

func (atr *AlarmTriggerRelationship) implementParameter() {}

func NewRelationshipType(value string) (*RelationshipType, error) {
	switch t := RelationshipTypeKind(value); t {
	case RelationshipTypeKindParent,
		RelationshipTypeKindChild,
		RelationshipTypeKindSibling:
		return &RelationshipType{Type: t, Value: value}, nil
	default:
		if token.IsXName(value) {
			return &RelationshipType{Type: RelationshipTypeKindXName, Value: value}, nil
		}
		return nil, fmt.Errorf("invalid RelationshipType %s", value)
	}
}

type RelationshipType struct {
	Type  RelationshipTypeKind
	Value string // for x-name
}

func (rt *RelationshipType) implementParameter() {}

func NewParticipationRole(value string) (*ParticipationRole, error) {
	switch t := ParticipationRoleType(value); t {
	case ParticipationRoleTypeChair,
		ParticipationRoleTypeRequestedParticipant,
		ParticipationRoleTypeOptionalParticipant,
		ParticipationRoleTypeNonParticipant:
		return &ParticipationRole{Type: t, Value: value}, nil
	default:
		if token.IsXName(value) {
			return &ParticipationRole{Type: ParticipationRoleTypeXName, Value: value}, nil
		}
		return nil, fmt.Errorf("invalid ParticipationRoleType %s", value)
	}
}

type ParticipationRole struct {
	Type  ParticipationRoleType
	Value string // for x-token
}

func (pr *ParticipationRole) implementParameter() {}

func NewRSVP(value string) (*RSVP, error) {
	b, err := types.NewBoolean(value)
	if err != nil {
		return nil, fmt.Errorf("convert value to boolean: %w", err)
	}
	return &RSVP{Value: b}, nil
}

type RSVP struct {
	Value types.Boolean
}

func (rsvp *RSVP) implementParameter() {}

func NewSentBy(value string) (*SentBy, error) {
	v, err := strconv.Unquote(value)
	if err != nil {
		return nil, fmt.Errorf("unquote input: %w", err)
	}
	uri, err := types.NewCalenderUserAddress(v)
	if err != nil {
		return nil, fmt.Errorf("parse input to CalenderUserAddress: %w", err)
	}
	return &SentBy{Address: uri}, nil
}

type SentBy struct {
	Address types.CalenderUserAddress
}

func (sb *SentBy) implementParameter() {}

func NewReferenceTimezone(value string) (*ReferenceTimezone, error) {
	for _, v := range value {
		r := rune(v)
		if unicode.IsControl(r) {
			return nil, fmt.Errorf("control character in %s", value)
		}
	}
	if strings.ContainsAny(value, "\",;:") {
		return nil, fmt.Errorf("not param safe character in %s", value)
	}
	return &ReferenceTimezone{Value: value}, nil
}

type ReferenceTimezone struct {
	Value string
}

func (rtz *ReferenceTimezone) implementParameter() {}

func NewValueType(value string) *ValueType {
	return &ValueType{Value: value}
}

type ValueType struct {
	Value string
}

func (vt *ValueType) implementParameter() {}

func NewXParam(param string, values []string) *XParam {
	return &XParam{
		Parameter: param,
		Value:     values,
	}
}

type XParam struct {
	Parameter string
	Value     []string
}

func (xp *XParam) implementParameter() {}
