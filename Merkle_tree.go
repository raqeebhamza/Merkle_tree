package Merkle_tree

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"

	q "github.com/golang-collections/collections/queue"
)

type Node struct {
	transaction   string
	parentpointer *Node
	leftchild     *Node
	rightchild    *Node
	hash          []byte
}

// calculate_Hash is use to compute the hash of the node
func (b *Node) calculate_Hash() []byte {
	var data []byte
	data = bytes.Join([][]byte{[]byte(b.transaction)}, []byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}

func preorder_traversal(curr *Node) {
	if curr == nil {
		return
	}
	fmt.Print(curr.transaction, " ")
	preorder_traversal(curr.leftchild)
	preorder_traversal(curr.rightchild)
}

// BuidTree function will return the root of the tree after creating the merkle tree
func BuidTree(queue *q.Queue) *Node {

	var tree []*Node
	var nod *Node
	for queue.Len() > 0 {
		curr := &Node{
			transaction:   queue.Dequeue().(string),
			parentpointer: nil,
			leftchild:     nil,
			rightchild:    nil,
		}
		curr.hash = curr.calculate_Hash()
		tree = append([]*Node{curr}, tree...)
		nod = curr
	}
	if len(tree)%2 != 0 {
		tree = append([]*Node{nod}, tree...)
	}
	height := int(math.Log2(float64(len(tree))))
	old_len := len(tree)

	for i := 0; i < height; i += 1 {
		var temp_tree []*Node
		for j := old_len - 1; j > 0; j -= 2 {
			curr := &Node{
				transaction:   tree[j].transaction + tree[j-1].transaction,
				parentpointer: nil,
				leftchild:     nil,
				rightchild:    nil,
			}
			curr.hash = curr.calculate_Hash()
			temp_tree = append([]*Node{curr}, temp_tree...)
		}
		old_len = len(temp_tree)
		tree = append(temp_tree, tree...)
	}
	for k := 0; k < len(tree); k++ {
		if k != 0 {
			tree[k].parentpointer = tree[int((k-1)/2)]
		}
		if k < len(tree)-4 {
			tree[k].leftchild = tree[2*k+1]
			tree[k].rightchild = tree[2*k+2]
		}
	}
	return tree[0]
}
