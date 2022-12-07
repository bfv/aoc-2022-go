package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bfv/aoc2022-go/aoc"
	"github.com/bfv/aoc2022-go/lib"
)

type node struct {
	name   string
	parent *node
	size   int
	ntype  string
	childs []*node
}

func main() {

	start := time.Now()

	day := "day07"
	var a, b int

	strs := aoc.GetStringArray("input.txt")

	root := createNode("/", nil, "d", 0)
	parseInput(strs, root)
	calculateSize(root)

	if false {
		displayFS(root, 0)
	}

	a = findWithMax(root, 100000)
	b = findSmallest(30000000-(70000000-root.size), root, root.size)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func parseInput(strs []string, root *node) {
	var cnode *node
	cnode = root

	for _, s := range strs {
		args := strings.Split(s, " ")
		switch args[0] {
		case "$":
			if args[1] == "cd" {
				cnode = changeDir(args[2], cnode)
			}
		case "dir":
			addNode(args[1], "d", 0, cnode)
		default:
			addNode(args[1], "f", lib.Atoi(args[0]), cnode)
		}
	}
}

func addNode(name string, ntype string, size int, parent *node) {
	node := createNode(name, parent, ntype, size)
	parent.childs = append(parent.childs, node)
}

func changeDir(name string, cnode *node) *node {
	var dir *node

	if name == "/" {
		dir = cnode
	} else if name == ".." {
		dir = cnode.parent
	} else {
		for _, n := range cnode.childs { // find node with name in cnode.childs
			if n.ntype == "d" && n.name == name {
				dir = n
				break
			}
		}
	}
	return dir
}

func createNode(name string, parent *node, ntype string, size int) *node {
	return &node{name, parent, size, ntype, []*node{}}
}

func calculateSize(node *node) int {
	size := 0
	if node.ntype == "f" {
		size = node.size
	} else {
		for _, child := range node.childs {
			size += calculateSize(child)
		}
		node.size = size
	}
	return size
}

func findWithMax(node *node, maxSize int) int {
	size := 0
	if node.ntype == "d" { // no files
		if node.size <= maxSize {
			size = node.size
		}
		for _, child := range node.childs {
			size += findWithMax(child, maxSize)
		}
	}
	return size
}

func findSmallest(minSize int, node *node, smallest int) int {
	if node.ntype == "d" {
		if node.size < smallest && node.size >= minSize {
			smallest = node.size
		}
		for _, child := range node.childs {
			smallest = findSmallest(minSize, child, smallest)
		}
	}
	return smallest
}

func displayFS(node *node, level int) {
	nodeStr := fmt.Sprintf("- %s", node.name)
	if node.ntype == "d" {
		nodeStr += " (dir)"
	} else {
		nodeStr += fmt.Sprintf(" (file, size=%v)", node.size)
	}
	fmt.Printf("%s%s\n", strings.Repeat(" ", level*2), nodeStr)

	for _, d := range node.childs {
		displayFS(d, level+1)
	}
}
