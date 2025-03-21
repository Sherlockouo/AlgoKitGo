package skiplist

import (
	"math/rand"
	"time"
)

const (
	maxLevel    = 16  // 最大层数，可根据数据量调整
	probability = 0.5 // 节点升级概率，影响层级分布
)

// SkipListNode 跳表节点结构
type SkipListNode struct {
	key     any             // 节点键值，用于比较和排序
	value   any             // 节点存储的值
	forward []*SkipListNode // 各层前进指针数组
	prev    *SkipListNode   // 底层前驱指针（实现双向链表特性）
}

// SkipList 跳表主结构
type SkipList struct {
	head    *SkipListNode      // 头节点（哨兵节点）
	level   int                // 当前最大有效层数
	length  int                // 节点数量
	rand    *rand.Rand         // 随机数生成器
	compare func(a, b any) int // 键比较函数
}

// New 创建跳表实例
// compare: 自定义比较函数，返回：
//
//	-1 当a < b
//	 0 当a == b
//	 1 当a > b
func New(compare func(a, b any) int) *SkipList {
	return &SkipList{
		head: &SkipListNode{
			forward: make([]*SkipListNode, maxLevel),
		},
		rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
		compare: compare,
	}
}

/* 核心方法实现 */

// randomLevel 生成随机层数（节点高度）
// 按照概率因子生成层级，保证上层节点数是下层的1/probability倍
// 时间复杂度：O(1) 空间复杂度：O(1)
func (s *SkipList) randomLevel() int {
	level := 1
	// 每次有probability的概率升级层数
	for s.rand.Float64() < probability && level < maxLevel {
		level++
	}
	return level
}

// Insert 插入/更新键值对
// 平均时间复杂度：O(logn) 最坏情况：O(n)
func (s *SkipList) Insert(key, value any) {
	update := make([]*SkipListNode, maxLevel) // 各层需要更新的节点
	current := s.head

	// 从最高层向下搜索插入位置
	for i := s.level - 1; i >= 0; i-- {
		// 在当前层找到最后一个小于key的节点
		for current.forward[i] != nil &&
			s.compare(current.forward[i].key, key) < 0 {
			current = current.forward[i]
		}
		update[i] = current // 记录该层的插入位置
	}

	// 生成新节点层数
	newLevel := s.randomLevel()
	// 扩展跳表层级（当新节点层级更高时）
	if newLevel > s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.head // 新增层的前驱指向头节点
		}
		s.level = newLevel
	}

	// 创建新节点
	newNode := &SkipListNode{
		key:     key,
		value:   value,
		forward: make([]*SkipListNode, newLevel),
	}

	// 更新各层指针
	for i := range newLevel {
		// 将新节点插入到各层链表
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode

		// 维护底层双向指针（仅第0层）
		if i == 0 {
			// 设置前驱指针
			if update[i] != s.head {
				newNode.prev = update[i]
			}
			// 更新后继节点的前驱指针
			if newNode.forward[i] != nil {
				newNode.forward[i].prev = newNode
			}
		}
	}

	s.length++
}

// Search 查找指定键的节点
// 平均时间复杂度：O(logn) 最坏情况：O(n)
func (s *SkipList) Search(key any) *SkipListNode {
	current := s.head

	// 从顶层到底层逐层搜索
	for i := s.level - 1; i >= 0; i-- {
		// 在当前层向前搜索
		for current.forward[i] != nil &&
			s.compare(current.forward[i].key, key) < 0 {
			current = current.forward[i]
		}
	}

	// 在底层验证结果
	target := current.forward[0]
	if target != nil && s.compare(target.key, key) == 0 {
		return target
	}
	return nil
}

// Delete 删除指定键的节点
// 平均时间复杂度：O(logn) 最坏情况：O(n)
func (s *SkipList) Delete(key any) {
	update := make([]*SkipListNode, maxLevel)
	current := s.head

	// 查找各层的前驱节点
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil &&
			s.compare(current.forward[i].key, key) < 0 {
			current = current.forward[i]
		}
		update[i] = current
	}

	// 定位到目标节点
	target := current.forward[0]
	if target == nil || s.compare(target.key, key) != 0 {
		return // 节点不存在
	}

	// 更新各层指针
	for i := 0; i < s.level; i++ {
		// 跳过不存在该节点的层级
		if update[i].forward[i] != target {
			break
		}
		update[i].forward[i] = target.forward[i]
	}

	// 维护双向链表指针
	if target.forward[0] != nil {
		target.forward[0].prev = target.prev
	}

	// 调整跳表层级高度
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}

	s.length--
}

/* 辅助方法 */

// Size 返回元素数量
func (s *SkipList) Size() int {
	return s.length
}

// GetMin 获取最小键的节点（利用双向链表特性）
func (s *SkipList) GetMin() *SkipListNode {
	return s.head.forward[0]
}

// GetMax 获取最大键的节点（利用双向链表特性）
func (s *SkipList) GetMax() *SkipListNode {
	current := s.head.forward[0]
	if current == nil {
		return nil
	}
	for current.forward[0] != nil {
		current = current.forward[0]
	}
	return current
}
