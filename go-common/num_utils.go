package common

func GetBinaryNum(num int) []int {
	var binaryNum []int
	for num > 0 {
		binaryNum = append(binaryNum, num%2)
		num = num / 2
	}
	return binaryNum
}

// GetBinaryNumBit 返回指定二进制数的特定位的值
// 参数num是待查询的二进制数，参数bitIndex是指定的位置索引
// 返回值是num的二进制中index（从0开始）位置的值（0或1）
func GetBinaryNumBit(num int, index int) int {
	// 默认假设位值为0
	bit := 0
	// 使用与操作检查指定位置的值是否为1
	// 如果结果不为0，则说明该位为1
	if (num & (1 << index)) != 0 {
		// 将bit设置为1
		bit = 1
	}
	// 返回计算出的位值
	return bit
}
