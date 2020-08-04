package ical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
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
	return c.UpdateVersion(params, types.NewText(v[0]), types.NewText(v[1]))
}

func (c *Calender) UpdateVersion(params parameter.Container, min, max types.Text) error {
	a, err := semver.NewVersion(string(min))
	if err != nil {
		return fmt.Errorf("convert %s to semvar: %w", min, err)
	}
	b, err := semver.NewVersion(string(min))
	if err != nil {
		return fmt.Errorf("convert %s to semvar: %w", max, err)
	}
	if a.GreaterThan(b) {
		return fmt.Errorf("min version %s is greater than max version %s", min, max)
	}
	c.Version.Min = min
	c.Version.Max = max
	c.Version.Param = params
	return nil
}

// SetXProp sets experimental property to calender.
// schema defined in https://tools.ietf.org/html/rfc5545#section-3.8.8.2
func (c *Calender) SetXProp(name string, params parameter.Container, values []types.Text) error {
	if c.XProperty == nil {
		c.XProperty = make(map[string]struct {
			Name   string
			Param  parameter.Container
			Values []types.Text
		})
	}
	if _, ok := c.XProperty[name]; ok {
		return fmt.Errorf("Property %s is already defined", name)
	}
	c.XProperty[name] = struct {
		Name   string
		Param  parameter.Container
		Values []types.Text
	}{
		Name:   name,
		Param:  params,
		Values: values,
	}
	return nil
}

func (c *Calender) Validate() error {
	return nil
}
