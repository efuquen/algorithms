package ch1

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
