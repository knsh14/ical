package ical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/knsh14/ical/parameter"
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
	ProdID struct {
		Param parameter.Container
		Value types.Text
	}
	Version struct {
		Param parameter.Container
		Max   types.Text
		Min   types.Text
	}

	// optional
	CalScale struct {
		Valid bool
		Param parameter.Container
		Value types.Text
	}
	Method struct {
		Valid bool
		Param parameter.Container
		Value types.Text
	}

	XProperty map[string]struct {
		Name   string
		Param  parameter.Container
		Values []types.Text
	}
	IANAProperty map[string]struct {
		Name   string
		Param  parameter.Container
		Values []types.Text
	}

	Component []CalenderComponent
}

// SetCalScale updates CALSCALE property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.1
func (c *Calender) SetCalScale(params parameter.Container, t types.Text) error {
	if t == "" {
		return ErrInputIsEmpty
	}
	if string(t) != "GREGORIAN" {
		return fmt.Errorf("Invalid CALSCALE Value %s, allow only GREGORIAN", string(t))
	}
	c.CalScale.Param = params
	c.CalScale.Value = t
	return nil
}

// SetMethod updates Method property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.2
func (c *Calender) SetMethod(params parameter.Container, t types.Text) error {
	if t == "" {
		return ErrInputIsEmpty
	}
	if isMethod(string(t)) {
		c.Method.Param = params
		c.Method.Value = t
		return nil
	}
	return fmt.Errorf("Invalid Method Value %s, allow only registerd IANA tokens", t)
}

// SetProdID updates PRODID property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.3
func (c *Calender) SetProdID(params parameter.Container, t types.Text) error {
	if t == "" {
		return ErrInputIsEmpty
	}
	c.ProdID.Param = params
	c.ProdID.Value = t
	return nil
}

// SetVersion updates VERSION property
// spec is https://tools.ietf.org/html/rfc5545#section-3.7.4
func (c *Calender) SetVersion(params parameter.Container, t types.Text) error {
	if t == "" {
		return ErrInputIsEmpty
	}
	isMatch, err := regexp.MatchString(`^\d+.\d+$`, string(t))
	if err != nil {
		return err
	}
	if isMatch {
		c.Version.Param = params
		c.Version.Max = t
		return nil
	}
	isMatch, err = regexp.MatchString(`^\d+.\d+;\d+.\d+$`, string(t))
	if err != nil {
		return err
	}
	if !isMatch {
		return fmt.Errorf("not required format, allow X.Y or W.X;Y.Z")
	}
	v := strings.SplitN(string(t), ";", 2)
	c.Version.Min = types.NewText(v[0])
	c.Version.Max = types.NewText(v[1])
	c.Version.Param = params
	return nil
}

func (c *Calender) Validate() error {
	return nil
}
