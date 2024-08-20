package tree2html

import "io"

type Cell struct {
	Node    *Tree
	Row     int
	Col     int
	Rowspan int
	Colspan int
}

func NewVCell(n *Tree, row, col, rowspan, colspan int) *Cell {
	return &Cell{
		Node:    n,
		Row:     row,
		Col:     col,
		Rowspan: rowspan,
		Colspan: colspan,
	}
}

type Row []*Cell

func (r Row) WriteTo(w io.Writer) (n int64, err error) {
	return r.Write(NewDefaultWriter(w))
}

func (r Row) Write(w Writer) (n int64, err error) {
	var n2 int64

	n2, err = w.OpenRow()
	n += n2
	if err != nil {
		return
	}

	for _, cell := range r {
		n2, err = WriteCell(w, cell)
		n += n2
		if err != nil {
			return
		}
	}

	n2, err = w.CloseRow()
	n += n2
	return
}

type Table []Row

func (t Table) WriteTo(w io.Writer) (n int64, err error) {
	return t.Write(NewDefaultWriter(w))
}

func (t Table) Write(w Writer) (n int64, err error) {
	var n2 int64

	for _, row := range t {
		n2, err = row.Write(w)
		n += n2
		if err != nil {
			return
		}
	}
	return
}
