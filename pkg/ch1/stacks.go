package ch1

import "github.com/efuquen/algorithms/pkg/algs4/iter"

/*
 *  FixedCapacityStackOfStrings start.
 */

type FixedCapacityStackOfStrings struct {
	a []string
	n int
}

func NewFixedCapacityStackOfStrings(capacity int) FixedCapacityStackOfStrings {
	return FixedCapacityStackOfStrings{
		a: make([]string, capacity),
		n: 0,
	}
}

func (s *FixedCapacityStackOfStrings) IsEmpty() bool {
	return s.n == 0
}

func (s *FixedCapacityStackOfStrings) Size() int {
	return s.n
}

func (s *FixedCapacityStackOfStrings) Push(item string) {
	s.a[s.n] = item
	s.n += 1
}

func (s *FixedCapacityStackOfStrings) Pop() string {
	s.n -= 1
	item := s.a[s.n]
	return item
}

/*
 *  ResizingArrayStack[Item] start.
 */

type ResizingArrayStack[Item any] struct {
	iter.Iterable[Item]

	a []*Item
	n int
}

func NewResizingArrayStack[Item any]() ResizingArrayStack[Item] {
	return ResizingArrayStack[Item]{
		a: make([]*Item, 1),
		n: 0,
	}
}

func (s *ResizingArrayStack[Item]) IsEmpty() bool {
	return s.n == 0
}

func (s *ResizingArrayStack[Item]) Size() int {
	return s.n
}

func (s *ResizingArrayStack[Item]) resize(max int) {
	temp := make([]*Item, max)
	for i := 0; i < s.n; i++ {
		temp[i] = s.a[i]
	}
	s.a = temp
}

func (s *ResizingArrayStack[Item]) Push(item *Item) {
	if s.n == len(s.a) {
		s.resize(2 * len(s.a))
	}
	s.a[s.n] = item
	s.n += 1
}

func (s *ResizingArrayStack[Item]) Pop() *Item {
	s.n -= 1
	item := s.a[s.n]
	s.a[s.n] = nil
	if s.n > 0 && s.n == len(s.a)/4 {
		s.resize(len(s.a) / 2)
	}
	return item
}

func (s *ResizingArrayStack[Item]) Iter() iter.Iterator[*Item] {
	iter := make(iter.Iterator[*Item])
	go func() {
		for i := s.n - 1; i >= 0; i-- {
			iter <- s.a[i]
		}
		close(iter)
	}()
	return iter
}

/*
 *  Stack[Item] start.
 */

type Stack[Item any] struct {
	iter.Iterable[Item]

	first *node[Item]
	n     int
}

func NewStack[Item any]() Stack[Item] {
	return Stack[Item]{
		first: nil,
		n:     0,
	}
}

func (s *Stack[Item]) IsEmpty() bool {
	return s.first == nil
}

func (s *Stack[Item]) Size() int {
	return s.n
}

func (s *Stack[Item]) Push(item *Item) {
	oldFirst := s.first
	s.first = &node[Item]{
		item: item,
		next: oldFirst,
	}
	s.n++
}

func (s *Stack[Item]) Pop() *Item {
	item := s.first.item
	s.first = s.first.next
	s.n--

	return item
}

func (s *Stack[Item]) Iter() iter.Iterator[*Item] {
	iter := make(iter.Iterator[*Item])
	go func() {
		for n := s.first; n != nil; n = n.next {
			iter <- n.item
		}
		close(iter)
	}()
	return iter
}
