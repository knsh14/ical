package parser

import (
	"fmt"
	"strings"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/contentline"
)

func Parse(path string) (*ical.Calender, error) {
	return nil, nil
}

func rebuildContentLines(raw []string) []string {
	var res []string
	for _, l := range raw {
		switch {
		case strings.HasPrefix(l, " "):
			res[len(res)-1] += strings.TrimPrefix(l, " ")
		case strings.HasPrefix(l, "\t"):
			res[len(res)-1] += strings.TrimPrefix(l, "\t")
			continue
		default:
			res = append(res, l)
		}
	}
	return res
}

type Parser struct {
	Lines        []*contentline.ContentLine
	CurrentIndex int
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
	if !p.isBeginComponent(ical.ComponentTypeCalender) {
		return nil, fmt.Errorf("not %s:%s, got %v", "BEGIN", ical.ComponentTypeCalender, l)
	}
	p.nextLine()
	c, err := p.parseCalender()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) isBeginComponent(component ical.ComponentType) bool {
	if p.getCurrentLine().Name != "BEGIN" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return ical.ComponentType(p.getCurrentLine().Values[0]) == component
}

func (p *Parser) isEndComponent(cl *contentline.ContentLine, component ical.ComponentType) bool {
	if p.getCurrentLine().Name != "END" {
		return false
	}
	if len(p.getCurrentLine().Values) != 1 {
		return false
	}
	return ical.ComponentType(p.getCurrentLine().Values[0]) == component
}
