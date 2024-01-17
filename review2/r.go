package main

func MergeSort(arr []int)[]int{
	if len(arr)==1{
		return arr
	}
	m:=len(arr)/2
	left:=MergeSort(arr[:m])
	right:=MergeSort(arr[m:])
	return Merge(left,right)
}
func Merge(left,right []int)[]int{
	i:=0
	j:=0
	k:=0
	MergedArray:=make([]int,len(left)+len(right))
	for i<len(left)&&j<len(right)
}
func QuickSort(arr []int,low,high int){
	if low>=high{
		return
	}
	s:=low
	e:=high
	m:=s+(e-s)/2
	p:=arr[m]
	for s<=e{
		for arr[s]<p{
			s++
		}
		for arr[e]>p{
			e--
		}
	}
}
func commonElements(a,b []int){
for i:=0;i<len(a);i++{
	h.insert(a[i])
}
for j:=0;j<len(b);j++{
	h.sea
}
}
type Node struct{
	data int
	next *Node
}
type Stack struct{
	top *Node
}
func(s *Stack)Push(value int){
	newNode:=&Node{data:value}
	newNode.next=s.top
	s.top=newNode
}
func SelectionSort(arr []int){
	for i:=0;i<len(arr);i++{
		last:=len(arr)-i-1
		max:=getMax(arr,last)
		arr[max],arr[last]=arr[last],arr[max]
	}
}
func getMax(arr []int,last)int{
	max:=0
	for i:=0;i<len(arr);i++{
		if arr[i]>arr[max]{
			max=i
		} 
	}
	return max
}
func BubbleSort(arr []int){
	swapped:=false
	for i:=0;i<len(arr);i++{
		for j:=1;j<len(arr)-i;j++{
			if arr[j-1]>arr[j]{
				arr[j-1],arr[j]=arr[j],arr[j-1]
				swapped=true
			}

		}
		if !swapped{
			break
		}
	}
}