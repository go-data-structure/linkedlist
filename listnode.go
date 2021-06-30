package linkedlist

type listNode interface {
	set(interface{})
	get() interface{}
	getNext() listNode
	setNext(node listNode)
}

// newListNode 创建链表节点
func newListNode(t Type, val interface{}) listNode {
	switch t {
	case IntType:
		node := intListNode{val: val.(int)}
		return &node
	case StringType:
		node := stringListNode{val: val.(string)}
		return &node
	case InterfaceType:
		node := interfaceListNode{val: val}
		return &node
	default:
		node := interfaceListNode{val: val}
		return &node
	}
}

type interfaceListNode struct {
	val  interface{}
	next listNode
}

func (ln *interfaceListNode) set(val interface{}) {
	ln.val = val
}

func (ln *interfaceListNode) get() interface{} {
	return ln.val
}

func (ln *interfaceListNode) getNext() listNode {
	return ln.next
}

func (ln *interfaceListNode) setNext(node listNode) {
	ln.next = node
}

type intListNode struct {
	val  int
	next listNode
}

func (ln *intListNode) set(val interface{}) {
	ln.val = val.(int)
}

func (ln *intListNode) get() interface{} {
	return ln.val
}

func (ln *intListNode) getNext() listNode {
	return ln.next
}

func (ln *intListNode) setNext(node listNode) {
	ln.next = node
}

type stringListNode struct {
	val  string
	next listNode
}

func (ln *stringListNode) set(val interface{}) {
	ln.val = val.(string)
}

func (ln *stringListNode) get() interface{} {
	return ln.val
}

func (ln *stringListNode) getNext() listNode {
	return ln.next
}

func (ln *stringListNode) setNext(node listNode) {
	ln.next = node
}
