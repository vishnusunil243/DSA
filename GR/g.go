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
	if q.front == nil {
		q.rear = nil
	}
	return toRemove
}
func (g *Graph) AddVertex(value int) {
	if contains(g.vertices, value) {
		fmt.Printf("existing vertex %v", value)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: value})
}
func contains(arr []*Vertex, k int) bool {
	for _, i := range arr {
		if i.key == k {
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
	} else if contains(fromVertex.adjacent, to) || contains(toVertex.adjacent, from) {
		fmt.Printf("existing edge %v => %v", from, to)
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}
func (g *Graph) getVertex(value int) *Vertex {
	for _, i := range g.vertices {
		if i.key == value {
			return i
		}
	}
	return nil
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vr := range v.adjacent {
			fmt.Printf(" %v ", vr.key)
		}
	}
}
func (g *Graph) BFS(startKey int) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		fmt.Printf("vertex not found")
		return
	}
	q := Queue{}
	q.Enqueue(startVertex)
	fmt.Print(startVertex)
	startVertex.visited = true
	for !q.IsEmpty() {
		fmt.Print("hii")
		current := q.front.data
		fmt.Printf(" %v =>", current.key)
		for _, neighbour := range current.adjacent {
			if !neighbour.visited {
				neighbour.visited = true
				q.Enqueue(neighbour)
			}
		}
		q.Dequeue()
	}
}
func (g *Graph) DFS(startKey int) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		fmt.Printf("vertex not found")
		return
	}
	s := Stack{}
	s.Push(startVertex)
	for !s.IsEmpty() {
		current := s.top.data
		if !current.visited {
			fmt.Printf("%v => ", current.key)
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
func (g *Graph) dfs(startKey int) {
	startVertex := g.getVertex(startKey)
	if startVertex == nil {
		fmt.Println("Vertex not found")
		return
	}

	stack := Stack{}
	stack.Push(startVertex)

	for !stack.IsEmpty() {
		current := stack.top.data
		if !current.visited {
			fmt.Printf("%v ", current.key)
			current.visited = true
		}
		stack.Pop()

		// Push unvisited adjacent vertices onto the stack
		for _, neighbor := range current.adjacent {
			if !neighbor.visited {
				stack.Push(neighbor)
			}
		}
	}
}
func main() {
	g := Graph{}
	for i := 0; i < 6; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(0, 2)
	g.AddEdge(4, 5)
	g.AddEdge(3, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(1, 2)
	g.Print()
	fmt.Println()
	g.DFS(0)
}
