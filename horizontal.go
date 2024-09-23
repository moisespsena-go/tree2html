package tree2html

func (t *Tree) HTable() (tb Table) {
	maxDepth := MaxDepth(t)

	var (
		firsts []*Tree
		ri     int
	)

	if t.IsLeaf() {
		return
	}

	dot := t.Children[0]
	for dot != nil {
		firsts, dot = FirstsOf(dot)
		var (
			endCell = firsts[len(firsts)-1]
			row     = make([]*Cell, len(firsts)+maxDepth-endCell.depth)
			i       int
		)

		for ; i < len(firsts); i++ {
			col := firsts[i]
			row[i] = &Cell{Node: col, Rowspan: col.leafCount, Row: ri, Col: i}
		}

		for ; i < len(row); i++ {
			row[i] = &Cell{Row: ri, Col: i}
		}

		tb = append(tb, row)
	}

	return
}

func FirstsOf(t *Tree) (firsts []*Tree, dot *Tree) {
	dot = t
	for !dot.IsLeaf() {
		firsts = append(firsts, dot)
		if dot = Next(dot); dot == nil {
			return
		}
	}
	firsts = append(firsts, dot)
	return firsts, Next(dot)
}
