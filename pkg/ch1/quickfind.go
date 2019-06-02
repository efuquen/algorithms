package ch1

type QuickFind struct {
	*BaseUF
}

func NewQuickFind(n int) QuickFind {
	var qf = QuickFind { }
	qf.count = n
	qf.id = make([]int, n)
	for i, _ := range qf.id {
		qf.id[i] = i
	}
	return qf
}

func (qf QuickFind) Find(p int) int {
	return qf.id[p]
}

func (qf QuickFind) Union(p int, q int) {
	var pId = qf.Find(p)
	var qId = qf.Find(q)
	if pId == qId {
		return
	}

	for i, id := range qf.id {
		if id == pId {
			qf.id[i] = qId
		}
	}
}