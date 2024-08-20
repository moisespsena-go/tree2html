package tree2html

func (t *Tree) HTable() (tb Table) {
	maxDepth := MaxDepth(t)

	var (
		firsts []*Tree
		ri     int
	)

	for t != nil {
		firsts, t = PopFirsts(t)
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

func PopFirsts(t *Tree) (firsts []*Tree, dot *Tree) {
	dot = popFirsts(t, &firsts)
	return
}

func popFirsts(t *Tree, firsts *[]*Tree) (dot *Tree) {
	if t.parent != nil {
		*firsts = append(*firsts, t)
	}

	if len(t.Children) > 0 {
		dot = popFirsts(t.Children[0], firsts)
	} else {
		dot = Next(t)
		t.parent.Children = t.parent.Children[1:]
		for _, child := range t.parent.Children {
			child.index--
		}
	}
	return
}
