package ical

import (
	"fmt"
	"strings"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

// defines Properties
// https://tools.ietf.org/html/rfc5545#section-3.8

// Attachment is ATTACH
// https://tools.ietf.org/html/rfc5545#section-3.8.1.1
type Attachment struct {
	Parameters parameter.Container
	Value      interface{} // TODO limit Binary or URI
}

func (a *Attachment) SetAttachment(params parameter.Container, value interface{}) error {
	enc, encOK := params[parameter.TypeNameInlineEncoding]
	val, valOK := params[parameter.TypeNameValueType]
	if encOK && valOK {
		if len(enc) == 0 {
			return fmt.Errorf("")
		}
		if len(val) == 0 {
			return fmt.Errorf("")
		}
		encoding, ok := enc[0].(*parameter.InlineEncoding)
		if !ok || encoding.Type != parameter.InlineEncodingTypeBASE64 {
			return fmt.Errorf("")
		}
		valueType, ok := val[0].(*parameter.ValueType)
		if !ok || valueType.Value != "BINARY" {
			return fmt.Errorf("")
		}
		v, ok := value.(types.Binary)
		if !ok {
			return fmt.Errorf("invalid type %T", value)
		}
		a.Parameters = params
		a.Value = v
		return nil
	} else if encOK {
		return fmt.Errorf("%s and %s are must be true", parameter.TypeNameInlineEncoding, parameter.TypeNameValueType)
	} else if valOK {
		return fmt.Errorf("%s and %s are must be true", parameter.TypeNameInlineEncoding, parameter.TypeNameValueType)
	}
	if len(params[parameter.TypeNameFormatType]) > 1 {
		return fmt.Errorf("%s must be set only 1", parameter.TypeNameFormatType)
	}
	v, ok := value.(types.URI)
	if !ok {
		return fmt.Errorf("invalid type %T", value)
	}
	a.Parameters = params
	a.Value = v
	return nil
}

// Categories is CATEGORIES
// https://tools.ietf.org/html/rfc5545#section-3.8.1.2
type Categories struct {
	Parameters parameter.Container
	Values     []types.Text
}

// SetCategories updates CATEGORIES
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.2
func (c *Categories) SetCategories(params parameter.Container, values []types.Text) error {
	if len(values) == 0 {
		return ErrInputIsEmpty
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("%s must be set only 1", parameter.TypeNameLanguage)
	}
	c.Parameters = params
	c.Values = values
	return nil
}

// Class is CLASS, optional property for components
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.3
type Class struct {
	Parameter parameter.Container
	Value     types.Text
}

func (c *Class) SetClass(params parameter.Container, value types.Text) error {
	switch ClassType(value) {
	case PUBLIC, PRIVATE, CONFIDENTIAL:
		c.Parameter = params
		c.Value = value
		return nil
	default:
		if token.IsXName(string(value)) {
			c.Parameter = params
			c.Value = value
			return nil
		}
		return fmt.Errorf("invalid value: %s", value)
	}
}

// Comment is COMMENT
// https://tools.ietf.org/html/rfc5545#section-3.8.1.4
type Comment struct {
	Parameter parameter.Container
	Value     types.Text
}

// SetComment updates property value
func (c *Comment) SetComment(params parameter.Container, value types.Text) error {
	if len(value) == 0 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	c.Parameter = params
	c.Value = value
	return nil
}

// SetDescription updates DESCRIPTION
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.5
type Description struct {
	Parameter parameter.Container
	Value     types.Text
}

func (d *Description) SetDescription(params parameter.Container, value types.Text) error {
	if len(value) == 0 {
		return fmt.Errorf("input is nil")
	}
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	d.Parameter = params
	d.Value = value
	return nil
}

// Geo is GEO
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.6
type Geo struct {
	Parameter parameter.Container
	Latitude  types.Float
	Longitude types.Float
}

func (g *Geo) SetGeo(params parameter.Container, latitude, longitude types.Float) error {
	if latitude > 180 || latitude < -180 {
		return fmt.Errorf("")
	}
	if longitude > 180 || longitude < -180 {
		return fmt.Errorf("")
	}
	g.Parameter = params
	g.Latitude = latitude
	g.Longitude = longitude
	return nil
}

func (g *Geo) SetGeoWithText(params parameter.Container, value types.Text) error {
	var lat, log types.Float
	var err error
	v := strings.SplitN(string(value), ";", 2)
	if len(v) != 2 {
		return fmt.Errorf("input %s cannot split with ;", value)
	}
	lat, err = types.NewFloat(v[0])
	if err != nil {
		return fmt.Errorf("convert %s to float: %w", v[0], err)
	}
	log, err = types.NewFloat(v[1])
	if err != nil {
		return fmt.Errorf("convert %s to float: %w", v[1], err)
	}
	return g.SetGeo(params, lat, log)
}

// Locaiton is LOCATION
// location of component
// https://tools.ietf.org/html/rfc5545#section-3.8.1.7
type Location struct {
	Parameter parameter.Container
	Value     types.Text
}

func (l *Location) SetLocation(params parameter.Container, value types.Text) error {
	if len(value) == 0 {
		return fmt.Errorf("input is nil")
	}
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	l.Parameter = params
	l.Value = value
	return nil
}

// PercentComplete is PERCENT-COMPLETE
// property for ToDo
// https://tools.ietf.org/html/rfc5545#section-3.8.1.8
type PercentComplete struct {
	Parameter parameter.Container
	Value     types.Integer
}

func (pc *PercentComplete) SetPercentComplete(params parameter.Container, value types.Integer) error {
	if value > 100 || value < 0 {
		return fmt.Errorf("")
	}
	pc.Parameter = params
	pc.Value = value
	return nil
}

// Priority is PRIORITY
// used for ToDo or Event
// https://tools.ietf.org/html/rfc5545#section-3.8.1.9
type Priority struct {
	Parameter parameter.Container
	Value     types.Integer
}

func (p *Priority) SetPriority(params parameter.Container, value types.Integer) error {
	if value > 9 || value < 0 {
		return fmt.Errorf("")
	}
	p.Parameter = params
	p.Value = value
	return nil
}

// Resources is RESOURCES
// https://tools.ietf.org/html/rfc5545#section-3.8.1.10
type Resources struct {
	Parameter parameter.Container
	Values    []types.Text
}

func (r *Resources) SetResources(params parameter.Container, values []types.Text) error {
	if len(values) == 0 {
		return fmt.Errorf("input is nil")
	}
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	r.Parameter = params
	r.Values = values
	return nil
}

// Status is STATUS
// https://tools.ietf.org/html/rfc5545#section-3.8.1.11
type Status struct {
	Parameter parameter.Container
	Value     StatusType
}

func (s *Status) SetStatus(params parameter.Container, value types.Text, kind component.ComponentType) error {
	v := StatusType(value)
	switch kind {
	case component.ComponentTypeEvent:
		switch v {
		case StatusTypeTentative, StatusTypeConfirmed, StatusTypeCancelled:
			s.Parameter = params
			s.Value = v
			return nil
		default:
			return fmt.Errorf("")
		}
	case component.ComponentTypeTODO:
		switch v {
		case StatusTypeNeedsAction, StatusTypeCompleted, StatusTypeInProcess, StatusTypeCancelled:
			s.Parameter = params
			s.Value = v
			return nil
		default:
			return fmt.Errorf("")
		}
	case component.ComponentTypeJournal:
		switch v {
		case StatusTypeDraft, StatusTypeFinal, StatusTypeCancelled:
			s.Parameter = params
			s.Value = v
			return nil
		default:
			return fmt.Errorf("")
		}
	default:
		return fmt.Errorf("")
	}
}

// Summary is SUMMARY
// used for "VEVENT", "VTODO", "VJOURNAL", or "VALARM"
// https://tools.ietf.org/html/rfc5545#section-3.8.1.12
type Summary struct {
	Parameter parameter.Container
	Value     types.Text
}

func (s *Summary) SetSummary(params parameter.Container, value types.Text) error {
	if len(value) == 0 {
		return fmt.Errorf("input is nil")
	}
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	s.Parameter = params
	s.Value = value
	return nil
}
