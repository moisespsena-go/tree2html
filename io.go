package tree2html

import (
	"fmt"
	"io"
)

type CellOpener func(w io.Writer, c *Cell) (doValue func(), end func())

type Writer interface {
	io.Writer

	Row() (end func() (n int64, err error), i int64, err error)
	Cell(c *Cell) (value func() (n int64, err error), end func() (n int64, err error), n int64, err error)
}

func WriteCell(w Writer, c *Cell) (n int64, err error) {
	var (
		n2 int64
		val,
		end func() (n int64, err error)
	)
	val, end, n, err = w.Cell(c)
	n += n2
	if err != nil {
		return
	}

	if c.Node != nil {
		n2, err = val()
		n += n2
		if err != nil {
			return
		}
	}

	n2, err = end()
	n += n2
	return
}

func defaultValueWriter(w io.Writer, t *Tree) (n int64, err error) {
	i, err := fmt.Fprint(w, t.Value)
	return int64(i), err
}
