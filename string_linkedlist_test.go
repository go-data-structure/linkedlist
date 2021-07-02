package linkedlist

import (
	"fmt"
	"testing"
)

func TestStringLinkedList(t *testing.T) {
	l := NewStringLinkedList()
	l.Add("hello")
	l.Add("world")
	l.Add("wnanbei")

	if l.Len != 3 {
		t.Fatal("长度异常")
	}
	fmt.Println("length: ", l.Len)

	err := l.GetAll(func(val string) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
