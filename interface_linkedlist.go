package linkedlist

type InterFaceListNode struct {
	Val  interface{}
	Next *InterFaceListNode
}

type InterFaceLinkedList struct {
	Len  int64
	Head *InterFaceListNode
	Tail *InterFaceListNode
}

func NewInterFaceLinkedList() *InterFaceLinkedList {
	return &InterFaceLinkedList{Len: 0, Head: nil, Tail: nil}
}

func (l *InterFaceLinkedList) Add(val interface{}) {
	if l.Head == nil || l.Tail == nil {
		l.Head = &InterFaceListNode{Val: val}
		l.Tail = l.Head
		l.Len += 1
		return
	}

	l.Tail.Next = &InterFaceListNode{Val: val}
	l.Tail = l.Tail.Next
	l.Len += 1
	return
}

func (l *InterFaceLinkedList) GetAll(f func(val interface{}) error) error {
	node := l.Head
	for node != nil {
		if err := f(node.Val); err != nil {
			return err
		}
		node = node.Next
	}
	return nil
}
