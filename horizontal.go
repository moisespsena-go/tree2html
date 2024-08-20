package tree2html

import (
	"io"
)

type HTable struct {
	Rows     [][]*Tree
	maxDepth int
}

func (t *HTable) WriteTo(w io.Writer) (n int64, err error) {
	return t.Write(&DefaultWriter{Writer: w})
}

func (t *HTable) Write(w Writer) (n int64, err error) {
	var n2 int64

	for _, row := range t.Rows {
		n2, err = w.OpenRow()
		n += n2
		if err != nil {
			return
		}

		for _, col := range row {
			n2, err = WriteCell(w, col, col.leafCount, 0)
			n += n2
			if err != nil {
				return
			}
		}

		for i := 0; i < t.maxDepth-row[len(row)-1].depth; i++ {
			n2, err = WriteCell(w, nil, 0, 0)
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

func NewHTable(tree *Tree) (t *HTable) {
	t = &HTable{
		maxDepth: MaxDepth(tree),
	}
	tree = tree.DeepCopy()

	var firsts []*Tree
	for tree != nil {
		firsts, tree = tree.PopFirsts()
		t.Rows = append(t.Rows, firsts)
	}
	return
}

func (t *Tree) PopFirsts() (firsts []*Tree, dot *Tree) {
	dot = t.popFirsts(&firsts)
	return
}

func (t *Tree) popFirsts(firsts *[]*Tree) (dot *Tree) {
	if t.parent != nil {
		*firsts = append(*firsts, t)
	}

	if len(t.Children) > 0 {
		dot = t.Children[0].popFirsts(firsts)
	} else {
		dot = Next(t)
		t.parent.Children = t.parent.Children[1:]
		for _, child := range t.parent.Children {
			child.index--
		}
	}
	return
}
