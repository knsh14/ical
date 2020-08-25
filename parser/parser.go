package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/lexer"
	"golang.org/x/sync/errgroup"
)

func Parse(r io.Reader) (*ical.Calender, error) {
	return parseFromScanner(bufio.NewScanner(r))
}

func ParseFile(path string) (*ical.Calender, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	return parseFromScanner(scanner)
}

func parseFromScanner(scanner *bufio.Scanner) (*ical.Calender, error) {
	lines, err := scanLines(scanner)
	if err != nil {
		return nil, err
	}
	var eg errgroup.Group
	contentlines := make([]*contentline.ContentLine, len(lines))
	for i := range lines {
		i := i
		eg.Go(func() error {
			l := lexer.New(lines[i])
			cl, err := contentline.ConvertContentLine(l)
			if err != nil {
				return fmt.Errorf("convert content line in line %d: %w", i, err)
			}
			contentlines[i] = cl
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	p := NewParser(contentlines)
	return p.parse()
}

func scanLines(scanner *bufio.Scanner) ([]string, error) {
	var res []string
	for scanner.Scan() {
		l := scanner.Text()
		switch {
		case strings.HasPrefix(l, " "):
			res[len(res)-1] += "\n" + strings.TrimPrefix(l, " ")
		case strings.HasPrefix(l, "\t"):
			res[len(res)-1] += "\n" + strings.TrimPrefix(l, "\t")
			continue
		default:
			res = append(res, l)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func NewParser(cls []*contentline.ContentLine) *Parser {
	return &Parser{
		Lines: cls,
	}
}

type Parser struct {
	Lines                []*contentline.ContentLine
	CurrentIndex         int
	currentComponentType component.Type
	// errors               []error
}

func (p *Parser) getCurrentLine() *contentline.ContentLine {
	if p.CurrentIndex >= len(p.Lines) {
		return nil
	}
	return p.Lines[p.CurrentIndex]
}

func (p *Parser) nextLine() {
	p.CurrentIndex++
}

func (p *Parser) Parse() (*ical.Calender, error) {
	c, err := p.parse()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) parse() (*ical.Calender, error) {
	l := p.getCurrentLine()
	if !p.isBeginComponent(component.TypeCalendar) {
		return nil, fmt.Errorf("not %s:%s, got %v", "BEGIN", component.TypeCalendar, l)
	}
	c, err := p.parseCalender()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) isBeginComponent(c component.Type) bool {
	if p.getCurrentLine().Name != "BEGIN" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return component.Type(p.getCurrentLine().Values[0]) == c
}

func (p *Parser) isEndComponent(c component.Type) bool {
	if p.getCurrentLine().Name != "END" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return component.Type(p.getCurrentLine().Values[0]) == c
}
