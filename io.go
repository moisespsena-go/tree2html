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
		cs = fmt.Sprintf(" colspan='%d'", colspan)
	}
	rs := ""
	if rowspan > 1 {
		rs = fmt.Sprintf(" rowspan='%d'", rowspan)
	}
	depth := fmt.Sprintf(" data-depth='%d'", t.depth)
	i, err := fmt.Fprintf(w, "<td%s%s%s>", cs, rs, depth)
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

func WriteCell(w Writer, col *Tree, rowspan, colspan int) (n int64, err error) {
	var n2 int64
	n2, err = w.OpenCell(col, rowspan, colspan)
	n += n2
	if err != nil {
		return
	}

	if col != nil {
		n2, err = w.CellValue(col.Value)
		n += n2
		if err != nil {
			return
		}
	}

	n2, err = w.CloseCell()
	n += n2
	return
}
