package linkedlist

type IntListNode struct {
	Val  int
	Next *IntListNode
}

type IntLinkedList struct {
	Len  int64
	Head *IntListNode
	Tail *IntListNode
}

func NewIntLinkedList() *IntLinkedList {
	return &IntLinkedList{Len: 0, Head: nil, Tail: nil}
}

func (l *IntLinkedList) Add(val int) {
	if l.Head == nil || l.Tail == nil {
		l.Head = &IntListNode{Val: val}
		l.Tail = l.Head
		l.Len += 1
		return
	}

	l.Tail.Next = &IntListNode{Val: val}
	l.Tail = l.Tail.Next
	l.Len += 1
	return
}

func (l *IntLinkedList) GetAll(f func(val int) error) error {
	node := l.Head
	for node != nil {
		if err := f(node.Val); err != nil {
			return err
		}
		node = node.Next
	}
	return nil
}
