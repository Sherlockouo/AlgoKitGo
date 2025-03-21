package skiplist

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// 测试比较函数
func intComparator(a, b interface{}) int {
	return a.(int) - b.(int)
}

func stringComparator(a, b interface{}) int {
	s1, s2 := a.(string), b.(string)
	switch {
	case s1 < s2:
		return -1
	case s1 > s2:
		return 1
	default:
		return 0
	}
}

/* 功能测试 */
func TestAsicOperations(t *testing.T) {
	t.Run("IntKeys", func(t *testing.T) {
		sl := New(intComparator)
		testBasicOperations(t, sl)
	})

	t.Run("StringKeys", func(t *testing.T) {
		sl := New(stringComparator)
		testBasicOperations(t, sl)
	})
}

func testBasicOperations(t *testing.T, sl *SkipList) {
	// 测试插入和查找
	sl.Insert(3, "three")
	sl.Insert(1, "one")
	sl.Insert(2, "two")

	// 正常查找
	if node := sl.Search(2); node == nil || node.value != "two" {
		t.Errorf("Search 2 failed")
	}

	// 查找不存在的键
	if sl.Search(4) != nil {
		t.Errorf("Found non-existent key")
	}

	// 测试删除
	sl.Delete(2)
	if sl.Search(2) != nil {
		t.Errorf("Delete failed")
	}

	// 测试重复插入
	sl.Insert(1, "new-one")
	if node := sl.Search(1); node.value != "new-one" {
		t.Errorf("Update failed")
	}

	// 测试长度
	if sl.Size() != 2 {
		t.Errorf("Size incorrect, got %d", sl.Size())
	}
}

func TestBoundaryCases(t *testing.T) {
	sl := New(intComparator)

	// 测试空表
	if sl.GetMin() != nil || sl.GetMax() != nil {
		t.Error("Empty list boundaries incorrect")
	}

	// 测试单元素表
	sl.Insert(100, "century")
	if sl.GetMin().key != 100 || sl.GetMax().key != 100 {
		t.Error("Single element boundaries incorrect")
	}

	// 测试多元素边界
	sl.Insert(200, "two")
	sl.Insert(50, "half")
	if sl.GetMin().key != 50 || sl.GetMax().key != 200 {
		t.Error("Boundaries incorrect")
	}
}

func TestOrdering(t *testing.T) {
	sl := New(intComparator)
	nums := []int{5, 3, 7, 1, 9, 2, 6}

	for _, n := range nums {
		sl.Insert(n, fmt.Sprintf("%d", n))
	}

	// 验证顺序
	expected := []int{1, 2, 3, 5, 6, 7, 9}
	current := sl.GetMin()
	for i, v := range expected {
		if current == nil || current.key != v {
			t.Errorf("Position %d: expected %d, got %v", i, v, current)
		}
		current = current.forward[0]
	}
}

/* 性能基准测试 */
const (
	smallSize  = 1_000
	mediumSize = 10_000
	largeSize  = 100_000
)

func BenchmarkInsert(b *testing.B) {
	sizes := []int{smallSize, mediumSize, largeSize}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			sl := New(intComparator)
			keys := generateSortedKeys(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, k := range keys {
					sl.Insert(k, struct{}{})
				}
				b.StopTimer()
				sl = New(intComparator)
				b.StartTimer()
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	sl := New(intComparator)
	keys := generateSortedKeys(largeSize)
	for _, k := range keys {
		sl.Insert(k, struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 搜索存在的元素
		sl.Search(keys[i%len(keys)])
		// 搜索不存在的元素（每10次）
		if i%10 == 0 {
			sl.Search(-1)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	sl := New(intComparator)
	keys := generateSortedKeys(largeSize)
	for _, k := range keys {
		sl.Insert(k, struct{}{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copyKeys := make([]int, len(keys))
		copy(copyKeys, keys)
		rand.Shuffle(len(copyKeys), func(i, j int) {
			copyKeys[i], copyKeys[j] = copyKeys[j], copyKeys[i]
		})
		b.StartTimer()

		for _, k := range copyKeys {
			sl.Delete(k)
		}
	}
}

/* 辅助函数 */
func generateSortedKeys(n int) []int {
	keys := make([]int, n)
	for i := 0; i < n; i++ {
		keys[i] = i*2 + 1 // 生成奇数避免重复
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	return keys
}

/* 对比测试（示例）*/
func BENCHMARKsMap(b *testing.B) {
	// 跳表插入
	b.Run("Insert", func(b *testing.B) {
		sl := New(intComparator)
		keys := generateSortedKeys(largeSize)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sl.Insert(keys[i%len(keys)], i)
		}
	})

	// Map插入
	b.Run("Map_Insert", func(b *testing.B) {
		m := make(map[int]int)
		keys := generateSortedKeys(largeSize)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m[keys[i%len(keys)]] = i
		}
	})

	// 跳表范围查询（利用双向链表）
	b.Run("Range", func(b *testing.B) {
		sl := New(intComparator)
		keys := generateSortedKeys(largeSize)
		for _, k := range keys {
			sl.Insert(k, struct{}{})
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := 0
			for curr := sl.GetMin(); curr != nil && count < 100; curr = curr.forward[0] {
				count++
			}
		}
	})

	// Map范围查询（需要排序）
	b.Run("Map_Range", func(b *testing.B) {
		m := make(map[int]struct{})
		keys := generateSortedKeys(largeSize)
		for _, k := range keys {
			m[k] = struct{}{}
		}
		sorted := make([]int, 0, len(keys))
		for k := range m {
			sorted = append(sorted, k)
		}
		// 预先排序
		sort.Ints(sorted)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			count := 0
			for _, k := range sorted {
				_ = k
				count++
				if count >= 100 {
					break
				}
			}
		}
	})
}
