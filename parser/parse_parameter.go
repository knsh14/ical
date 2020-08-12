package parser

import (
	"fmt"

	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/parameter"
)

func (p *Parser) parseParameter(cl *contentline.ContentLine) (parameter.Container, error) {
	params := parameter.Container(map[parameter.TypeName][]parameter.Base{})
	for _, v := range cl.Parameters {
		switch t := parameter.TypeName(v.Name); t {
		case parameter.TypeNameAlternateTextRepresentation:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewAlternateTextRepresentation(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameCommonName:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p := parameter.NewCommonName(v.Values[0])
			params[t] = append(params[t], p)
		case parameter.TypeNameCalenderUserType:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewCalenderUserType(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameDelegator:
			p, err := parameter.NewDelegator(v.Values)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameDelegatee:
			p, err := parameter.NewDelegatee(v.Values)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameDirectoryEntry:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewDirectoryEntry(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameInlineEncoding:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewInlineEncoding(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameFormatType:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewFormatType(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameFreeBusyTimeType:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewFreeBusyTimeType(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameLanguage:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewLanguage(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameMembership:
			p, err := parameter.NewMembership(v.Values)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameParticipationStatus:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			// TODO get current component type
			p, err := parameter.NewParticipationStatus(v.Values[0], p.currentComponentType)
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameRecurrenceIDRange:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewRecurrenceIDRange(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameAlarmTriggerRelationship:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewAlarmTriggerRelationship(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameRelationshipType:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewRelationshipType(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameParticipationRole:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewParticipationRole(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameRSVP:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewRSVP(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameSentBy:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewSentBy(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameReferenceTimezone:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p, err := parameter.NewReferenceTimezone(v.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse %s: %w", t, err)
			}
			params[t] = append(params[t], p)
		case parameter.TypeNameValueType:
			if len(v.Values) != 1 {
				return nil, fmt.Errorf("value for %s must be 1, but %d", t, len(v.Values))
			}
			p := parameter.NewValueType(v.Values[0])
			params[t] = append(params[t], p)
		default:
			// if x-token
		}
	}
	return params, nil
}
