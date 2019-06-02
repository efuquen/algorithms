package ch1

type UF interface {
	Connected(p int, q int) bool
	Find(p int) int
	Union(p int, q int)
}
