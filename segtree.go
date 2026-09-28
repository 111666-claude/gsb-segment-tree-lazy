// Package segtree 支持区间加与区间最大/求和。
// 缺陷：懒标记从不下推，查询也不把祖先的懒标记算进来。
package segtree

import "errors"

// ErrRange 表示区间不合法。
var ErrRange = errors.New("segtree: 区间不合法")

// Tree 是带懒标记的线段树。
type Tree struct {
	size int
	max  []int64
	sum  []int64
	lazy []int64
}

func maxOf(left, right int64) int64 {
	if left > right {
		return left
	}
	return right
}

// New 建树。
func New(values []int64) *Tree {
	size := 1
	for size < len(values) && size > 0 {
		size *= 2
	}
	tree := &Tree{
		size: size,
		max:  make([]int64, 2*size),
		sum:  make([]int64, 2*size),
		lazy: make([]int64, 2*size),
	}
	for index, value := range values {
		tree.max[size+index] = value
		tree.sum[size+index] = value
	}
	for index := size - 1; index >= 1; index-- {
		tree.max[index] = maxOf(tree.max[2*index], tree.max[2*index+1])
		tree.sum[index] = tree.sum[2*index] + tree.sum[2*index+1]
	}
	return tree
}

// Size 是树的叶位数（补到 2 的幂）。
func (t *Tree) Size() int {
	return t.size
}

// RangeAdd 区间加。缺陷：完全覆盖时只记 lazy，却从不往下推。
func (t *Tree) RangeAdd(left, right int, delta int64) error {
	if left < 0 || right >= t.size || left > right {
		return ErrRange
	}
	t.update(1, 0, t.size-1, left, right, delta)
	return nil
}

func (t *Tree) update(node, nodeLeft, nodeRight, left, right int, delta int64) {
	if left <= nodeLeft && nodeRight <= right {
		span := int64(nodeRight - nodeLeft + 1)
		t.max[node] += delta
		t.sum[node] += delta * span
		t.lazy[node] += delta
		return
	}
	middle := (nodeLeft + nodeRight) / 2
	if left <= middle {
		t.update(2*node, nodeLeft, middle, left, right, delta)
	}
	if right > middle {
		t.update(2*node+1, middle+1, nodeRight, left, right, delta)
	}
	// 缺陷：这里本该先 push(node)，再把 lazy[node] 合进孩子的聚合值
	t.max[node] = maxOf(t.max[2*node], t.max[2*node+1])
	t.sum[node] = t.sum[2*node] + t.sum[2*node+1]
}

// RangeMax 区间最大值。缺陷：查询时忽略沿途的懒标记。
func (t *Tree) RangeMax(left, right int) (int64, error) {
	if left < 0 || right >= t.size || left > right {
		return 0, ErrRange
	}
	return t.queryMax(1, 0, t.size-1, left, right), nil
}

func (t *Tree) queryMax(node, nodeLeft, nodeRight, left, right int) int64 {
	if left <= nodeLeft && nodeRight <= right {
		return t.max[node]
	}
	middle := (nodeLeft + nodeRight) / 2
	best := int64(-1) << 62
	if left <= middle {
		best = maxOf(best, t.queryMax(2*node, nodeLeft, middle, left, right))
	}
	if right > middle {
		best = maxOf(best, t.queryMax(2*node+1, middle+1, nodeRight, left, right))
	}
	return best
}

// RangeSum 区间和。缺陷：查询时忽略沿途的懒标记。
func (t *Tree) RangeSum(left, right int) (int64, error) {
	if left < 0 || right >= t.size || left > right {
		return 0, ErrRange
	}
	return t.querySum(1, 0, t.size-1, left, right), nil
}

func (t *Tree) querySum(node, nodeLeft, nodeRight, left, right int) int64 {
	if left <= nodeLeft && nodeRight <= right {
		return t.sum[node]
	}
	middle := (nodeLeft + nodeRight) / 2
	total := int64(0)
	if left <= middle {
		total += t.querySum(2*node, nodeLeft, middle, left, right)
	}
	if right > middle {
		total += t.querySum(2*node+1, middle+1, nodeRight, left, right)
	}
	return total
}
