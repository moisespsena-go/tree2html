package tree2html

func LeafCount(t *Tree) int {
	if len(t.Children) == 0 {
		return 1
	}
	w := 0
	for _, child := range t.Children {
		w += LeafCount(child)
	}
	return w
}

func MaxDepth(t *Tree) (i int) {
	maxDepth(t, 0, &i)
	return
}

func maxDepth(t *Tree, d int, dst *int) {
	if len(t.Children) > 0 {
		d++
		for _, child := range t.Children {
			maxDepth(child, d, dst)
		}
	} else if d > *dst {
		*dst = d
	}
}

func WriteResult(n int, err error) (int64, error) {
	return int64(n), err
}
