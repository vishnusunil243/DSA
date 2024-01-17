package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key      int
	adjacent []*Vertex
	visited  bool
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
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}
func (q *Queue) Dequeue() *Vertex {
	if q.front == nil {
		return nil
	}
	toRemove := q.front.data
	q.front = q.front.next
	return toRemove
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
func (g *Graph) BFS(startKey int) {
	startVertx := g.getVertex(startKey)
	if startVertx == nil {
		fmt.Println("vertex not found")
		return
	}
	q := Queue{}
	q.Enqueue(startVertx)
	startVertx.visited = true
	for !q.IsEmpty() {
		current := q.front.data
		fmt.Printf("%v ", current.key)
		for _, neighbour := range current.adjacent {
			if !neighbour.visited {
				neighbour.visited = true
				q.Enqueue(neighbour)
			}
		}
		fmt.Println(q.Dequeue())
	}
}
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v->%v)", from, to)
		fmt.Println(err.Error())
		return
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if k == v.key {
			return g.vertices[i]
		}
	}
	return nil
}
func (g *Graph) DFS(startKey int) {
	startVertx := g.getVertex(startKey)
	if startVertx == nil {
		fmt.Println("vertex not found")
		return
	}
	g.dfsRecursive(startVertx)
}
func (g *Graph) dfsRecursive(v *Vertex) {
	if v == nil {
		return
	}
	fmt.Printf("%v ", v.key)
	v.visited = true
	for _, neighbour := range v.adjacent {
		if !neighbour.visited {
			g.dfsRecursive(neighbour)
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

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vr := range v.adjacent {
			fmt.Printf(" %v ", vr.key)
		}
	}
}
func main() {
	test := Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	// Add edges
	test.AddEdge(0, 1)
	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(2, 4)
	test.AddEdge(2, 5)
	test.AddEdge(3, 6)

	test.Print()
}
