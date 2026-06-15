package main

import "fmt"

type Node struct {
	left  *Node
	data  int
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) addNode(data int) {
	node := Node{left: nil, data: data, right: nil}
	if bst.root == nil {
		bst.root = &node
		return
	}

	curr := bst.root
	for {
		if curr.data <= data {
			if curr.right == nil {
				curr.right = &node
				return
			}
			curr = curr.right
		} else {
			if curr.left == nil {
				curr.left = &node
				return
			}
			curr = curr.left
		}
	}

}

func (bst *BinarySearchTree) inOrder() {
	var traverse func(n *Node)
	traverse = func(n * Node) {
		if n == nil {
			return
		}
		traverse(n.left)
		fmt.Printf("%d =>", n.data)
		traverse(n.right)
	}
	traverse(bst.root)
}


func (bst *BinarySearchTree) preOrder() {
        var traverse func(n *Node)
        traverse = func(n * Node) {
                if n == nil {
                        return
                }
		fmt.Printf("%d =>", n.data)
                traverse(n.left)
                traverse(n.right)
        }
        traverse(bst.root)
}

func (bst *BinarySearchTree) postOrder() {
        var traverse func(n *Node)
        traverse = func(n * Node) {
                if n == nil {
                        return
                }
                traverse(n.left)
                traverse(n.right)
		fmt.Printf("%d =>", n.data)
        }
        traverse(bst.root)
}

func (bst *BinarySearchTree) levelOrder() {
	queue := [] *Node {bst.root}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Printf(" %d => ", curr.data)
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}

func main() {
	bst := &BinarySearchTree{}
	for i := range 10 {
		bst.addNode(i+1)
	}
	fmt.Println("\n================In Order ===============")
	bst.inOrder()
	fmt.Println("\n================Pre Order ===============")
	bst.preOrder()
	fmt.Println("\n================Post Order ===============")
	bst.postOrder()
	fmt.Println("\n================Level Order ===============")
	bst.levelOrder()
}
