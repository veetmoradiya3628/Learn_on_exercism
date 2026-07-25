package pov

type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	if tr == nil {
		return ""
	}
	return tr.value
}

// Children returns a slice containing the children of a tree.
func (tr *Tree) Children() []*Tree {
	if tr == nil {
		return nil
	}
	return tr.children
}

func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// findPath returns a slice of *Tree representing the nodes from root to target.
func findPath(node *Tree, target string) []*Tree {
	if node == nil {
		return nil
	}
	if node.value == target {
		return []*Tree{node}
	}
	for _, child := range node.children {
		if path := findPath(child, target); path != nil {
			return append([]*Tree{node}, path...)
		}
	}
	return nil
}

// removeChild removes target child from node's children slice.
func (tr *Tree) removeChild(target *Tree) {
	for i, child := range tr.children {
		if child == target {
			tr.children = append(tr.children[:i], tr.children[i+1:]...)
			return
		}
	}
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	path := findPath(tr, from)
	if path == nil {
		return nil
	}

	// Reverse parent-child relationships along the path
	for i := 0; i < len(path)-1; i++ {
		parent := path[i]
		child := path[i+1]

		parent.removeChild(child)
		child.children = append(child.children, parent)
	}
	return path[len(path)-1]
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	rerooted := tr.FromPov(from)
	if rerooted == nil {
		return nil
	}

	pathNodes := findPath(rerooted, to)
	if pathNodes == nil {
		return nil
	}

	result := make([]string, len(pathNodes))
	for i, node := range pathNodes {
		result[i] = node.value
	}
	return result
}