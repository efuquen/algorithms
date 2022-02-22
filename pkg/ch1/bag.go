package ch1

import "github.com/efuquen/algorithms/pkg/algs4/iter"

type Bag[Item any] struct {
	iter.Iterable[Item]

	first *node[Item]
	n     int
}

func NewBag[Item any]() Bag[Item] {
	return Bag[Item]{
		first: nil,
		n:     0,
	}
}

func (b *Bag[Item]) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag[Item]) Size() int {
	return b.n
}

func (b *Bag[Item]) Add(item *Item) {
	oldFirst := b.first
	b.first = &node[Item]{
		item: item,
		next: oldFirst,
	}
	b.n++
}

func (b *Bag[Item]) Iter() iter.Iterator[*Item] {
	iter := make(iter.Iterator[*Item])
	go func() {
		for n := b.first; n != nil; n = n.next {
			iter <- n.item
		}
		close(iter)
	}()
	return iter
}
