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

func (g *Graph) AddVertex(value int) {
	if contains(g.vertices, value) {
		fmt.Printf("vertex %v already exists", value)
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
		fmt.Println("vertex not found")
		return
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("existing edge %v->%v", from, to)
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
func (g *Graph) getVertex(value int) *Vertex {
	for _, v := range g.vertices {
		if v.key == value {
			return v
		}
	}
	return nil
}

type Node struct {
	data *Vertex
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
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
	if q.front == nil {
		q.rear = nil
	}
	return toRemove
}
func (g *Graph) BFS(start int) {
	startVertex := g.getVertex(start)
	q := Queue{}
	q.Enqueue(startVertex)
	startVertex.visited = true
	for !q.IsEmpty() {
		current := q.front.data
		fmt.Printf(" %d ", current.key)
		for _, neighbour := range current.adjacent {
			if !neighbour.visited {
				q.Enqueue(neighbour)
				neighbour.visited = true
			}
		}
		q.Dequeue()
	}
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
func (g *Graph) DFS(start int) {
	startVertex := g.getVertex(start)
	if startVertex == nil {
		fmt.Printf("vertex not found")
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
			if !neighbour.visited {
				s.Push(neighbour)
			}
		}
	}
}
func (g *Graph) hasCycle() bool {
	s := Stack{}
	for _, v := range g.vertices {
		if g.hasCycleDFS(v, &s) {
			return true
		}
	}
	return false
}
func (g *Graph) hasCycleDFS(vertex *Vertex, s *Stack) bool {
	vertex.visited = true
	vertex.inStack = true
	for _, neighbour := range vertex.adjacent {
		if !neighbour.visited {
			if g.hasCycleDFS(neighbour, s) {
				return true
			}
		} else if vertex.inStack {
			return true
		}
	}
	vertex.inStack = false
	s.Push(vertex)
	return false
}
func main() {
	g := Graph{}
	for i := 0; i < 6; i++ {
		g.AddVertex(i)
	}
	// g.AddEdge(1, 2)
	// g.AddEdge(0, 3)
	// g.AddEdge(1, 4)
	// g.AddEdge(2, 4)
	// g.AddEdge(3, 5)
	// g.AddEdge(4, 6)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(1, 5)
	g.Print()
	fmt.Println()
	fmt.Println(g.hasCycle())

}
