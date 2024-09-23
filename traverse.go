package tree2html

func Next(t *Tree) *Tree {
	if !t.IsLeaf() {
		return t.Children[0]
	}
	for t.parent != nil {
		if len(t.parent.Children) > t.index+1 {
			return t.parent.Children[t.index+1]
		}
		t = t.parent
	}
	return nil
}
