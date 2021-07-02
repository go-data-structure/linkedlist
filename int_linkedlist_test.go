package linkedlist

import (
	"fmt"
	"testing"
)

func TestIntLinkedList(t *testing.T) {
	l := NewIntLinkedList()
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)

	if l.Len != 4 {
		t.Fatal("长度异常")
	}
	fmt.Println("length: ", l.Len)

	err := l.GetAll(func(val int) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
