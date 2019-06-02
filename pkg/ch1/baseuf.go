package ch1

type BaseUF struct {
	UF
	id []int
	count int
}

func (b BaseUF) Connected(p int, q int) bool {
	return b.Find(p) == b.Find(q)
}