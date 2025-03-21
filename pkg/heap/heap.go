package heap

func Heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, n, largest)
	}
}

func BuildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Heapify(arr, i, 0)
	}
}
