# segment-tree-lazy

带懒标记的线段树（区间加 + 区间最大 + 区间求和），只用 Go 标准库。

```
go test ./...
go run ./cmd/segment-tree-lazy --values 1,2,3,4 --ops 'add,0,2,10;max,0,3;sum,0,3'
go run ./cmd/segment-tree-lazy --values 1,2,3,4 --ops 'add,1,1,5;sum,0,3;max,0,1'
```

## 口径（README 为准）

- **懒标记必须下推**：`RangeAdd` 命中完全覆盖的节点时可以只记 `lazy`，
  但访问孩子前必须 `push`（把 `lazy` 加到孩子、清空自己）。
  `--ops 'add,0,1,10;add,0,0,1;sum,0,0;sum,0,1'` 按口径是 `sum(0,0)=12`、`sum(0,1)=23`，
  现在第二次 `add` 把 node(0,1) 上挂着的 lazy 丢了，给 `sum(0,0)=2`。
- **查询也要下推**：`--ops 'add,0,1,10;sum,0,0;sum,0,1'` 里第一次 `sum(0,0)` 要走孩子，
  按口径是 11，现在给 1（懒标记没算进孩子）。
- **查询要把懒标记算进来**：`RangeMax`/`RangeSum` 递归前先下推，再取孩子。
- 不变量：任意操作序列之后，`RangeSum(l,r)` 与 `RangeMax(l,r)` 等于对原数组逐步加完之后的真实值
  （与 O(n) 暴力实现逐条对照，随机 5000 次操作差异必须为 0）；越界或 `l > r` 报 `ErrRange`。
- 规模：10^5 个元素、10^5 次操作 1 秒内完成，单次操作 O(log n)；不许每次操作都重建整棵树。

## 现在的行为

- `update` 里注释的位置本该先 `push(node)`：现在完全不下推，孩子的聚合值也没加上 `lazy[node]`。
- `queryMax`/`querySum` 直接用 `max[node]`/`sum[node]`，查询区间没完全覆盖被加节点时就偏小。
- 整体加（`add` 覆盖全树）碰巧是对的，因为根节点被直接更新过。

## 计划新增的能力

```
--values 初值     逗号分隔
--ops 操作脚本    add,l,r,v / max,l,r / sum,l,r，分号分隔
```

## 输出

```
sum(0,0)=12
sum(0,1)=23
```

## 目录

```
segtree.go                   线段树
cmd/segment-tree-lazy        命令行入口
segtree_test.go              go test 用例
```
