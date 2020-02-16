package main

import (
	"fmt"

	q "github.com/golang-collections/collections/queue"
)

func main() {
	transactions := []string{"A", "B", "C"}
	queue := q.New()
	for i, trans := range transactions {
		_ = i
		queue.Enqueue(trans)
	}
	var root *Node
	root = BuidTree(queue)
	fmt.Println("PreOrder Traversal Merkle_tree:")
	preorder_traversal(root)
}
