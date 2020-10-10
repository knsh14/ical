package property

import (
	"fmt"
	"io"
	"time"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// Change Management Component Properties
// https://tools.ietf.org/html/rfc5545#section-3.8.7

// DateTimeCreated is CREATED
// https://tools.ietf.org/html/rfc5545#section-3.8.7.1
type DateTimeCreated struct {
	Parameter parameter.Container
	Value     types.DateTime
}

func (dc *DateTimeCreated) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameDateTimeCreated, dc.Parameter.String(), dc.Value)
	return nil
}

func (dc *DateTimeCreated) Validate() error {
	// TODO
	return nil
}

func (dc *DateTimeCreated) SetDateTimeCreated(params parameter.Container, value types.DateTime) error {
	if value == types.DateTime(time.Time{}) {
		return fmt.Errorf("input is nil")
	}
	dc.Parameter = params
	dc.Value = value
	return nil
}

// DateTimeStamp is DTSTAMP
// https://tools.ietf.org/html/rfc5545#section-3.8.7.2
type DateTimeStamp struct {
	Parameter parameter.Container
	Value     types.DateTime
}

func (ds *DateTimeStamp) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameDateTimeStamp, ds.Parameter.String(), ds.Value)
	return nil
}

func (ds *DateTimeStamp) Validate() error {
	// TODO
	return nil
}

func (ds *DateTimeStamp) SetDateTimeStamp(params parameter.Container, value types.DateTime) error {
	if value == types.DateTime(time.Time{}) {
		return fmt.Errorf("input is nil")
	}
	ds.Parameter = params
	ds.Value = value
	return nil
}

// LastModified is LAST-MODIFIED
// https://tools.ietf.org/html/rfc5545#section-3.8.7.3
type LastModified struct {
	Parameter parameter.Container
	Value     types.DateTime
}

func (lm *LastModified) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameLastModified, lm.Parameter.String(), lm.Value)
	return nil
}

func (lm *LastModified) Validate() error {
	panic("implement me")
}

func (lm *LastModified) SetLastModified(params parameter.Container, value types.DateTime) error {
	if value == types.DateTime(time.Time{}) {
		return fmt.Errorf("input is nil")
	}
	lm.Parameter = params
	lm.Value = value
	return nil
}

// SequenceNumber is SEQUENCE
// https://tools.ietf.org/html/rfc5545#section-3.8.7.4
type SequenceNumber struct {
	Parameter parameter.Container
	Value     types.Integer // default is 0
}

func (sn *SequenceNumber) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s%s:%s", NameLastModified, lm.Parameter.String(), lm.Value)
	return nil
}

func (sn *SequenceNumber) Validate() error {
}

func (sn *SequenceNumber) SetSequenceNumber(params parameter.Container, value types.Integer) error {
	if value < 0 {
		return fmt.Errorf("value must be > 0")
	}
	sn.Parameter = params
	sn.Value = value
	return nil
}
