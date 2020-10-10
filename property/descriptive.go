package property

import (
	"fmt"
	"io"
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
	Parameter parameter.Container
	Value     types.AttachmentValue
}

func (a *Attachment) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameAttachment, a.Parameter.String(), s)
	return nil
}

func (a *Attachment) Validate() error {
}

func NewAttachmentValue(params parameter.Container, s string) (types.AttachmentValue, error) {
	if checkAttachmentIsBinary(params) {
		b, err := types.NewBinary(s)
		if err != nil {
			return nil, fmt.Errorf("convert %s to Binary for Attachment: %w", s, err)
		}
		return b, nil
	}
	uri, err := types.NewURI(s)
	if err != nil {
		return nil, fmt.Errorf("convert %s to URL for Attachment: %w", s, err)
	}
	return uri, nil
}

func checkAttachmentIsBinary(params parameter.Container) bool {
	if len(params) != 2 {
		return false
	}
	enc, encOK := params[parameter.TypeNameInlineEncoding]
	if !encOK {
		return false
	}
	encoding, ok := enc[0].(*parameter.InlineEncoding)
	if !(ok && encoding.Type == parameter.InlineEncodingTypeBASE64) {
		return false
	}
	val, valOK := params[parameter.TypeNameValueType]
	if !valOK {
		return false
	}
	valueType, ok := val[0].(*parameter.ValueType)
	if !ok || valueType.Value != "BINARY" {
		return false
	}
	return true
}

func (a *Attachment) SetAttachment(params parameter.Container, value types.AttachmentValue) error {
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
		a.Parameter = params
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
	a.Parameter = params
	a.Value = v
	return nil
}

// Categories is CATEGORIES
// https://tools.ietf.org/html/rfc5545#section-3.8.1.2
type Categories struct {
	Parameter parameter.Container
	Values    []types.Text
}

func (c *Categories) Decoce(w io.Writer) error {
	var s []string
	for _, v := range c.Values {
		s = append(s, string(v))
	}
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameCategories, c.Parameter.String(), strings.Join(s, ",")); err != nil {
		return err
	}
	return nil
}

func (c *Categories) Validate() error {
	// TODO
	return nil
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
	c.Parameter = params
	c.Values = values
	return nil
}

// Class is CLASS, optional property for components
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.3
type Class struct {
	Parameter parameter.Container
	Value     types.Text
}

func (c *Class) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameClass, c.Parameter.String(), c.Value); err != nil {
		return err
	}
	return nil
}

func (c *Class) Validate() error {
	// TODO
	return nil
}

func (c *Class) SetClass(params parameter.Container, value types.Text) error {
	switch ClassType(value) {
	case ClassTypePublic, ClassTypePrivate, ClassTypeConfidential:
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

func (c *Comment) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameComment, c.Parameter.String(), c.Value); err != nil {
		return err
	}
	return nil
}

func (c *Comment) Validate() error {
	// TODO
	return nil
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

// Description is DESCRIPTION
// https://tools.ietf.org/html/rfc5545#section-3.8.1.5
type Description struct {
	Parameter parameter.Container
	Value     types.Text
}

func (d *Description) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDescription, d.Parameter.String(), d.Value); err != nil {
		return err
	}
	return nil
}

func (d *Description) Validate() error {
	// TODO
	return nil
}

func (d *Description) SetDescription(params parameter.Container, value types.Text) error {
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

func (g *Geo) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%d;%d", NameGeo, g.Parameter.String(), g.Latitude, g.Longitude); err != nil {
		return err
	}
	return nil
}

func (g *Geo) Validate() error {
	// TODO
	return nil
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

// Location is LOCATION
// location of component
// https://tools.ietf.org/html/rfc5545#section-3.8.1.7
type Location struct {
	Parameter parameter.Container
	Value     types.Text
}

func (l *Location) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameLocation, l.Parameter.String(), l.Value); err != nil {
		return err
	}
	return nil
}

func (l *Location) Validate() error {
	// TODO
	return nil
}

func (l *Location) SetLocation(params parameter.Container, value types.Text) error {
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

func (pc *PercentComplete) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NamePercentComplete, pc.Parameter.String(), pc.Value); err != nil {
		return err
	}
	return nil
}

func (pc *PercentComplete) Validate() error {
	// TODO
	return nil
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

func (p *Priority) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NamePriority, p.Parameter.String(), p.Value); err != nil {
		return err
	}
	return nil
}

func (p *Priority) Validate() error {
	// TODO
	return nil
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

func (r *Resources) Decoce(w io.Writer) error {
	var s []string
	for _, v := range r.Values {
		s = append(s, string(v))
	}
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameResources, r.Parameter.String(), strings.Join(s, ",")); err != nil {
		return err
	}
	return nil
}

func (r *Resources) Validate() error {
	// TODO
	return nil
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

func (s *Status) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameStatus, s.Parameter.String(), s.Value); err != nil {
		return err
	}
	return nil
}

func (s *Status) Validate() error {
	// TODO
	return nil
}

func (s *Status) SetStatus(params parameter.Container, value types.Text, kind component.Type) error {
	v := StatusType(value)
	switch kind {
	case component.TypeEvent:
		switch v {
		case StatusTypeTentative, StatusTypeConfirmed, StatusTypeCancelled:
			s.Parameter = params
			s.Value = v
			return nil
		default:
			return fmt.Errorf("")
		}
	case component.TypeTODO:
		switch v {
		case StatusTypeNeedsAction, StatusTypeCompleted, StatusTypeInProcess, StatusTypeCancelled:
			s.Parameter = params
			s.Value = v
			return nil
		default:
			return fmt.Errorf("")
		}
	case component.TypeJournal:
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

func (s *Summary) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameSummary, s.Parameter.String(), s.Value); err != nil {
		return err
	}
	return nil
}

func (s *Summary) Validate() error {
	// TODO
	return nil
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
