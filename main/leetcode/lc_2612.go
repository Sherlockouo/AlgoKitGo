package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

// 给你一个整数 n 和一个在范围 [0, n - 1] 以内的整数 p ，它们表示一个长度为 n 且下标从 0 开始的数组 arr ，数组中除了下标为 p 处是 1 以外，其他所有数都是 0 。
//
// 同时给你一个整数数组 banned ，它包含数组中的一些位置。banned 中第 i 个位置表示 arr[banned[i]] = 0 ，题目保证 banned[i] != p 。
//
// 你可以对 arr 进行 若干次 操作。一次操作中，你选择大小为 k 的一个 子数组 ，并将它 翻转 。在任何一次翻转操作后，你都需要确保 arr 中唯一的 1 不会到达任何 banned 中的位置。换句话说，arr[banned[i]] 始终 保持 0 。
//
// 请你返回一个数组 ans ，对于 [0, n - 1] 之间的任意下标 i ，ans[i] 是将 1 放到位置 i 处的 最少 翻转操作次数，如果无法放到位置 i 处，此数为 -1 。
//
// 子数组 指的是一个数组里一段连续 非空 的元素序列。
// 对于所有的 i ，ans[i] 相互之间独立计算。
// 将一个数组中的元素 翻转 指的是将数组中的值变成 相反顺序 。
//
//
// 示例 1：
//
// 输入：n = 4, p = 0, banned = [1,2], k = 4
// 输出：[0,-1,-1,1]
// 解释：k = 4，所以只有一种可行的翻转操作，就是将整个数组翻转。一开始 1 在位置 0 处，所以将它翻转到位置 0 处需要的操作数为 0 。
// 我们不能将 1 翻转到 banned 中的位置，所以位置 1 和 2 处的答案都是 -1 。
// 通过一次翻转操作，可以将 1 放到位置 3 处，所以位置 3 的答案是 1
//
//

func minReverseOperations(n int, p int, banned []int, k int) []int {
	ban := make(map[int]bool)
	for _, b := range banned {
		ban[b] = true
	}

	sets := make([][]int, 2)
	for i := range n {
		if i != p && !ban[i] {
			sets[i%2] = append(sets[i%2], i)
		}
	}

	for i := range sets {
		sort.Ints(sets[i])
	}

	ans := make([]int, n)
	for i := range n {
		ans[i] = -1
	}
	ans[p] = 0
	// BFS
	queue := []int{p}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		mn := max(i-k+1, k-i-1)
		mx := min(i+k-1, 2*n-k-i-1)
		targetSet := sets[mx%2]
		toRemove := []int{}
		left := sort.SearchInts(targetSet, mn)
		right := sort.SearchInts(targetSet, mx+1)
		for j := left; j < right; j++ {
			v := targetSet[j]
			ans[v] = ans[i] + 1
			queue = append(queue, v)
			toRemove = append(toRemove, v)
		}
		for _, val := range toRemove {
			idx := sort.SearchInts(targetSet, val)
			if idx < len(targetSet) && targetSet[idx] == val {
				targetSet = slices.Delete(targetSet, idx, idx+1)
			}
		}

		sets[mx%2] = targetSet

	}
	return ans
}

func LC_2612(_r io.Reader, _w io.Writer) { // 方便测试，见 10C_test.go
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, p, k int
	Fscan(in, &n)
	Fscan(in, &p)
	Fscan(in, &k)

	ans := make([][]int, 0)

	Fprint(out, ans)
}

func main() { LC_2612(os.Stdin, os.Stdout) }
