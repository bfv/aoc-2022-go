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

var cnode *node
var root *node

func main() {

	start := time.Now()

	day := "day07"
	var a, b int

	strs := aoc.GetStringArray("input.txt")

	root = createNode("/", nil, 0, "d")
	parseInput(strs)
	calculateSize(root)

	a = findWithMax(root, 100000)
	b = findSmallest(30000000-(70000000-root.size), root, root.size)

	elapsed := time.Since(start)
	fmt.Printf("%v, a: %v, b: %v, time: %v", day, a, b, elapsed)
}

func parseInput(strs []string) {
	for _, s := range strs {
		line := strings.Split(s, " ")
		switch line[0] {
		case "$":
			if line[1] == "cd" {
				cnode = changeDir(line[2])
			}
		case "dir":
			addDir(line[1])
		default:
			addFile(lib.Atoi(line[0]), line[1])
		}
	}
}

func addDir(name string) {
	dir := createNode(name, cnode, 0, "d")
	cnode.childs = append(cnode.childs, dir)
}

func addFile(size int, name string) {
	file := createNode(name, cnode, size, "f")
	cnode.childs = append(cnode.childs, file)
}

func changeDir(name string) *node {
	var dir *node

	if name == "/" {
		dir = root
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

func createNode(name string, parent *node, size int, ntype string) *node {
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

	if node.ntype == "f" { // no files
		return 0
	}

	size := 0
	if node.size <= maxSize {
		size = node.size
	}
	for _, child := range node.childs {
		size += findWithMax(child, maxSize)
	}

	return size
}

func findSmallest(minSize int, node *node, smallest int) int {

	if node.ntype == "f" {
		return smallest
	}

	if node.size < smallest && node.size >= minSize {
		smallest = node.size
	}
	for _, child := range node.childs {
		smallest = findSmallest(minSize, child, smallest)
	}

	return smallest
}
