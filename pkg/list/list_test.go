package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(8)
	l.PushBack(9)
	l.PushBack(10)
	l.PushBack(11)
	l.PushBack(12)
	l.PushBack(13)
	l.PushBack(14)
	l.PushBack(15)
	l.PushBack(16)
	l.PushBack(17)
	l.PushBack(18)
	l.PushBack(19)
	l.PushBack(20)
	l.PushBack(21)
	l.PushBack(22)
	l.PushBack(23)
	l.PushBack(24)
	l.PushBack(25)

	for i := 0; i < l.Len(); i++ {
		fmt.Println(l.Index(i).Value)
	}
}
