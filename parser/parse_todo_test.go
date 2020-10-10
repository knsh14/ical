package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/types"
)

func TestParseToDo(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       []*contentline.ContentLine
		expected    *ical.ToDo
		assertError func(*testing.T, error)
	}{}
	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			p := NewParser(tc.input)
			actual, err := p.parseTodo()
			tc.assertError(t, err)
			if diff := cmp.Diff(tc.expected, actual, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
