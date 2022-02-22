package ch1

type node[Item any] struct {
	item *Item
	next *node[Item]
}
