package util

import "container/list"

type Stack struct {
	l list.List
}

func (s Stack) Push(x any) {
	s.l.PushBack(x)
}

func (s Stack) Pop() any {
	return s.l.Remove(s.l.Back())
}

func (s Stack) Top() any {
	return s.l.Back().Value
}

func (s Stack) IsEmpty() bool {
	return s.l.Len() == 0
}
