package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func (m *MaxHeap) insert(value int) {
	m.arr = append(m.arr, value)
	m.heapifyUp(len(m.arr) - 1)
}
func (m *MaxHeap) heapifyUp(index int) {
	for m.arr[parent(index)] < m.arr[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}
func (m *MaxHeap) extract() int {
	extracted := m.arr[0]
	l := len(m.arr) - 1
	if len(m.arr) == 0 {
		fmt.Println("cannot extract from empty max heap")
		return -1
	}
	m.arr[0] = m.arr[l]
	m.arr = m.arr[:l]
	m.heapifyDown(0)
	return extracted
}
func (m *MaxHeap) heapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.arr[l] > m.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if m.arr[index] < m.arr[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return (2 * i) + 1
}
func rightChild(i int) int {
	return (2 * i) + 2
}
func (m *MaxHeap) swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}
func main() {
	m := MaxHeap{}
	arr := []int{1, 22, 44, 55, 606, 33}
	for _, num := range arr {
		m.insert(num)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(m.extract())
		fmt.Println(m)
	}
}
