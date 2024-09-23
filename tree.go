package tree2html

type Tree struct {
	parent    *Tree
	index     int
	leafCount int
	depth     int
	Value     any     `json:"Value,omitempty"`
	Children  []*Tree `json:"Children,omitempty"`
}

func Node(val any, children ...*Tree) *Tree {
	return &Tree{Value: val, Children: children}
}

func New(children ...*Tree) *Tree {
	return Node(nil, children...).Build()
}

func (t *Tree) Parent() *Tree {
	return t.parent
}

func (t *Tree) Root() *Tree {
	for t.parent != nil {
		t = t.parent
	}
	return t
}

func (t *Tree) Index() int {
	return t.index
}

func (t *Tree) Depth() int {
	return t.depth
}

func (t *Tree) IsLeaf() bool {
	return len(t.Children) == 0
}

func (t Tree) DeepCopy() *Tree {
	children := make([]*Tree, len(t.Children))
	for i, child := range t.Children {
		child = child.DeepCopy()
		child.parent = &t
		children[i] = child
	}
	t.Children = children
	t.parent = nil
	return &t
}

func (t *Tree) Walk(f func(p []*Tree, t *Tree, i int)) {
	t.walk(nil, f)
}

func (t *Tree) walk(p []*Tree, f func(path []*Tree, t *Tree, i int)) {
	p = append(p, t)
	for i, child := range t.Children {
		f(p, child, i)
		child.walk(p, f)
	}
}

func (t *Tree) Append(child ...*Tree) *Tree {
	t.Children = append(t.Children, child...)
	return t
}

func (t *Tree) Build() *Tree {
	t.build(0)
	return t
}

func (t *Tree) build(depth int) {
	t.depth = depth

	for i, child := range t.Children {
		child.index = i
		child.parent = t
		child.build(depth + 1)
		if child.IsLeaf() {
			t.leafCount++
		} else {
			t.leafCount += child.leafCount
		}
	}
}
