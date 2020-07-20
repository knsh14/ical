package parser

import (
	"fmt"

	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/lexer"
	"golang.org/x/sync/errgroup"
)

func convertContentLines(raw []string) ([]*contentline.ContentLine, error) {

	var eg errgroup.Group
	lines := make([]*contentline.ContentLine, len(raw))

	for i, s := range raw {
		i := i
		eg.Go(func() error {
			l := lexer.New(s)
			cl, err := contentline.ConvertContentLine(l)
			if err != nil {
				return fmt.Errorf("line:%d convert content line: %w", i, err)
			}
			lines[i] = cl
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return lines, nil
}
