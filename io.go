package tree2html

import (
	"fmt"
	"io"
)

type Writer interface {
	io.Writer

	OpenRow() (i int64, err error)
	CloseRow() (i int64, err error)
	OpenCell(t *Tree, rowspan, colspan int) (i int64, err error)
	CloseCell() (i int64, err error)
	CellValue(data any) (i int64, err error)
}

type DefaultWriter struct {
	io.Writer
	valueWriter func(w io.Writer, node *Tree) (n int64, err error)
}

func NewDefaultWriter(writer io.Writer) *DefaultWriter {
	return &DefaultWriter{Writer: writer, valueWriter: defaultValueWriter}
}

func (w *DefaultWriter) SetValueWriter(valueWriter func(w io.Writer, node *Tree) (n int64, err error)) *DefaultWriter {
	w.valueWriter = valueWriter
	return w
}

func (w *DefaultWriter) OpenRow() (int64, error) {
	i, err := w.Write([]byte(`<tr>`))
	return int64(i), err
}

func (w *DefaultWriter) CloseRow() (int64, error) {
	i, err := w.Write([]byte(`</tr>`))
	return int64(i), err
}

func (w *DefaultWriter) OpenCell(t *Tree, rowspan, colspan int) (int64, error) {
	if t == nil {
		i, err := w.Write([]byte(`<td>`))
		return int64(i), err
	}
	cs := ""
	if colspan > 1 {
		cs = fmt.Sprintf(` colspan="%d"`, colspan)
	}
	rs := ""
	if rowspan > 1 {
		rs = fmt.Sprintf(` rowspan="%d"`, rowspan)
	}
	i, err := fmt.Fprintf(w, "<td%s%s>", cs, rs)
	return int64(i), err
}

func (w *DefaultWriter) CloseCell() (int64, error) {
	i, err := w.Write([]byte(`</td>`))
	return int64(i), err
}

func (w *DefaultWriter) CellValue(data any) (int64, error) {
	i, err := fmt.Fprint(w, data)
	return int64(i), err
}

func WriteCell(w Writer, c *Cell) (n int64, err error) {
	var n2 int64
	n2, err = w.OpenCell(c.Node, c.Rowspan, c.Colspan)
	n += n2
	if err != nil {
		return
	}

	if c.Node != nil {
		n2, err = w.CellValue(c.Node.Value)
		n += n2
		if err != nil {
			return
		}
	}

	n2, err = w.CloseCell()
	n += n2
	return
}

func defaultValueWriter(w io.Writer, t *Tree) (n int64, err error) {
	i, err := fmt.Fprint(w, t.Value)
	return int64(i), err
}
