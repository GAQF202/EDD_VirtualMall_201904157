package hashes

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Hash [20]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func hash(data []byte) Hash {
	return sha1.Sum(data)
}

type Hashable interface {
	hash() Hash
}

type Block string

func (b Block) hash() Hash {
	return hash([]byte(b)[:])
}

type EmptyBlock struct {
}

func (_ EmptyBlock) hash() Hash {
	return [20]byte{}
}

type Node struct {
	left  Hashable
	right Hashable
}

func (n Node) hash() Hash {
	var l, r [sha1.Size]byte
	l = n.left.hash()
	r = n.right.hash()
	return hash(append(l[:], r[:]...))
}

func BuildTree(parts []Hashable) []Hashable {
	var nodes []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			nodes = append(nodes, Node{left: parts[i], right: parts[i+1]})
		} else {
			nodes = append(nodes, Node{left: parts[i], right: EmptyBlock{}})
		}
	}
	if len(nodes) == 1 {
		return nodes
	} else {
		return BuildTree(nodes)
	}
}

func PrintTree(node Node) {
	printNode(node, 0)
}

func printNode(node Node, level int) {
	//fmt.Printf("(%d) %s %s\n", level, strings.Repeat("hola", level), node.hash())
	//fmt.Println(level, node.hash())
	//fmt.Println(node.hash(), "->", level+1)
	fmt.Println("node", node.hash(), "[label=\"", node.hash(), "\"];")
	//fmt.Println("PADRE", node.hash())
	if l, ok := node.left.(Node); ok {
		fmt.Println("node", node.hash(), "->node", l.hash())
		printNode(l, level+1)
	} else if l, ok := node.left.(Block); ok {
		fmt.Println("node", l.hash(), "[label=\"", l.hash(), "\\n", l, "\"];")
		fmt.Println("node", node.hash(), "->node", l.hash())
		//fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("hola", level+1), l.hash(), l)
	}
	if r, ok := node.right.(Node); ok {
		fmt.Println("node", node.hash(), "->node", r.hash())
		printNode(r, level+1)
	} else if r, ok := node.right.(Block); ok {
		fmt.Println("node", r.hash(), "[label=\"", r.hash(), "\\n", r, "\"];")
		fmt.Println("node", node.hash(), "->node", r.hash())
		//fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat("hola", level+1), r.hash(), r)
	}
}
