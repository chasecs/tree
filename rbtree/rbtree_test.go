package rbtree

import (
	"math/rand"
	"testing"
	"time"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func TestFind(t *testing.T) {

	tree := New()

	tree.Insert(key(1), "val_1")
	tree.Insert(key(3), "val_3")
	tree.Insert(key(4), "val_4")
	tree.Insert(key(6), "val_6")
	tree.Insert(key(7), "val_7")

	n := tree.Find(key(4))
	if n.Value != "val_4" {
		t.Error("Error value")
	}
	n.Value = "v_4"
	if n.Value != "v_4" {
		t.Error("Error value modify")
	}
}

func TestBalance(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i++ {
		tree.Insert(key(i), i)
		if !tree.IsBalance() {
			t.Error("tree balance Error")
		}
	}
}

func TestInsert(t *testing.T) {

	tree := New()

	tree.Insert(key(1), "val_1")
	tree.Insert(key(3), "val_3")
	tree.Insert(key(4), "val_4")

	root := tree.Root
	if root.Value != "val_3" {
		t.Error("Insert Position Error")
	}
	if !root.IsBlack {
		t.Error("Insert Color Error")
	}

	node4 := root.Right
	if node4.Value != "val_4" || node4.IsBlack {
		t.Error("Insert Position Error")
	}
	node1 := root.Left
	if node1.Value != "val_1" || node1.IsBlack {
		t.Error("Insert Position Error")
	}

	tree.Insert(key(6), "val_6")
	root = tree.Root
	if root.Value != "val_3" {
		t.Error("Insert Position Error", root)
	}
	if !root.IsBlack {
		t.Error("Insert Color Error")
	}
	node4 = root.Right
	if node4.Value != "val_4" || !node4.IsBlack {
		t.Error("Insert Position Error", node4)
	}
	node1 = root.Left
	if node1.Value != "val_1" || !node1.IsBlack {
		t.Error("Insert Position Error", node1)
	}

	tree.Insert(key(7), "val_7")
	root = tree.Root
	node7 := root.Right.Right

	if node7.Value != "val_7" || node7.IsBlack {
		t.Error("Insert Position Error", node7)
	}
}

func TestDeletion(t *testing.T) {
	tree := New()
	rand.Seed(time.Now().UnixNano())
	count := rand.Perm(100)
	for _, k := range count {
		tree.Insert(key(k), k)
	}

	if !tree.IsBalance() {
		t.Error("tree balance Error")
	}
	for i, k := range count {
		n := tree.Delete(key(k))
		if !tree.IsBalance() {
			t.Error("tree balance Error", n, i)
		}
	}
}
