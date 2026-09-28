package segtree

import "testing"

func TestInitialQueries(t *testing.T) {
	tree := New([]int64{1, 2, 3, 4})
	if value, _ := tree.RangeMax(0, 3); value != 4 {
		t.Fatalf("初值最大应该是 4：%d", value)
	}
	if value, _ := tree.RangeSum(0, 3); value != 10 {
		t.Fatalf("初值和应该是 10：%d", value)
	}
}

func TestSingleElementAdd(t *testing.T) {
	tree := New([]int64{1, 2, 3, 4})
	if err := tree.RangeAdd(2, 2, 5); err != nil {
		t.Fatalf("RangeAdd 出错：%v", err)
	}
	if value, _ := tree.RangeMax(2, 2); value != 8 {
		t.Fatalf("单点加不对：%d", value)
	}
}

func TestWholeRangeAdd(t *testing.T) {
	tree := New([]int64{1, 2, 3, 4})
	tree.RangeAdd(0, 3, 10)
	if value, _ := tree.RangeSum(0, 3); value != 50 {
		t.Fatalf("整体加之后和应该是 50：%d", value)
	}
}

func TestBadRange(t *testing.T) {
	tree := New([]int64{1, 2, 3, 4})
	if err := tree.RangeAdd(3, 1, 1); err != ErrRange {
		t.Fatalf("非法区间应该报 ErrRange：%v", err)
	}
	if _, err := tree.RangeMax(0, 9); err != ErrRange {
		t.Fatalf("越界应该报 ErrRange：%v", err)
	}
}

func TestSizeIsPowerOfTwo(t *testing.T) {
	if New([]int64{1, 2, 3}).Size() != 4 {
		t.Fatalf("补到 2 的幂之后大小应该是 4：%d", New([]int64{1, 2, 3}).Size())
	}
}
