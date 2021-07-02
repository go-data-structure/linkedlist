package linkedlist

type StringListNode struct {
	Val  string
	Next *StringListNode
}

type StringLinkedList struct {
	Len  int64
	Head *StringListNode
	Tail *StringListNode
}

func NewStringLinkedList() *StringLinkedList {
	return &StringLinkedList{Len: 0, Head: nil, Tail: nil}
}

func (l *StringLinkedList) Add(val string) {
	if l.Head == nil || l.Tail == nil {
		l.Head = &StringListNode{Val: val}
		l.Tail = l.Head
		l.Len += 1
		return
	}

	l.Tail.Next = &StringListNode{Val: val}
	l.Tail = l.Tail.Next
	l.Len += 1
	return
}

func (l *StringLinkedList) GetAll(f func(val string) error) error {
	node := l.Head
	for node != nil {
		if err := f(node.Val); err != nil {
			return err
		}
		node = node.Next
	}
	return nil
}
