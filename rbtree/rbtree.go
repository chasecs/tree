package rbtree

type Keytype interface {
	LessThan(interface{}) bool
}

type RBTree struct {
	Root *Node
	Size int
}

type Node struct {
	Parent, Left, Right *Node
	IsBlack             bool
	Key                 Keytype
	Value               interface{}
}

func New() *RBTree {
	return &RBTree{}
}

// Find finds the node by key and return it, if not exists return nil.
func (t *RBTree) Find(key Keytype) *Node {
	x := t.Root
	for x != nil {
		if key.LessThan(x.Key) {
			x = x.Left
		} else {
			if key == x.Key {
				return x
			}
			x = x.Right
		}
	}
	return nil
}

//Insert insert item into the tree
func (t *RBTree) Insert(key Keytype, value interface{}) error {
	x := t.Root
	var p *Node

	for x != nil {
		p = x
		if key.LessThan(x.Key) {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	n := &Node{Parent: p, IsBlack: false, Key: key, Value: value}
	t.Size++

	if p == nil {
		t.Root = n
	} else if n.Key.LessThan(p.Key) {
		p.Left = n
	} else {
		p.Right = n
	}

	t.insertFixup(n)

	return nil
}

func (t *RBTree) insertFixup(n *Node) {
	if n.Parent == nil {
		//case 1 n is root
		n.IsBlack = true
		return
	} else if isRed(n.Parent) {
		if n.Parent == n.Parent.Parent.Left {
			if isRed(n.Parent.Parent.Right) {
				// case 2: p red, u red
				n.Parent.Parent.Left.IsBlack = true
				n.Parent.Parent.Right.IsBlack = true
				n.Parent.Parent.IsBlack = false
				t.insertFixup(n.Parent.Parent)
			} else {
				if n == n.Parent.Right {
					// case 4-1: p red, u IsBlack, n == p.right
					n = n.Parent
					t.rotateLeft(n)
				}
				//case 4-2
				t.rotateRight(n.Parent.Parent)
				n.Parent.IsBlack = true
				n.Parent.Right.IsBlack = false
				return
			}
		} else {
			if isRed(n.Parent.Parent.Left) {
				// case 2: p red, u red
				n.Parent.Parent.Left.IsBlack = true
				n.Parent.Parent.Right.IsBlack = true
				n.Parent.Parent.IsBlack = false
				t.insertFixup(n.Parent.Parent)
			} else {
				if n == n.Parent.Left {
					// case 4-3: p red, u IsBlack, n == p.left
					n = n.Parent
					t.rotateRight(n)
				}
				//case 4-4
				t.rotateLeft(n.Parent.Parent)
				n.Parent.IsBlack = true
				n.Parent.Left.IsBlack = false
				return
			}
		}
	}
	//case 3: p red, do nothing
	return
}

// Delete move a specified item from the tree
func (t *RBTree) Delete(key Keytype) *Node {

	deletedNode := t.Find(key)
	if deletedNode == nil {
		return nil
	}

	var s *Node
	delItem := &Node{}
	*delItem = *deletedNode
	if deletedNode.Left != nil && deletedNode.Right != nil {
		s = successor(deletedNode)
		// replace deletedNode with successor
		deletedNode.Key = s.Key
		deletedNode.Value = s.Value
	} else {
		s = deletedNode
	}

	// Now z must be a leaf or a single child parent
	// case 1-1, z is red(it must be a no child leaf)
	if isRed(s) {
		// case1-1: z is a red leaf
		if s == s.Parent.Left {
			s.Parent.Left = nil
		} else {
			s.Parent.Right = nil
		}

		return delItem
	}

	//z is black

	var sChild *Node
	if s.Right != nil {
		sChild = s.Right
	} else {
		// succChil maybe nil
		sChild = s.Left
	}

	sParent := s.Parent
	if sParent == nil {
		t.Root = sChild
	} else if s == sParent.Left {
		sParent.Left = sChild
	} else {
		sParent.Right = sChild
	}
	if sChild != nil {
		sChild.Parent = sParent
	}

	t.deleteFixup(sChild, sParent)
	return delItem

}
func (t *RBTree) deleteFixup(n *Node, nParent *Node) {

	for !isRed(n) && n != t.Root {
		if n != nil {
			nParent = n.Parent
		}
		if n == nParent.Left {
			sibling := nParent.Right
			if isRed(sibling) {
				// has a red sibling
				t.rotateLeft(nParent)
				nParent.IsBlack = false
				nParent.Parent.IsBlack = true
				sibling = nParent.Right
			} else if !isRed(sibling.Right) && !isRed(sibling.Left) {
				// sibling and its children both are black
				sibling.IsBlack = false
				if isRed(nParent) {
					nParent.IsBlack = true
					return
				} else {
					n = nParent
				}
			} else {
				// one of sibling's chilren must be red
				if !isRed(sibling.Right) {
					t.rotateRight(sibling)
					sibling = nParent.Right
					sibling.IsBlack = true
					sibling.Right.IsBlack = false
				}

				t.rotateLeft(nParent)
				nParent.Parent.IsBlack = nParent.IsBlack
				nParent.IsBlack = true
				nParent.Parent.Right.IsBlack = true
				return
			}
		} else {
			sibling := nParent.Left
			if isRed(sibling) {
				t.rotateRight(nParent)
				nParent.IsBlack = false
				nParent.Parent.IsBlack = true
				sibling = nParent.Left
			} else if !isRed(sibling.Right) && !isRed(sibling.Left) {

				sibling.IsBlack = false
				// if isRed(nParent) {
				// 	nParent.IsBlack = true
				// 	return
				// } else {
				n = nParent
				// }
			} else {

				if !isRed(sibling.Left) {
					if sibling.Right != nil {
						sibling.Right.IsBlack = true
					}
					sibling.IsBlack = false
					t.rotateLeft(sibling)
					sibling = nParent.Left
				}

				sibling.IsBlack = nParent.IsBlack
				nParent.IsBlack = true
				if sibling.Left != nil {
					sibling.Left.IsBlack = true
				}
				t.rotateRight(nParent)
				return
			}
		}

	}

	if n != nil {
		n.IsBlack = true
	}
}

func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return !n.IsBlack
}

func (t *RBTree) rotateRight(n *Node) *Node {
	// necessary ?
	// YES
	if n.Left == nil {
		return nil
	}

	y := n.Left

	y.Parent = n.Parent
	if n.Parent == nil {
		t.Root = y
	} else if n == n.Parent.Right {
		n.Parent.Right = y
	} else {
		n.Parent.Left = y
	}

	n.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = n
	}
	y.Right = n
	n.Parent = y

	return nil
}

func (t *RBTree) rotateLeft(n *Node) *Node {
	// When we do a left rotation on a node x, we assume that its right child y is not T:nil;
	// x may be any node in the tree whose right child is not T:nil.
	if n.Right == nil {
		return nil
	}
	y := n.Right
	y.Parent = n.Parent
	if n.Parent == nil {
		t.Root = y
	} else if n == n.Parent.Right {
		n.Parent.Right = y
	} else {
		n.Parent.Left = y
	}
	n.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = n
	}
	n.Parent = y
	y.Left = n
	return nil
}

// IsBalance check if all  paths from the root to descendant leaves
// contain the same number of black nodes
func (t *RBTree) IsBalance() bool {

	blackNum := 0
	x := t.Root
	if x == nil {
		return true
	}
	//count one of the paths
	for x != nil {
		if x.IsBlack {
			blackNum++
		}
		x = x.Left
	}

	return isBalance(t.Root, blackNum)
}

func isBalance(node *Node, blackNum int) bool {
	if node == nil {
		return blackNum == 0
	} else if isRed(node) && isRed(node.Parent) {
		return false
	} else if node.IsBlack {
		blackNum--
	}
	return isBalance(node.Left, blackNum) && isBalance(node.Right, blackNum)
}

// min of right subtree
func successor(x *Node) *Node {
	rightMin := x.Right
	if rightMin == nil {
		return x
	}
	for rightMin.Left != nil {
		rightMin = rightMin.Left
	}
	return rightMin
}
