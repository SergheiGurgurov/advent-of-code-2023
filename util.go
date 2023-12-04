package util

import (
	"fmt"
	"os"
)

func CharToInt[T rune | byte](c T) int {
	return int(c - '0')
}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func First[T any](list []T, fn func(T) bool) T {
	for _, item := range list {
		if fn(item) {
			return item
		}
	}
	panic("No item found")
}

func Last[T any](list []T, fn func(T) bool) T {
	for i := len(list) - 1; i >= 0; i-- {
		if fn(list[i]) {
			return list[i]
		}
	}
	panic("No item found")
}

func ReadFileSync(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func Every[T any](ts []T, fn func(T) bool) bool {
	for i := range ts {
		if !fn(ts[i]) {
			return false
		}
	}
	return true
}

// managing optional return in generic functions
type Option[T any] struct {
	Value  T
	IsSome bool
}

type LinkedListNode[T any] struct {
	Next  *LinkedListNode[T]
	Value T
}

type LinkedList[T any] struct {
	Head *LinkedListNode[T]
	Tail *LinkedListNode[T]
}

func (list *LinkedList[T]) Add(value T) {
	node := LinkedListNode[T]{Value: value}

	if list.Head == nil {
		list.Head = &node
		list.Tail = &node
	} else {
		list.Tail.Next = &node
		list.Tail = &node
	}
}

func (list *LinkedList[T]) Shift() Option[*LinkedListNode[T]] {
	if list.Head == nil {
		return Option[*LinkedListNode[T]]{IsSome: false}
	}

	value := list.Head
	list.Head = list.Head.Next

	return Option[*LinkedListNode[T]]{Value: value, IsSome: true}
}

func ToLinkedList[T any](arr []T) *LinkedList[T] {
	list := LinkedList[T]{}

	for _, elem := range arr {
		list.Add(elem)
	}

	return &list
}

func (list *LinkedList[T]) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Printf("%v ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}

func InsertNodeAfter[T any](nodeBefore *LinkedListNode[T], nodeAfter *LinkedListNode[T]) {
	nodeAfter.Next = nodeBefore.Next
	nodeBefore.Next = nodeAfter
}
