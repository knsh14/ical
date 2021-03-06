package property

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// properties for only Calender component
// https://tools.ietf.org/html/rfc5545#section-3.7

// CalScale is CALSCALE
// https://tools.ietf.org/html/rfc5545#section-3.7.1
type CalScale struct {
	Parameter parameter.Container
	Value     types.Text
}

func (cs *CalScale) SetCalScale(params parameter.Container, value types.Text) error {
	if value == "" {
		return ErrInputIsEmpty
	}
	if value != types.Text("GREGORIAN") {
		return fmt.Errorf("Invalid CALSCALE Value %s, allow only GREGORIAN", value)
	}
	cs.Parameter = params
	cs.Value = value
	return nil
}

func (cs *CalScale) Decode(w io.Writer) error {
	if err := cs.Validate(); err != nil {
		return err
	}
	fmt.Fprintf(w, "%s%s:%s", NameCalScale, cs.Parameter.String(), cs.Value)
	return nil
}

func (cs *CalScale) Validate() error {
	if cs.Value != types.Text("GREGORIAN") {
		return fmt.Errorf("allow only \"GREGORIAN\", but %s", cs.Value)
	}
	return nil
}

// Method is Method
// https://tools.ietf.org/html/rfc5545#section-3.7.2
type Method struct {
	Parameter parameter.Container
	Value     types.Text
}

func (m *Method) SetMethod(params parameter.Container, value types.Text) error {
	if value == "" {
		return ErrInputIsEmpty
	}
	if isMethod(string(value)) {
		m.Parameter = params
		m.Value = value
		return nil
	}
	return fmt.Errorf("Invalid Method Value %s, allow only registerd IANA tokens", value)
}

func (m *Method) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameMethod, m.Parameter.String(), m.Value)
	return nil
}
func (m *Method) Validate() error {
	if m.Value == "" {
		return ErrInputIsEmpty
	}
	if !isMethod(string(m.Value)) {
		return fmt.Errorf("Invalid Method Value %s, allow only registerd IANA tokens", m.Value)
	}
	return nil
}

// ProdID is PRODID
// https://tools.ietf.org/html/rfc5545#section-3.7.3
type ProdID struct {
	Parameter parameter.Container
	Value     types.Text
}

func (p *ProdID) SetProdID(params parameter.Container, value types.Text) error {
	if value == "" {
		return ErrInputIsEmpty
	}
	p.Parameter = params
	p.Value = value
	return nil
}

func (p *ProdID) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameProdID, p.Parameter.String(), p.Value)
	return nil
}

func (p *ProdID) Validate() error {
	if p.Value == "" {
		return ErrInputIsEmpty
	}
	return nil
}

func NewVersion() *Version {
	return &Version{
		Max: types.Text("2.0"),
	}
}

// Version is VERSION
// https://tools.ietf.org/html/rfc5545#section-3.7.4
type Version struct {
	Parameter parameter.Container
	Min, Max  types.Text
}

func (v *Version) Decode(w io.Writer) error {
	s := v.Max
	if v.Min == "" {
		s = v.Min + ";" + s
	}
	fmt.Fprintf(w, "%s%s:%s", NameVersion, v.Parameter.String(), s)
	return nil
}

func (v *Version) Validate() error {
	// TODO
	return nil
}

func (v *Version) SetVersion(params parameter.Container, value types.Text) error {
	if value == "" {
		return ErrInputIsEmpty
	}
	isMatch, err := regexp.MatchString(`^\d+.\d+$`, string(value))
	if err != nil {
		return err
	}
	if isMatch {
		v.Parameter = params
		v.Max = value
		return nil
	}
	isMatch, err = regexp.MatchString(`^\d+.\d+;\d+.\d+$`, string(value))
	if err != nil {
		return err
	}
	if !isMatch {
		return fmt.Errorf("not required format, allow X.Y or W.X;Y.Z")
	}
	versions := strings.SplitN(string(value), ";", 2)
	if len(versions) != 2 {
		return fmt.Errorf("versions must be 2, but %d", len(versions))
	}
	return v.UpdateVersion(params, types.NewText(versions[0]), types.NewText(versions[1]))
}

func (v *Version) UpdateVersion(params parameter.Container, min, max types.Text) error {
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
	v.Parameter = params
	v.Min = min
	v.Max = max
	return nil
}
