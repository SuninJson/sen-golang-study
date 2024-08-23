package leetcode

import "fmt"

func findCircleNum(isConnected [][]int) int {
	fmt.Println("isConnected:", isConnected)
	n := len(isConnected)
	provinceUnionFind := NewProvinceUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				provinceUnionFind.Union(i, j)
			}
		}
	}
	fmt.Println("provinceUnionFind:", provinceUnionFind)
	return provinceUnionFind.setNum
}

type ProvinceUnionFind struct {
	parent []int
	size   []int
	help   []int
	setNum int
}

func NewProvinceUnionFind(n int) *ProvinceUnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	help := make([]int, n)
	setNum := n

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &ProvinceUnionFind{parent, size, help, setNum}
}

func (uf *ProvinceUnionFind) Find(province int) int {
	helpIndex := 0
	for province != uf.parent[province] {
		uf.help[helpIndex] = province
		helpIndex++

		province = uf.parent[province]
	}

	for helpIndex > 0 {
		helpIndex--
		uf.parent[uf.help[helpIndex]] = province
	}

	return province
}

func (uf *ProvinceUnionFind) Union(province1 int, province2 int) {
	root1 := uf.Find(province1)
	root2 := uf.Find(province2)
	if root1 == root2 {
		return
	}

	if uf.size[root1] >= uf.size[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += uf.size[root2]
	} else {
		uf.parent[root1] = root2
		uf.size[root2] += uf.size[root1]
	}
	uf.setNum--
}
