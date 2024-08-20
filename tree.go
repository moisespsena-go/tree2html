package tree2html

type Tree struct {
	parent    *Tree
	index     int
	Value     string
	Children  []*Tree
	leafCount int
	depth     int
}

func New(val string, children ...*Tree) *Tree {
	return &Tree{Value: val, Children: children}
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

func (t *Tree) Build() *Tree {
	t.Walk(func(p []*Tree, t *Tree, i int) {
		t.index = i
		t.parent = p[len(p)-1]
		t.depth = len(p)
		t.leafCount = LeafCount(t)
	})
	return t
}
