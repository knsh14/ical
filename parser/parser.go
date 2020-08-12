package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/lexer"
	"golang.org/x/sync/errgroup"
)

func Parse(path string) (*ical.Calender, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
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

func rebuildContentLines(raw []string) []string {
	var res []string
	for _, l := range raw {
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
	return res
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
	currentComponentType component.ComponentType
	errors               []error
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
		// TODO: wrap
		return nil, err
	}
	return c, nil
}

func (p *Parser) parse() (*ical.Calender, error) {
	l := p.getCurrentLine()
	if !p.isBeginComponent(component.ComponentTypeCalendar) {
		return nil, fmt.Errorf("not %s:%s, got %v", "BEGIN", component.ComponentTypeCalendar, l)
	}
	p.nextLine()
	c, err := p.parseCalender()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) isBeginComponent(c component.ComponentType) bool {
	if p.getCurrentLine().Name != "BEGIN" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return component.ComponentType(p.getCurrentLine().Values[0]) == c
}

func (p *Parser) isEndComponent(c component.ComponentType) bool {
	if p.getCurrentLine().Name != "END" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return component.ComponentType(p.getCurrentLine().Values[0]) == c
}
