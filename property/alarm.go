package property

import (
	"fmt"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

// Alarm Component Properties
// https://tools.ietf.org/html/rfc5545#section-3.8.6

// Action is ACTION
// https://tools.ietf.org/html/rfc5545#section-3.8.6.1
type Action struct {
	Parameter parameter.Container
	Value     types.Text
}

func (a *Action) SetAction(params parameter.Container, value types.Text) error {
	switch ActionType(value) {
	case ActionTypeAudio, ActionTypeDisplay, ActionTypeEMail:
		a.Parameter = params
		a.Value = value
		return nil
	default:
		if token.IsXName(string(value)) {
			a.Parameter = params
			a.Value = value
			return nil
		}
		return fmt.Errorf("%s is invalid value for ACTION", value)
	}
}

// RepeatCount is REPEAT
// https://tools.ietf.org/html/rfc5545#section-3.8.6.2
type RepeatCount struct {
	Parameter parameter.Container
	Value     types.Integer
}

func (rc *RepeatCount) SetRepeatCount(params parameter.Container, value types.Integer) error {
	if value < 0 {
		return fmt.Errorf("value must be > 0")
	}
	rc.Parameter = params
	rc.Value = value
	return nil
}

func NewTriggerValue(params parameter.Container, value string) (types.TriggerValue, error) {
	if len(params[parameter.TypeNameValueType]) > 1 {
		return nil, fmt.Errorf("invalid %s parameter count", parameter.TypeNameValueType)
	}
	if len(params[parameter.TypeNameValueType]) == 0 {
		d, err := types.NewDuration(value)
		if err != nil {
			return nil, fmt.Errorf("value %s is not DURATION: %w", value, err)
		}
		return d, nil
	}

	valueType, ok := params[parameter.TypeNameValueType][0].(*parameter.ValueType)
	if !ok {
		return nil, fmt.Errorf("invalid type %T in %s ", params[parameter.TypeNameValueType][0], parameter.TypeNameValueType)
	}

	switch valueType.Value {
	case "DURATION":
		d, err := types.NewDuration(value)
		if err != nil {
			return nil, fmt.Errorf("value %s is not DURATION: %w", value, err)
		}
		return d, nil
	case "DATE-TIME":
		tz := params.GetTimezone()
		dt, err := types.NewDateTime(value, tz)
		if err != nil {
			return nil, fmt.Errorf("value %s is not DATE-TIME: %w", value, err)
		}
		return dt, nil
	default:
		return nil, fmt.Errorf("invalid value type %s, must be DURATION or DATE-TIME", valueType.Value)
	}
}

// Trigger is TRIGGER
// https://tools.ietf.org/html/rfc5545#section-3.8.6.3
type Trigger struct {
	Parameter parameter.Container
	Value     types.TriggerValue
}

func (t *Trigger) SetTrigger(params parameter.Container, value types.TriggerValue) error {
	if len(params[parameter.TypeNameValueType]) > 1 {
		return fmt.Errorf("invalid %s parameter count", parameter.TypeNameValueType)
	}
	if len(params[parameter.TypeNameValueType]) == 0 {
		if _, ok := value.(types.Duration); !ok {
			return fmt.Errorf("parameter value type is DURATION, but acutual type is %T", value)
		}
		t.Parameter = params
		t.Value = value
		return nil
	}

	valueType, ok := params[parameter.TypeNameValueType][0].(*parameter.ValueType)
	if !ok {
		return fmt.Errorf("invalid type %T in %s ", params[parameter.TypeNameValueType][0], parameter.TypeNameValueType)
	}

	switch valueType.Value {
	case "DURATION":
		if _, ok := value.(types.Duration); !ok {
			return fmt.Errorf("parameter value type is DURATION, but acutual type is %T", value)
		}
		t.Parameter = params
		t.Value = value
		return nil
	case "DATE-TIME":
		if _, ok := value.(types.DateTime); !ok {
			return fmt.Errorf("parameter value type is DATE-TIME, but acutual type is %T", value)
		}
		t.Parameter = params
		t.Value = value
		return nil
	default:
		return fmt.Errorf("invalid value type %s, must be DURATION or DATE-TIME", valueType.Value)
	}
}
