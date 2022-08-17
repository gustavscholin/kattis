package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(key string) {
	t.root = t.root.add(key)
}

func (t *AVLTree) Remove(key string) {
	t.root = t.root.remove(key)
}

func (t *AVLTree) Update(oldKey string, newKey string) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey)
}

func (t *AVLTree) Search(key string) (node *AVLNode) {
	return t.root.search(key)
}

func (t *AVLTree) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

func (t *AVLTree) GetRank(key string) int {
	return t.root.getRank(key, 0)
}

// AVLNode structure
type AVLNode struct {
	key string

	// height counts nodes (not edges)
	height int
	size   int
	left   *AVLNode
	right  *AVLNode
}

// Adds a new node
func (n *AVLNode) add(key string) *AVLNode {
	if n == nil {
		return &AVLNode{key, 1, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key)
	} else if key > n.key {
		n.right = n.right.add(key)
	}
	n.size++
	return n.rebalanceTree()
}

// Removes a node
func (n *AVLNode) remove(key string) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			// delete smallest node that we replaced
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}

	}
	n.size--
	return n.rebalanceTree()
}

// Searches for a node
func (n *AVLNode) search(key string) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

// Displays nodes left-depth first (used for debugging)
func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, ": ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) getSize() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// Checks if node is balanced and rebalance
func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.size = 1 + n.left.getSize() + n.right.getSize()
	newRoot.size = 1 + newRoot.left.getSize() + newRoot.right.getSize()

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.size = 1 + n.left.getSize() + n.right.getSize()
	newRoot.size = 1 + newRoot.left.getSize() + newRoot.right.getSize()

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Finds the smallest child (based on the key) for the current node
func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

// Returns max number - TODO: std lib seemed to only have a method for floats!
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *AVLNode) getRank(key string, rank int) int {
	if n == nil {
		return 0
	}

	if key < n.key {
		if n.left == nil {
			return rank
		}
		return n.left.getRank(key, rank)
	} else if key > n.key {
		if n.right == nil {
			return rank + 1
		}
		if n.left == nil {
			return n.right.getRank(key, rank+1)
		}
		return n.right.getRank(key, rank+1+n.left.size)
	} else {
		return rank
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandParts := strings.Split(scanner.Text(), " ")
	command, _ := strconv.Atoi(commandParts[0])
	genders := make(map[string]int)
	boys := new(AVLTree)
	girls := new(AVLTree)
	for command != 0 {
		switch command {
		case 1:
			gender, _ := strconv.Atoi(commandParts[2])
			name := commandParts[1]
			genders[name] = gender
			if gender == 1 {
				boys.Add(name)
			} else if gender == 2 {
				girls.Add(name)
			}
		case 2:
			name := commandParts[1]
			gender, _ := genders[name]
			if gender == 1 {
				boys.Remove(name)
			} else if gender == 2 {
				girls.Remove(name)
			}
			delete(genders, name)
		case 3:
			gender, _ := strconv.Atoi(commandParts[3])
			start := commandParts[1]
			end := commandParts[2]
			var matches int
			switch gender {
			case 0:
				matches = boys.GetRank(end) - boys.GetRank(start)
				matches += girls.GetRank(end) - girls.GetRank(start)
			case 1:
				matches = boys.GetRank(end) - boys.GetRank(start)
			case 2:
				matches = girls.GetRank(end) - girls.GetRank(start)
			}
			fmt.Println(matches)
		}
		scanner.Scan()
		commandParts = strings.Split(scanner.Text(), " ")
		command, _ = strconv.Atoi(commandParts[0])
	}
}
