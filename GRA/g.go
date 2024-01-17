package main

import "fmt"

type Vertex struct {
	key      int
	adjacent []*Vertex
	visited  bool
}
type Graph struct {
	vertices []*Vertex
}
type Node struct {
	data *Vertex
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) Push(value *Vertex) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
func (s *Stack) Pop() *Vertex {
	if s.IsEmpty() {
		return nil
	}
	toRemove := s.top.data
	s.top = s.top.next
	return toRemove
}

func (q *Queue) Enqueue(value *Vertex) {
	newNode := &Node{data: value}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
func (q *Queue) Dequeue() *Vertex {
	if q.IsEmpty() {
		return nil
	}
	toRemove := q.front.data
	q.front = q.front.next
	return toRemove
}

func (g *Graph) AddVertex(value int) {
	if contains(g.vertices, value) {
		fmt.Printf("vertex already exists")
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: value})
}
func contains(arr []*Vertex, value int) bool {
	for _, v := range arr {
		if v.key == value {
			return true
		}
	}
	return false
}
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("vertex not found")
		return
	}
	if contains(fromVertex.adjacent, to) {
		fmt.Printf("existing edge %v -> %v", from, to)
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vr := range v.adjacent {
			fmt.Printf(" %v ", vr.key)
		}
	}
}
func (g *Graph) getVertex(i int) *Vertex {
	for _, v := range g.vertices {
		if v.key == i {
			return v
		}
	}
	return nil
}
func (g *Graph) BFS(startKey int) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		fmt.Print("vertex not found")
		return
	}
	q := Queue{}
	q.Enqueue(startVertex)
	startVertex.visited = true
	for !q.IsEmpty() {
		current := q.front.data
		fmt.Printf(" %v ", current.key)

		for _, neighbour := range current.adjacent {
			if !neighbour.visited {
				neighbour.visited = true
				q.Enqueue(neighbour)
			}
		}
		q.Dequeue()
	}
}
func (g *Graph) DFS(starKey int) {
	startVertex := g.getVertex(starKey)
	if startVertex == nil {
		fmt.Print("vertex not found")
		return
	}
	s := Stack{}
	s.Push(startVertex)
	for !s.IsEmpty() {
		current := s.top.data
		if !current.visited {
			fmt.Printf(" %v ", current.key)
			current.visited = true
		}
		s.Pop()
		for _, neighbour := range current.adjacent {
			s.Push(neighbour)
		}
	}
}
func main() {
	g := Graph{}
	arr := []int{1, 2, 3, 4, 5}
	for _, num := range arr {
		g.AddVertex(num)
	}
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(1, 5)
	g.Print()
	fmt.Println()
	g.BFS(1)
}
