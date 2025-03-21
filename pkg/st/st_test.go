package st

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestSt(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	st := NewST(a)

	for i := 0; i < m; i++ {
		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())
		l--
		r--
		fmt.Println(st.Query(l, r))
	}
}
