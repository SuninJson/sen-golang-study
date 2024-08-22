package union_find

import (
	"testing"
)

// TestNewUnionFind 测试 NewUnionFind 函数是否正确初始化 UnionFind 结构体。
func TestNewUnionFind(t *testing.T) {
	n := 5 // 假定集合的大小为5
	uf := NewUnionFind(n)

	if len(uf.parent) != n || len(uf.size) != n || len(uf.help) != n {
		t.Errorf("NewUnionFind did not initialize UnionFind with the correct size")
	}

	for i := 0; i < n; i++ {
		if uf.parent[i] != i || uf.size[i] != 1 {
			t.Errorf("NewUnionFind did not initialize UnionFind with the correct values")
		}
	}
}

// TestUnionFind_Find 测试 UnionFind 结构体的 Find 方法。
func TestUnionFind_Find(t *testing.T) {
	uf := NewUnionFind(5)

	// 测试 Find 方法是否返回正确的根节点
	for i := 0; i < 5; i++ {
		root := uf.Find(i)
		if root != i {
			t.Errorf("Find method did not return the correct root node, got %d, want %d", root, i)
		}
	}

	expectedRoot := 0 // 假定的根节点
	uf.Union(0, 1)
	uf.Union(0, 2)
	uf.Union(0, 3)
	uf.Union(0, 4)

	for i := 0; i < 5; i++ {
		root := uf.Find(i)
		if root != expectedRoot {
			t.Errorf("Find method did not return the correct root node, got %d, want %d", root, expectedRoot)
		}
	}

}

// TestUnionFind_Union 测试 UnionFind 结构体的 Union 方法。
func TestUnionFind_Union(t *testing.T) {
	uf := NewUnionFind(5)

	// 测试 Union 方法是否能正确地合并两个集合
	uf.Union(0, 1)
	if !uf.isSameSet(0, 1) {
		t.Errorf("Union method did not merge the sets correctly")
	}
}

// TestUnionFind_isSameSet 测试 UnionFind 结构体的 isSameSet 方法。
func TestUnionFind_isSameSet(t *testing.T) {
	uf := NewUnionFind(5)

	// 测试 isSameSet 方法是否能正确判断两个节点是否在同一个集合中
	if uf.isSameSet(0, 1) {
		t.Errorf("isSameSet method incorrectly reported that nodes 0 and 1 are in the same set")
	}
}
