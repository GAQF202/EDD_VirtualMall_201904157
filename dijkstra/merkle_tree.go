package dijkstra

import (
	"crypto/sha1"
	"encoding/hex"
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

var DotMerkleTree = "digraph { node [shape=box, style=\"filled\", fillcolor=\"#61e665\"];"

func (b Block) toStringBlock() string {
	d := []byte(b)[:]
	return string(d)
}

func printNode(node Node, level int) {
	DotMerkleTree += "node" + node.hash().String() + "[label=\"" + node.hash().String() + "\"];\n"
	//fmt.Println("node", node.hash(), "[label=\"", node.hash(), "\"];")
	if l, ok := node.left.(Node); ok {
		//fmt.Println("node", node.hash(), "->node", l.hash())
		DotMerkleTree += "node" + node.hash().String() + "->node" + l.hash().String() + "\n"
		printNode(l, level+1)
	} else if l, ok := node.left.(Block); ok {
		DotMerkleTree += "node" + l.hash().String() + "[label=\"" + l.hash().String() + "\\n" + l.toStringBlock() + "\"];\n"
		//fmt.Println("node", l.hash(), "[label=\"", l.hash(), "\\n", l, "\"];")
		DotMerkleTree += "node" + node.hash().String() + "->node" + l.hash().String() + "\n"
		//fmt.Println("node", node.hash(), "->node", l.hash())
	}
	if r, ok := node.right.(Node); ok {
		//fmt.Println("node", node.hash(), "->node", r.hash())
		DotMerkleTree += "node" + node.hash().String() + "->node" + r.hash().String() + "\n"
		printNode(r, level+1)
	} else if r, ok := node.right.(Block); ok {
		DotMerkleTree += "node" + r.hash().String() + "[label=\"" + r.hash().String() + "\\n" + r.toStringBlock() + "\"];\n"
		//fmt.Println("node", r.hash(), "[label=\"", r.hash(), "\\n", r, "\"];")
		DotMerkleTree += "node" + node.hash().String() + "->node" + r.hash().String() + "\n"
		//fmt.Println("node", node.hash(), "->node", r.hash())
	}
}
