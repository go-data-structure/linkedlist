package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := New(StringType)
	l.Add("dddd")
	l.Add("cccc")
	l.Add("9")

	fmt.Println(l.Len())

	err := l.GetAll(func(val interface{}) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
