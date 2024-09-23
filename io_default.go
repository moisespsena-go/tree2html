package tree2html

import (
	"fmt"
	"io"
	"strconv"
)

type DefaultWriter struct {
	io.Writer
	valueWriter     func(w io.Writer, node *Tree) (n int64, err error)
	cellTagHandlers []func(c *Cell, tag *Tag)
	rowTagHandlers  []func(c *Tag)
}

func NewDefaultWriter(writer io.Writer) *DefaultWriter {
	return &DefaultWriter{
		Writer:      writer,
		valueWriter: defaultValueWriter,
	}
}

func (w *DefaultWriter) SetValueWriter(valueWriter func(w io.Writer, node *Tree) (n int64, err error)) *DefaultWriter {
	w.valueWriter = valueWriter
	return w
}

func (w *DefaultWriter) CellTagHandler(handlers ...func(c *Cell, tag *Tag)) *DefaultWriter {
	w.cellTagHandlers = append(w.cellTagHandlers, handlers...)
	return w
}

func (w *DefaultWriter) RowTagHandler(handlers ...func(c *Tag)) *DefaultWriter {
	w.rowTagHandlers = append(w.rowTagHandlers, handlers...)
	return w
}

func (w *DefaultWriter) Row() (end func() (n int64, err error), i int64, err error) {
	tag := &Tag{TagName: "tr"}
	for _, h := range w.rowTagHandlers {
		h(tag)
	}
	if i, err = WriteResult(w.Write([]byte(tag.Open()))); err != nil {
		return
	}
	end = func() (n int64, err error) {
		return WriteResult(w.Write([]byte(tag.Close())))
	}
	return
}

func (w *DefaultWriter) Cell(c *Cell) (value func() (n int64, err error), end func() (n int64, err error), n int64, err error) {
	tag := &Tag{TagName: "td"}
	if c.Node != nil {
		if c.Colspan > 1 {
			tag.Attr(&CellTagAttr{
				Name:  "colspan",
				Value: strconv.Itoa(c.Colspan),
			})
		}

		if c.Rowspan > 1 {
			tag.Attr(&CellTagAttr{
				Name:  "rowspan",
				Value: strconv.Itoa(c.Rowspan),
			})
		}
		for _, h := range w.cellTagHandlers {
			h(c, tag)
		}
	}

	if n, err = WriteResult(w.Write([]byte(tag.Open()))); err != nil {
		return
	}

	value = func() (n int64, err error) {
		if wt, _ := c.Node.Value.(io.WriterTo); wt != nil {
			return wt.WriteTo(w.Writer)
		}

		return WriteResult(fmt.Fprint(w, c.Node.Value))
	}

	end = func() (n int64, err error) {
		return WriteResult(w.Write([]byte(tag.Close())))
	}
	return
}
