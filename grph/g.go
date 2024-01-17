package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key       int
	adjacent  []*Vertex
	isVisited bool
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("%v already exists", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}
func contains(vertices []*Vertex, k int) bool {
	for _, v := range vertices {
		if v.key == k {
			return true
		}
	}
	return false
}
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("vertex not found")
		fmt.Println(err.Error())
		return
	} else if contains(fromVertex.adjacent, to) || contains(toVertex.adjacent, from) {
		fmt.Printf("existing edge %v -> %v", from, to)
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}
func (g *Graph) RemoveEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("vertex not found")
		return
	} else if !contains(fromVertex.adjacent, to) || !contains(toVertex.adjacent, from) {
		fmt.Println("edge not found")
		return
	}
	fromVertex.adjacent = removeEdge(fromVertex.adjacent, to)
	toVertex.adjacent = removeEdge(toVertex.adjacent, from)
}
func removeEdge(adjacent []*Vertex, target int) []*Vertex {
	var result []*Vertex
	for _, v := range adjacent {
		if v.key != target {
			result = append(result, v)
		}
	}
	return result
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vr := range v.adjacent {
			fmt.Printf(" %v ", vr.key)
		}
	}
}
func main() {
	g := Graph{}
	for i := 0; i < 6; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.RemoveEdge(2, 4)
	g.Print()
}
