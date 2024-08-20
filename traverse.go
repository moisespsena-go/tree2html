package tree2html

func Next(t *Tree) *Tree {
	return NextAfter(t, t.index)
}

func NextAfter(t *Tree, i int) *Tree {
	if len(t.Children) > i {
		return t.Children[i]
	}
	if t.parent == nil {
		return nil
	}
	return NextAfter(t.parent, t.index+1)
}
