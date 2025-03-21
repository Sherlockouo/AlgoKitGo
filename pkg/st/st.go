package st

import "math/bits"

type ST [][]int

// a 的下标从 0 开始
func NewST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
		// if i != 2 {
		// }
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = st.Op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// 查询区间 [l,r)    0 <= l < r <= n
func (st ST) Query(l, r int) int {
	k := bits.Len32(uint32(r-l)) - 1
	return st.Op(st[l][k], st[r-1<<k][k])
}

// min, max, gcd, ...
func (ST) Op(a, b int) int { return max(a, b) }
