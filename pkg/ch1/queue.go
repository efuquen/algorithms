package ch1

import (
	"github.com/efuquen/algorithms/pkg/algs4/iter"
)

type Queue[Item any] struct {
	iter.Iterable[Item]

	first *node[Item]
	last  *node[Item]
	n     int
}

func NewQueue[Item any]() Queue[Item] {
	return Queue[Item]{
		first: nil,
		last:  nil,
		n:     0,
	}
}

func (q *Queue[Item]) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue[Item]) Size() int {
	return q.n
}

func (q *Queue[Item]) Enqueue(item *Item) {
	oldLast := q.last
	q.last = &node[Item]{
		item: item,
		next: nil,
	}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.n++
}

func (q *Queue[Item]) Dequeue() *Item {
	oldFirst := q.first
	q.first = oldFirst.next
	q.n--
	if q.IsEmpty() {
		q.last = nil
	}

	return oldFirst.item
}

func (q *Queue[Item]) Iter() iter.Iterator[*Item] {
	iter := make(iter.Iterator[*Item])
	go func() {
		for n := q.first; n != nil; n = n.next {
			iter <- n.item
		}
		close(iter)
	}()
	return iter
}