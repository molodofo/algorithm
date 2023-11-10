package datastructure

type Queue []any

func (q *Queue) Push(item any) {
	*q = append(*q, item)
}

func (q *Queue) Pop() any {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
