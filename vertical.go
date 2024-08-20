package tree2html

import (
	"sort"
)

type vTable struct{}

func (t *vTable) width(tree *Tree) int {
	if len(tree.Children) == 0 {
		return 1
	}
	w := 0
	for _, child := range tree.Children {
		w += t.width(child)
	}
	return w
}

func (t *vTable) lcm(a, b int) int {
	c := a * b
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return c / a
}

func (t *vTable) rowsToUse(tree *Tree) int {
	childrenRows := 0
	if len(tree.Children) > 0 {
		childrenRows = 1
	}
	for _, child := range tree.Children {
		childrenRows = t.lcm(childrenRows, t.rowsToUse(child))
	}
	return 1 + childrenRows
}

func (t *vTable) buildCells(tree *Tree, row, col, rowsLeft int) []*Cell {
	rootRows := rowsLeft / t.rowsToUse(tree)
	cells := []*Cell{NewVCell(tree, row, col, rootRows, t.width(tree))}

	for _, child := range tree.Children {
		cells = append(cells, t.buildCells(child, row+rootRows, col, rowsLeft-rootRows)...)
		col += t.width(child)
	}

	return cells
}

func (t *Tree) VTable() (tb Table) {
	vt := &vTable{}
	cells := vt.buildCells(t, 0, 0, vt.rowsToUse(t))

	sort.Slice(cells, func(i, j int) bool {
		if cells[i].Row != cells[j].Row {
			return cells[i].Row < cells[j].Row
		}
		return cells[i].Col < cells[j].Col
	})

	for i, ri := 0, 0; i < len(cells); ri++ {
		if ri == 0 {
			i++
			continue
		}

		var row []*Cell

		for ; i < len(cells) && cells[i].Row == ri; i++ {
			row = append(row, cells[i])
		}

		tb = append(tb, row)
	}
	return
}
