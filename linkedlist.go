package linkedlist

type LinkedList struct {
	listType Type
	len      int64
	head     listNode
}

func New(t Type) *LinkedList {
	return &LinkedList{listType: t, len: 0, head: nil}
}

func (l *LinkedList) Add(val interface{}) {
	if l.head == nil {
		l.head = newListNode(l.listType, val)
		l.len += 1
		return
	}

	node := l.head
	for {
		if node.getNext() == nil {
			node.setNext(newListNode(l.listType, val))
			l.len += 1
			return
		}
		node = node.getNext()
	}
}

func (l *LinkedList) GetAll(f func(val interface{}) error) error {
	node := l.head
	for node != nil {
		if err := f(node.get()); err != nil {
			return err
		}
		node = node.getNext()
	}
	return nil
}

// Len return the number of list nodes.
func (l *LinkedList) Len() int64 {
	return l.len
}
