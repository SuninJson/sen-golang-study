package leetcode

// findMaximumXOR 寻找数组中最大的异或值。
// nums: 整数数组。
// 返回数组中任意两个数字的最大异或值。
func findMaximumXOR(nums []int) int {
	// 初始化最大异或值为0。
	ans := 0
	// 如果数组为空，直接返回0。
	if len(nums) == 0 {
		return ans
	}

	// 创建一个二进制数字Trie树，并插入第一个数字。
	numTrie := &BinaryNumTrie{}
	numTrie.Insert(nums[0])
	// 遍历数组中的每个数字，计算当前数字与Trie树中所有数字的异或值，并更新最大异或值。
	for i := 1; i < len(nums); i++ {
		// 计算当前数字与Trie树中所有数字的最大异或值，并更新ans。
		ans = max(ans, numTrie.getMaxXOR(nums[i]))
		// 将当前数字插入Trie树。
		numTrie.Insert(nums[i])
	}

	// 返回最大异或值。
	return ans
}

// BinaryNumTrie 定义了一个二进制数字的Trie树结构，用于高效地查询和插入二进制表示的数字。
type BinaryNumTrie struct {
	// next 是一个长度为2的数组，
	// next[0]表示当前位为0的子节点，next[1]表示当前位为1的子节点。
	next [2]*BinaryNumTrie
}

// Insert 将一个整数插入到Trie树中。
// num: 需要插入的整数。
// 从最高位到最低位，根据num的二进制表示逐位创建节点并插入到Trie树中。
func (root *BinaryNumTrie) Insert(num int) {
	cur := root
	// 从高位到低位遍历num的二进制表示，插入到Trie树中。
	// 因为题目中的数字范围为0 ~ 2的31次方，最高位为符号位，因此最多需要31位。
	for i := 30; i >= 0; i-- {
		// 获取当前位的路径
		path := 0
		if (num & (1 << i)) != 0 {
			path = 1
		}
		// 如果当前路径下子节点为空，则创建新的子节点。
		if cur.next[path] == nil {
			cur.next[path] = &BinaryNumTrie{}
		}
		// 移动到下一个比特位。
		cur = cur.next[path]
	}
}

// getMaxXOR 计算给定数字与Trie树中所有数字的异或最大值。
// num: 需要计算的数字。
// 返回num与Trie树中所有数字的异或最大值。
func (root *BinaryNumTrie) getMaxXOR(num int) int {
	cur := root
	max := 0
	// 从高位到低位遍历num的二进制表示，尝试找到与当前位相反的比特位，以最大化异或值。
	for i := 30; i >= 0; i-- {
		numBit := 0
		// 如果当前位为1，则通过numBit记录。
		if (num & (1 << i)) != 0 {
			numBit = 1
		}
		// 因为当前位异或结果为1时，能获得更大的异或结果，所以期望能找到与当前位相反的比特位。
		wantBit := numBit ^ 1
		// 如果期望的路径不存在，则只能选择相同的比特位。
		if cur.next[wantBit] == nil {
			wantBit = numBit
		}
		// 更新最大异或值，并移动到下一个比特位。
		max |= (wantBit ^ numBit) << i
		cur = cur.next[wantBit]
	}

	// 返回最大异或值。
	return max
}
