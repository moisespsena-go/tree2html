package tree2html

import (
	"io"
	"sort"
)

type VCell struct {
	t       *Tree
	Row     int
	Col     int
	Rowspan int
	Colspan int
	Data    any
}

func NewVCell(t *Tree, row, col, rowspan, colspan int) *VCell {
	return &VCell{t: t, Row: row, Col: col, Rowspan: rowspan, Colspan: colspan}
}

type VTable struct {
	Cells []*VCell
}

func (t *VTable) width(tree *Tree) int {
	if len(tree.Children) == 0 {
		return 1
	}
	w := 0
	for _, child := range tree.Children {
		w += t.width(child)
	}
	return w
}

func (t *VTable) lcm(a, b int) int {
	c := a * b
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return c / a
}

func (t *VTable) rowsToUse(tree *Tree) int {
	childrenRows := 0
	if len(tree.Children) > 0 {
		childrenRows = 1
	}
	for _, child := range tree.Children {
		childrenRows = t.lcm(childrenRows, t.rowsToUse(child))
	}
	return 1 + childrenRows
}

func (t *VTable) buildCells(tree *Tree, row, col, rowsLeft int) []*VCell {
	rootRows := rowsLeft / t.rowsToUse(tree)
	cells := []*VCell{NewVCell(tree, row, col, rootRows, t.width(tree))}

	for _, child := range tree.Children {
		cells = append(cells, t.buildCells(child, row+rootRows, col, rowsLeft-rootRows)...)
		col += t.width(child)
	}

	return cells
}

func (t *VTable) WriteTo(w io.Writer) (n int64, err error) {
	return t.Write(NewDefaultWriter(w))
}

func (t *VTable) Write(w Writer) (n int64, err error) {
	var n2 int64

	for i, row := 0, 0; i < len(t.Cells); row++ {
		if row == 0 {
			i++
			continue
		}

		n2, err = w.OpenRow()
		n += n2
		if err != nil {
			return
		}

		for ; i < len(t.Cells) && t.Cells[i].Row == row; i++ {
			c := t.Cells[i]
			n2, err = WriteCell(w, c.t, c.Rowspan, c.Colspan)
			n += n2
			if err != nil {
				return
			}
		}

		n2, err = w.CloseRow()
		n += n2
		if err != nil {
			return
		}
	}
	return
}

func NewVTable(tree *Tree) (t *VTable) {
	t = &VTable{}
	cells := t.buildCells(tree, 0, 0, t.rowsToUse(tree))
	sort.Slice(cells, func(i, j int) bool {
		if cells[i].Row != cells[j].Row {
			return cells[i].Row < cells[j].Row
		}
		return cells[i].Col < cells[j].Col
	})
	t.Cells = cells
	return
}
