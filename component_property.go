package ical

import (
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical/types"
)

type Property struct {
	// optional field
	Attachment *types.Text
	Categories *types.TextList
	Class      *types.Text
	Comment    *types.Text

	Description *types.Text
	Geo         [2]*types.Float
	Location    *types.Text
}

// SetCategories updates CATEGORIES
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.2
func (p *Property) SetCategories(tl *types.TextList) error {
	if tl == nil {
		return fmt.Errorf("input is nil")
	}
	// TODO: check params has language pattern parameter more than one
	p.Categories = tl
	return nil
}

// SetClass updates CLASS
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.3
func (p *Property) SetClass(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	switch ClassType(t.Value) {
	case PUBLIC, PRIVATE, CONFIDENTIAL:
		p.Class = t
		return nil
	}
	// iana-token
	if m, err := regexp.MatchString(`^[0-9,a-z,A-Z,-]*$`, t.Value); err == nil && m {
		p.Class = t
		return nil
	}
	// x-token
	if m, err := regexp.MatchString(`^X-([0-9a-zA-Z]{3}-)?[0-9a-zA-Z-]+$`, t.Value); err == nil && m {
		p.Class = t
		return nil
	}
	return fmt.Errorf("invalid value: %s", t.Value)
}

// SetComment updates COMMENT
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.4
func (p *Property) SetComment(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	// TODO: check param
	p.Comment = t
	return nil
}

// SetDescription updates DESCRIPTION
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.5
func (p *Property) SetDescription(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	// TODO: check param
	p.Description = t
	return nil
}

// SetGeo updates GEO
// specification https://tools.ietf.org/html/rfc5545#section-3.8.1.6
func (p *Property) SetGeo(latitude, longitude *types.Float) error {
	if latitude == nil {
		return fmt.Errorf("latitude is nil")
	}
	if longitude == nil {
		return fmt.Errorf("longitude is nil")
	}
	if diff := cmp.Diff(latitude.Parameters, longitude.Parameters); diff != "" {
		return fmt.Errorf("parameters are not equal %s", diff)
	}
	// TODO: check param
	p.Geo = [2]*types.Float{latitude, longitude}
	return nil
}
