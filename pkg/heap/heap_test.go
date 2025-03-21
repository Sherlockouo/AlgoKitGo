package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{1, 6, 34, 5, 2, 7, 8, 9, 10}
	BuildMaxHeap(arr)
	fmt.Println(arr)
}
