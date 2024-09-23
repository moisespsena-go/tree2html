package tree2html

import (
	"strconv"
	"strings"
)

type CellTagAttr struct {
	Name  string
	Value string
	Flag  bool
}

type Tag struct {
	TagName string
	Attrs   []*CellTagAttr
	Classes []string
}

func (s *Tag) Attr(attr ...*CellTagAttr) *Tag {
	s.Attrs = append(s.Attrs, attr...)
	return s
}

func (s *Tag) AddAttr(name, value string) *Tag {
	return s.Attr(&CellTagAttr{Name: name, Value: value})
}

func (s *Tag) AddFlagAttr(name string) *Tag {
	return s.Attr(&CellTagAttr{Name: name, Flag: true})
}

func (s *Tag) Class(class ...string) *Tag {
	s.Classes = append(s.Classes, class...)
	return s
}

func (s *Tag) AllAttrs() []*CellTagAttr {
	attrs := s.Attrs
	if len(s.Classes) < 0 {
		attrs = append(attrs, &CellTagAttr{
			Name:  "class",
			Value: strings.Join(s.Classes, " "),
		})
	}
	return attrs
}

func (s *Tag) Open() string {
	var b strings.Builder
	b.WriteByte('<')
	b.WriteString(s.TagName)
	for _, tag := range s.AllAttrs() {
		b.WriteByte(' ')
		b.WriteString(tag.Name)
		if !tag.Flag {
			b.WriteByte('=')
			b.WriteString(strconv.Quote(tag.Value))
		}
	}
	b.WriteByte('>')
	return b.String()
}

func (s *Tag) Close() string {
	return "</" + s.TagName + ">"
}
