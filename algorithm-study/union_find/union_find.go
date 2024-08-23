// Package union_find 包提供了用于处理并查集的实现。
package union_find

// UnionFind 结构体表示并查集，其中包含了父节点索引、集合大小和辅助数组。
type UnionFind struct {
	parent []int // 存储每个节点的父节点索引
	size   []int // 存储每个集合的大小
	help   []int // 辅助数组，用于路径压缩
}

// NewUnionFind 函数创建并返回一个新的 UnionFind 实例。
// 参数 n 指定初始集合的大小。
func NewUnionFind(n int) *UnionFind {
	// 初始化 parent 和 size 数组
	parent := make([]int, n)
	size := make([]int, n)
	help := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	// 返回初始化后的 UnionFind 实例
	return &UnionFind{parent, size, help}
}

// Find 方法用于查找 num 所在集合的根节点。
// 参数 num 是要查找的节点。
func (unionFind *UnionFind) Find(num int) int {
	// helpIndex 用于记录路径压缩的当前位置
	helpIndex := 0
	// 循环查找根节点
	for ; unionFind.parent[num] != num; helpIndex++ {
		unionFind.help[helpIndex] = num
		num = unionFind.parent[num]
	}
	// 路径压缩，将路径上的所有节点的父节点指向根节点
	for helpIndex > 0 {
		// 注意helpIndex中的最后一个元素是根节点，不需要处理，所以先将helpIndex减一
		helpIndex--
		onTheWayNum := unionFind.help[helpIndex]
		unionFind.parent[onTheWayNum] = num
	}
	// 返回根节点
	return num
}

// Union 方法用于合并 num1 和 num2 所在的两个集合。
// 参数 num1 和 num2 是要合并的两个节点。
func (unionFind *UnionFind) Union(num1, num2 int) {
	// 找到 num1 和 num2 的根节点
	root1 := unionFind.Find(num1)
	root2 := unionFind.Find(num2)

	// 如果根节点不同且 root1 的集合大于 root2，则将 root2 连接到 root1
	if root1 != root2 && unionFind.size[root1] >= unionFind.size[root2] {
		unionFind.parent[root2] = root1
		unionFind.size[root1] += unionFind.size[root2]
	} else {
		// 否则，将 root1 连接到 root2
		unionFind.parent[root1] = root2
		unionFind.size[root2] += unionFind.size[root1]
	}
}

// isSameSet 方法判断 num1 和 num2 是否在同一个集合中。
// 参数 num1 和 num2 是要判断的两个节点。
func (unionFind *UnionFind) isSameSet(num1, num2 int) bool {
	// 如果两个节点的根节点相同，则它们在同一个集合中
	return unionFind.Find(num1) == unionFind.Find(num2)
}
