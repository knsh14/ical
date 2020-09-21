package ical

import (
	"fmt"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

func NewCalender() *Calender {
	return &Calender{Version: property.NewVersion()}
}

// Calender is root object of ical
// https://tools.ietf.org/html/rfc5545#section-3.4
// https://tools.ietf.org/html/rfc5545#section-3.5
// https://tools.ietf.org/html/rfc5545#section-3.6
type Calender struct {
	// required field
	ProdID  *property.ProdID
	Version *property.Version

	// optional
	CalScale *property.CalScale
	Method   *property.Method

	XProperties    []*property.NonStandard // https://tools.ietf.org/html/rfc5545#section-3.8.8.2
	IANAProperties []*property.IANA

	Components []CalenderComponent
}

func (c *Calender) SetCalScale(params parameter.Container, value types.Text) error {
	if c.CalScale != nil {
		return c.CalScale.SetCalScale(params, value)
	}
	cs := &property.CalScale{}
	if err := cs.SetCalScale(params, value); err != nil {
		return err
	}
	c.CalScale = cs
	return nil
}

func (c *Calender) SetMethod(params parameter.Container, value types.Text) error {
	if c.Method != nil {
		return c.Method.SetMethod(params, value)
	}
	m := &property.Method{}
	if err := m.SetMethod(params, value); err != nil {
		return err
	}
	c.Method = m
	return nil
}

func (c *Calender) SetProdID(params parameter.Container, value types.Text) error {
	if c.ProdID != nil {
		return c.ProdID.SetProdID(params, value)
	}
	pid := &property.ProdID{}
	if err := pid.SetProdID(params, value); err != nil {
		return err
	}
	c.ProdID = pid
	return nil
}

func (c *Calender) SetVersion(params parameter.Container, value types.Text) error {
	if c.Version != nil {
		return c.Version.SetVersion(params, value)
	}
	ver := property.NewVersion()
	if err := ver.SetVersion(params, value); err != nil {
		return err
	}
	c.Version = ver
	return nil
}

func (c *Calender) Validate() error {
	if c.ProdID == nil {
		return NewValidationError(component.TypeCalendar, property.NameProdID, "must not to be nil")
	}
	if c.ProdID.Value == "" {
		return NewValidationError(component.TypeCalendar, property.NameProdID, "must not to be empty")
	}
	if c.Version == nil {
		return NewValidationError(component.TypeCalendar, property.NameVersion, "must not to be nil")
	}
	if c.Version.Max == "" {
		return NewValidationError(component.TypeCalendar, property.NameVersion, "max must not to be empty")
	}
	for _, component := range c.Components {
		if err := component.Validate(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}
