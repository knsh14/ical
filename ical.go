package ical

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/knsh14/ical/types"
)

func NewCalender() *Calender {
	return &Calender{}
}

// Calender is top object of ical
// https://tools.ietf.org/html/rfc5545#section-3.4
// https://tools.ietf.org/html/rfc5545#section-3.5
// https://tools.ietf.org/html/rfc5545#section-3.6
type Calender struct {
	// required field
	ProdID  *types.Text
	version *types.Text

	// optional
	CalScale *types.Text
	Method   *types.Text

	XProperty    []*types.Text // TODO fix to implement component
	IANAProperty []*types.Text // TODO fix to implement component

	Component []CalenderComponent
}

// SetCalScale updates CALSCALE property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.1
func (c *Calender) SetCalScale(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	if t.Value != "GREGORIAN" {
		return fmt.Errorf("Invalid CALSCALE Value %s, allow only GREGORIAN", t.Value)
	}
	c.CalScale = t
	return nil
}

// SetMethod updates Method property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.2
func (c *Calender) SetMethod(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	if isMethod(t.Value) {
		c.Method = t
		return nil
	}
	return fmt.Errorf("Invalid Method Value %s, allow only registerd IANA tokens", t.Value)
}

// SetProdID updates PRODID property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.3
func (c *Calender) SetProdID(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	c.ProdID = t
	return nil
}

// SetVersion updates VERSION property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.4
func (c *Calender) SetVersion(t *types.Text) error {
	if t == nil {
		return fmt.Errorf("input is nil")
	}
	isMatch, err := regexp.MatchString(`^\d+.\d+$`, t.Value)
	if err != nil {
		return err
	}
	if isMatch {
		c.version = t
		return nil
	}
	isMatch, err = regexp.MatchString(`^\d+.\d+:\d+.\d+$`, t.Value)
	if err != nil {
		return err
	}
	if !isMatch {
		return fmt.Errorf("not required format, allow X.Y or W.X:Y.Z")
	}
	c.version = t
	return nil
}

// GetVersion returns min and max version of iCal.
// if property have only one version, then it will be max version.
// this specification is defined in https://tools.ietf.org/html/rfc5545#section-3.7.4
func (c *Calender) GetVersion() (string, string) {
	v := strings.SplitN(c.version.Value, ":", 2)
	if len(v) == 2 {
		return v[0], v[1]
	}
	return "", v[0]
}

type CalenderComponent interface {
	implementCalender()
	Decode(io.Writer) error
}
