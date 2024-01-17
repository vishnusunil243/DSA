package main

import "fmt"

type Vertex struct {
	key      int
	adjacent []*Vertex
	visited  bool
	inStack  bool
}
type Graph struct {
	vertices []*Vertex
}
type Node struct {
	data *Vertex
	next *Node
}
type Stack struct {
	top *Node
}

//	func (g *Graph) numberOfIslands(startKey int) int {
//		startVertex := g.getVertex(startKey)
//		if startVertex == 0 {
//			return -1
//		}
//		s := Stack{}
//		s.Push(startVertex)
//		nofislands := 0
//		for !s.IsEmpty() {
//			current := s.top.data
//			if current != 1 {
//				nofislands += 1
//			}
//			for _, neighbour := range current.adjacent {
//				if !neighbour.visited {
//					s.Push(neighbour)
//				}
//			}
//		}
//	}
func reverseString(s string) string {
	b := []byte(s)
	reverse(b, 0, len(b)-1)
	return string(b)
}
func reverse(s []byte, st, e int) {
	if st >= e {
		return
	}
	s[st], s[e] = s[e], s[st]
	st++
	e--
	reverse(s, st, e)
}
func (g *Graph) getVertex(startkey int) *Vertex {
	for _, v := range g.vertices {
		if v.key == startkey {
			return v
		}
	}
	return nil
}
func main() {
	fmt.Println(reverseString("vishnu"))
}
