package linkedlist

import (
	"fmt"
	"testing"
)

func TestInterfaceLinkedList(t *testing.T) {
	l := NewInterFaceLinkedList()
	l.Add(3)
	l.Add("wnanbei")
	l.Add(5)
	l.Add("hello world")

	if l.Len != 4 {
		t.Fatal("长度异常")
	}
	fmt.Println("length: ", l.Len)

	err := l.GetAll(func(val interface{}) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
