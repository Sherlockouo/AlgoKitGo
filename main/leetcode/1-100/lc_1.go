package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func LC_2610(_r io.Reader, _w io.Writer) { // 方便测试，见 10C_test.go
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	var arr []int
	for i := 0; i < n; i++ {
		var x int
		Fscan(in, &x)
		arr = append(arr, x)
	}

	// way 1: use map
	ans := make([][]int, 0)
	// count := make([]int, 0)
	// for _, x := range arr {
	// 	count[x]++
	//    if count[x] > len(ans){
	//        ans = append(ans, []int{x})
	//    }else{
	//
	//    ans[count[x]-1] = append(ans[count[x]-1], x)
	//    }
	// }

	// ans := make([][]int64, ansLen)
	// for k, v := range count {
	// 	for i := 0; i < v; i++ {
	// 		ans[i] = append(ans[i], int64(k))
	// 	}
	// }

	Fprint(out, ans)
}

func main() { LC_2610(os.Stdin, os.Stdout) }
