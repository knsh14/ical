package ical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

// Miscellaneous Component Properties
// https://tools.ietf.org/html/rfc5545#section-3.8.8

// IANA is IANA-registered property
// https://tools.ietf.org/html/rfc5545#section-3.8.8.1
type IANA struct {
	Name      string
	Parameter parameter.Container
	Value     interface{} // any types can be
}

func NewIANA(name string, params parameter.Container, value interface{}) *IANA {
	return &IANA{
		Name:      name,
		Parameter: params,
		Value:     value,
	}
}

// NonStandard is property name with a "X-" prefix
// https://tools.ietf.org/html/rfc5545#section-3.8.8.2
type NonStandard struct {
	Name      string
	Parameter parameter.Container
	Value     interface{} // any types can be
}

func NewNonStandard(name string, params parameter.Container, value interface{}) (*NonStandard, error) {
	if !token.IsXName(name) {
		return nil, fmt.Errorf("%s", name)
	}
	return &NonStandard{
		Name:      name,
		Parameter: params,
		Value:     value,
	}, nil
}

// RequestStatus is REQUEST-STATUS
// https://tools.ietf.org/html/rfc5545#section-3.8.8.3
type RequestStatus struct {
	Parameter parameter.Container

	// StatusCode list is defined in https://tools.ietf.org/html/rfc5546#section-3.6
	StatusCode        types.Text
	StatusDescription types.Text
	ExtraData         types.Text
}

func (rs *RequestStatus) SetRequestStatus(params parameter.Container, value types.Text) error {
	values := strings.SplitN(string(value), ";", 3)
	if len(values) < 2 {
		return fmt.Errorf("")
	}
	var exData types.Text
	if len(values) == 3 {
		exData = types.Text(values[2])
	}
	return rs.Update(params, types.Text(values[0]), types.Text(values[1]), exData)
}

func (rs *RequestStatus) Update(params parameter.Container, code, desc, exdata types.Text) error {
	if desc == "" {
		return fmt.Errorf("description is empty")
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("too many language parameter, %d", len(params[parameter.TypeNameLanguage]))
	}

	// consider to check by defined status
	found, err := regexp.MatchString(`^\d\.\d{1,2}$`, string(code))
	if err != nil {
		return fmt.Errorf("find version strings: %w", err)
	}
	if !found {
		return fmt.Errorf("invalid pattern %s, must be X.YY", code)
	}

	rs.Parameter = params
	rs.StatusCode = code
	rs.StatusDescription = desc
	rs.ExtraData = exdata

	return nil
}
