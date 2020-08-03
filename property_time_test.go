package ical

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical/parameter"
)

func TestDateTimeEnd(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		input struct {
			param parameter.Container
			value interface{}
		}
		expect      *DateTimeEnd
		expectError error
	}{}

	for i, tt := range testcases {
		tt := tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			v := &DateTimeEnd{}
			err := v.SetEnd(tt.input.param, tt.input.value)
			if err != tt.expectError {
				t.Fatalf("unexpected %s, %s", err, tt.expectError)
			}
			if diff := cmp.Diff(v, tt.expect); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
